using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class FormatFunction
    {

        static void TypeAndNameLeftParenthesis(FormatData formatData, ref string? returnType, ref string functionName, Token firstToken)
        {
            const string thisFunctionName = "TypeAndNameLeftParenthesis";

            List<Token> returnTypeTokenList = new List<Token>();
            FormatUtils.LoopUntilRightParenthesis(formatData, returnTypeTokenList);

            returnType = TokenUtils.JoinTextInListOfTokens(returnTypeTokenList);

            Token? nextToken = formatData.GetToken();
            if (nextToken == null)
            {
                formatData.EndOfFileError(firstToken, thisFunctionName);
                return;
            }
            if (nextToken.Value.Type != TokenType.Identifier)
            {
                formatData.MissingExpectedTypeError(nextToken, "missing identifier in function name", thisFunctionName);
                return;
            }
            functionName = nextToken.Value.Text;
            formatData.Increment();
        }

        static void TypeAndNameOther(FormatData formatData, ref string? returnType, ref string functionName, Token firstToken)
        {
            const string thisFunctionName = "TypeAndNameIdentifier";



            List<Token> returnTypeTokenList = new List<Token>();
            FormatUtils.FillVarType(formatData, returnTypeTokenList);

            returnType = TokenUtils.JoinTextInListOfTokens(returnTypeTokenList);

            Token? nextToken = formatData.GetToken();
            if (nextToken == null) {
                formatData.EndOfFileError(firstToken, thisFunctionName);
                return;
            }
            if (nextToken.Value.Type != TokenType.Identifier) {
                formatData.MissingExpectedTypeError(nextToken, "missing identifier in function name", thisFunctionName);
                return;
            }
            functionName = nextToken.Value.Text;
            formatData.Increment();
        }

        static void TypeAndNameIdentifier(FormatData formatData, ref string? ReturnType, ref string functionName, Token firstToken)
        {
            const string thisFunctionName = "TypeAndNameIdentifier";

            int index = formatData.TokenIndex;
            Token? tempToken = formatData.GetNextToken();

            if (tempToken == null)
            {
                formatData.EndOfFileError(tempToken, thisFunctionName);
                return;
            }
            Token nextToken = tempToken.Value;
            //void type
            if (nextToken.Type == TokenType.LeftParenthesis)
            {
                ReturnType = null;
                functionName = firstToken.Text;
                return;
            }
            if (nextToken.Type == TokenType.Identifier) 
            {
                tempToken = formatData.GetNextToken();
                if (tempToken == null)
                {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }
                if (tempToken.Value.Type != TokenType.LeftParenthesis)
                {
                    formatData.MissingExpectedTypeError(tempToken, "missing expected '(' in function", thisFunctionName);
                    return;
                }

                ReturnType = firstToken.Text;
                functionName = nextToken.Text;
                return;
            }
            formatData.TokenIndex = index;
            TypeAndNameOther(formatData, ref ReturnType, ref functionName, firstToken);
        }

        static void GetFunctionTypeAndName(FormatData formatData, ref string? returnType, ref string functionName)
        {
            const string thisFunctionName = "GetFunctionTypeAndName";

            Token? tempToken = null;
            tempToken = formatData.GetNextToken();
            if (tempToken == null) {
                formatData.EndOfFileError(tempToken, thisFunctionName);
                return;
            }
            Token firstToken = tempToken.Value;
            
            switch (firstToken.Type)
            {
                case TokenType.Identifier:
                    TypeAndNameIdentifier(formatData, ref returnType, ref functionName, firstToken);
                    break;

                case TokenType.LeftParenthesis:
                    TypeAndNameLeftParenthesis(formatData, ref returnType, ref functionName, firstToken);
                    break;

                default:
                    TypeAndNameOther(formatData, ref returnType, ref functionName, firstToken);
                    break;
            }
        }

        static void GetInterfaceMethodTypeAndName(FormatData formatData, ref string? returnType, ref string functionName)
        {
            const string thisFunctionName = "GetFunctionTypeAndName";

            Token? tempToken = null;
            tempToken = formatData.GetToken();
            if (tempToken == null)
            {
                formatData.EndOfFileError(tempToken, thisFunctionName);
                return;
            }
            Token firstToken = tempToken.Value;
            switch (firstToken.Type)
            {
                case TokenType.Identifier:
                    TypeAndNameIdentifier(formatData, ref returnType, ref functionName, firstToken);
                    break;

                case TokenType.LeftParenthesis:
                    TypeAndNameLeftParenthesis(formatData, ref returnType, ref functionName, firstToken);
                    break;

                default:
                    TypeAndNameOther(formatData, ref returnType, ref functionName, firstToken);
                    break;
            }

        }

        public static void ProcessFunction(FormatData formatData, List<Function> functions, Token fnToken)
        {
            const string thisFunctionName = "ProcessFunction";

            List<Variable> parameters = new();
            string? returnType = null;
            string functionName = "";

            GetFunctionTypeAndName(formatData, ref returnType, ref functionName);
            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return;
            }
            //Expect '('
            if (formatData.ExpectType(TokenType.LeftParenthesis, "Missing expected '(' in fn", thisFunctionName) == false) {
                return;
            }
            formatData.Increment();

            FunctionUtils.FindParameters(ref formatData, parameters);
            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return;
            }

            //Expect ')'
            if (formatData.ExpectType(TokenType.RightParenthesis, "Missing expected ')' in fn", thisFunctionName) == false) {
                return;
            }
            //Expect '{'
            if (formatData.ExpectNextType(TokenType.LeftBrace, "Missing expected '{' in fn", thisFunctionName) == false) {
                return;
            }
            formatData.Increment();
            CodeBlock? innerBlock = FormatBody.FillBody(formatData);
            //Expect '}'
            if (formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' in fn", thisFunctionName) == false) {
                return;
            }

            Function function = new Function { 
                InnerBlock  = innerBlock,
                Parameters = parameters,
                Name = functionName,
                ReturnType = returnType,
                startingToken = fnToken,
            };
            functions.Add(function);
        }

        public static void ProcessInterfaceFunction(FormatData formatData, List<Function> functions, Token fnToken)
        {
            const string thisFunctionName = "ProcessInterfaceFunction";

            List<Variable> parameters = new();
            string? returnType = null;
            string functionName = "";

            GetInterfaceMethodTypeAndName(formatData, ref returnType, ref functionName);
            if (formatData.IsError())
            {
                formatData.AddTrace(thisFunctionName);
                return;
            }
            //Expect '('
            if (formatData.ExpectType(TokenType.LeftParenthesis, "Missing expected '(' in interface", thisFunctionName) == false)
            {
                return;
            }
            formatData.Increment();

            FunctionUtils.FindParameters(ref formatData, parameters);
            if (formatData.IsError())
            {
                formatData.AddTrace(thisFunctionName);
                return;
            }

            //Expect ')'
            if (formatData.ExpectType(TokenType.RightParenthesis, "Missing expected ')' in interface", thisFunctionName) == false)
            {
                return;
            }
            formatData.Increment();

            Function function = new Function
            {
                InnerBlock = null,
                Parameters = parameters,
                Name = functionName,
                ReturnType = returnType,
            };
            functions.Add(function);
        }
    }
}
