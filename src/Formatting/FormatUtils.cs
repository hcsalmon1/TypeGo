using System;
using System.Collections.Generic;

namespace TypeGo
{
    public static class FormatUtils
    {

        static WhileLoopAction LoopUntilEndInner(FormatData formatData, BlockData blockData, Token token, TokenType lastTokenType, ref int openCurlyBraceCount, bool stopAtSemicolon)
        {
            if (token.Type == TokenType.LeftBrace) {
                openCurlyBraceCount += 1;
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.RightBrace) {
                openCurlyBraceCount -= 1;
                return WhileLoopAction.Continue;
            }

            if (openCurlyBraceCount != 0) {
                return WhileLoopAction.Continue;
            }

            if (stopAtSemicolon == true) {

                if (token.Type == TokenType.Semicolon) {

                    blockData.Tokens.Add(token);
                    return WhileLoopAction.Break;
                }
            }
            if (token.Type == TokenType.NewLine) {

                blockData.Tokens.Add(token);
                if (IsLineContinuingToken(lastTokenType) == false) {
                    formatData.Increment();
                    return WhileLoopAction.Break;
                }
            }
            if (token.Type == TokenType.EndComment) {
                return WhileLoopAction.Break;
            }
            return WhileLoopAction.Continue;
        }

        public static void LoopTokensUntilLineEnd(FormatData formatData, BlockData blockData, bool stopAtSemicolon)
        {
            TokenType lastTokenType = TokenType.NA;

            int openCurlyBraceCount = 0;

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                Token? tempToken = formatData.GetToken();
                if (tempToken == null) {
                    break;
                }
                Token token = tempToken.Value;
                WhileLoopAction loopResult = LoopUntilEndInner(formatData, blockData, token, lastTokenType, ref openCurlyBraceCount, stopAtSemicolon);
                if (loopResult == WhileLoopAction.Break) {
                    break;
                }

                formatData.Increment();
                blockData.Tokens.Add(token);
                lastTokenType = token.Type;
            }
            //while (formatData.TokenIndex < formatData.TokenList.Count) {
            //    Token? finalToken = formatData.GetToken();
            //    if (finalToken == null)
            //    {
            //        formatData.EndOfFileError(finalToken, "LoopTokensUntilLineEnd");
            //        return;
            //    }
            //    if (finalToken.Value.Type == TokenType.NewLine) {
            //        formatData.Increment();
            //        break;
            //    }
            //    formatData.Increment();
            //}
        }

        static bool IsLineContinuingToken(TokenType type)
        {
            switch (type)
            {
                case TokenType.Minus:
                case TokenType.Plus:
                case TokenType.Divide:
                case TokenType.Multiply:
                case TokenType.Equals:
                case TokenType.And:
                case TokenType.AndAnd:
                case TokenType.Or:
                case TokenType.OrOr:
                case TokenType.PlusEquals:
                case TokenType.MinusEquals:
                case TokenType.MultiplyEquals:
                case TokenType.DivideEquals:
                case TokenType.GreaterThan:
                case TokenType.LessThan:
                case TokenType.EqualsEquals:
                case TokenType.GreaterThanEquals:
                case TokenType.LessThanEquals:
                case TokenType.Modulus:
                case TokenType.ModulusEquals:
                case TokenType.NotEquals:
                case TokenType.LeftBrace:
                case TokenType.LeftSquareBracket:
                case TokenType.Comma:
                case TokenType.FullStop:
                    return true;
                
                default:
                    return false;
            }
        }

        public static bool IsPointerDeclaration(FormatData formatData)
        {
            const string thisFunctionName = "IsPointerDeclaration";

            int index = formatData.TokenIndex;
            TokenType lastType = TokenType.NA;
            bool identifierFound = false;

            while (index < formatData.TokenList.Count) {

                if (index >= formatData.TokenList.Count) {
                    formatData.EndOfFileError(null, thisFunctionName);
                    return false;
                }
                Token token = formatData.TokenList[index];
                if (token.Type == TokenType.LeftSquareBracket) {

                    if (identifierFound == true) {
                        return false;
                    }
                    return true;
                }
                if (token.Type == TokenType.Identifier) {

                    identifierFound = true;
                    if (lastType == TokenType.Identifier) {
                        return true;
                    }
                }
                else if (TokenUtils.IsVarType(token)) {
                    return true;
                } 
                else if (TokenUtils.IsAssignmentOperator(token.Type)) {
                    return false;
                }
                lastType = token.Type;
                index += 1;
            }
            return false;
        }

        static bool HandleToken(ref FormatData formatData, Token token, List<Token> returnTypeTokenList, ref TokenType lastType)
        {
            if (TokenUtils.IsVarType(token))
            {
                returnTypeTokenList.Add(token);
                return false;
            }


            if (token.Type == TokenType.Identifier)
            {
                if (lastType == TokenType.Identifier || lastType == TokenType.RightBrace) {
                    return true;
                }
                if (TokenUtils.IsVarType(lastType)) {
                    return true;
                }

                bool wasPartOfType =
                    lastType == TokenType.FullStop ||
                    lastType == TokenType.RightSquareBracket ||
                    lastType == TokenType.Multiply ||
                    lastType == TokenType.LeftParenthesis ||
                    lastType == TokenType.NA;

                if (wasPartOfType)
                {
                    returnTypeTokenList.Add(token);
                    return false;
                }

                formatData.UnexpectedTypeError(token, "Found unexpected identifier in parameters", "HandleToken");
                return true;
            }
            bool isSkippableToken =
                token.Type == TokenType.Tab ||
                token.Type == TokenType.NewLine ||
                token.Type == TokenType.RightParenthesis ||
                token.Type == TokenType.Comma;

            if (isSkippableToken)
            {
                return false;
            }
            returnTypeTokenList.Add(token);
            return false;
        }

        static WhileLoopAction VarTypeInnerCode(ref FormatData formatData, List<Token> returnTypeTokenList, ref int whileCount, ref TokenType lastType)
        {
            const string thisFunctionName = "VarTypeInnerCode";
            const int MAX = 10000;

            if (Debugging.InfiniteWhileCheck(ref whileCount, MAX)) {
                formatData.Result = FormatResult.Internal_Error;
                formatData.ErrorDetail = "Infinite while loop in VarTypeInnerLoop, FormatUtils";
                return WhileLoopAction.Error;
            }
            int indexBefore = formatData.TokenIndex;

            Token? token = formatData.GetToken();
            if (token == null) {
                formatData.EndOfFileError(null, thisFunctionName);
                return WhileLoopAction.Return;
            }

            bool shouldBreak = HandleToken(ref formatData, token.Value, returnTypeTokenList, ref lastType);
            if (shouldBreak == true) {
                return WhileLoopAction.Break;
            }

            formatData.IncrementIfSame(indexBefore);
            lastType = token.Value.Type;
            return WhileLoopAction.Continue;
        }

        public static void FillVarType(FormatData formatData, List<Token> returnTypeTokenList)
        {
            const string thisFunctionName = "FillVarType";

            int whileCount = 0;

            TokenType lastType = TokenType.NA;

            //Get to ')' and then don't increment
            while (formatData.TokenIndex < formatData.TokenList.Count) {

                WhileLoopAction iterationResult = VarTypeInnerCode(ref formatData, returnTypeTokenList, ref whileCount, ref lastType);
                if (iterationResult == WhileLoopAction.Break) {
                    break;
                }
                if (iterationResult == WhileLoopAction.Return) {
                    return;
                }
            }
        }
    
        
        public static void LoopUntilRightParenthesis(FormatData formatData, List<Token> returnTypeTokenList)
        {
            const string thisFunctionName = "LoopUntilRightParenthesis";

            int whileCount = 0;

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {
                const int MAX = 10000;

                if (Debugging.InfiniteWhileCheck(ref whileCount, MAX)) {
                    formatData.Result = FormatResult.Internal_Error;
                    formatData.ErrorDetail = "Infinite while loop in LoopUntilRightParenthesis, FormatUtils";
                    return;
                }
                int indexBefore = formatData.TokenIndex;

                Token? token = formatData.GetToken();
                if (token == null) {
                    formatData.EndOfFileError(null, thisFunctionName);
                    return;
                }

                returnTypeTokenList.Add(token.Value);

                formatData.IncrementIfSame(indexBefore);

                if (token.Value.Type == TokenType.RightParenthesis) {
                    break;
                }
            }
        }
    
    }
}
