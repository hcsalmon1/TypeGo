using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class FunctionUtils
    {

        static void AddParameter(List<Variable> parameters, ref Variable tempParameter)
        {
            if (tempParameter.NameToken == null) {
                return;
            }
            if (tempParameter.TypeList.Count == 0) {
                return;
            }
            Variable parameterCopy = new();
            parameterCopy.TypeList = new List<Token>();
            for (int i = 0; i < tempParameter.TypeList.Count; i++) {
                parameterCopy.TypeList.Add(tempParameter.TypeList[i]);
            }
            parameterCopy.NameToken = TokenUtils.CopyTokenList(tempParameter.NameToken);

            parameters.Add(parameterCopy);
            tempParameter.SetToDefaults();
        }

        static void HandleToken(ref FormatData formatData, Token token, ref Variable tempParameter, ref TokenType lastType)
        {
            if (token.Type == TokenType.Identifier)
            {
                if (lastType == TokenType.Identifier || lastType == TokenType.RightBrace) {
                    tempParameter.NameToken.Add(token);
                    return;
                }
                if (TokenUtils.IsVarType(lastType))
                {
                    tempParameter.NameToken.Add(token);
                    return;
                }

                bool wasPartOfType =
                    lastType == TokenType.FullStop ||
                    lastType == TokenType.RightSquareBracket ||
                    lastType == TokenType.Multiply ||
                    lastType == TokenType.Comma ||
                    lastType == TokenType.LeftParenthesis;

                if (wasPartOfType) {
                    tempParameter.TypeList.Add(token);
                    return;
                }

                formatData.UnexpectedTypeError(token, "Found unexpected identifier in parameters", "HandleToken");
                return;
            }
            bool isSkippableToken =
                token.Type == TokenType.Tab ||
                token.Type == TokenType.NewLine ||
                token.Type == TokenType.RightParenthesis ||
                token.Type == TokenType.Comma;

            if (isSkippableToken) {
                return;
            }
            tempParameter.TypeList.Add(token);
        }

        static WhileLoopAction ParameterInnerLoop(ref FormatData formatData, List<Variable> parameters, ref Variable tempParameter, ref int whileCount, ref TokenType lastType)
        {
            const string thisFunctionName = "ParameterInnerLoop";
            const int MAX = 10000;

            if (Debugging.InfiniteWhileCheck(ref whileCount, MAX)) {
                formatData.Result = FormatResult.Internal_Error;
                formatData.ErrorDetail = "infinite while loop in ParameterInnerLoop, FunctionUtils";
                return WhileLoopAction.Error;
            }
            int indexBefore = formatData.TokenIndex;

            Token? token = formatData.GetToken();
            if (token == null) {
                formatData.EndOfFileError(null, thisFunctionName);
                return WhileLoopAction.Return;
            }

            HandleToken(ref formatData, token.Value, ref tempParameter, ref lastType);

            if (token.Value.Type == TokenType.Comma) {
                AddParameter(parameters, ref tempParameter);
            }
            if (token.Value.Type == TokenType.RightParenthesis){
                AddParameter(parameters, ref tempParameter);
                return WhileLoopAction.Break;
            }

            formatData.IncrementIfSame(indexBefore);
            lastType = token.Value.Type;
            return WhileLoopAction.Continue;
        }

        public static void FindParameters(ref FormatData formatData, List<Variable> parameters)
        {
            const string thisFunctionName = "FindParameters";

            int whileCount = 0;

            Variable tempParameter = new();
            tempParameter.SetToDefaults();
            TokenType lastType = TokenType.LeftParenthesis;

            //Get to ')' and then don't increment
            while (formatData.TokenIndex < formatData.TokenList.Count) {

                WhileLoopAction iterationResult = ParameterInnerLoop(ref formatData, parameters, ref tempParameter, ref whileCount, ref lastType);
                if (iterationResult == WhileLoopAction.Break) {
                    break;
                }
                if (iterationResult == WhileLoopAction.Return) {
                    return;
                }
            }
        }
    }
}
