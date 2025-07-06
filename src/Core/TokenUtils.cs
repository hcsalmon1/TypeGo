using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class TokenUtils
    {
        public const int INVALID_INDEX = -1;


        public static List<Token> CopyTokenList(List<Token> tokenList)
        {
            List<Token> copyTokenList = new List<Token>();
            if (tokenList.Count == 0)
            {
                return copyTokenList;
            }
            for (int i = 0; i < tokenList.Count; i++)
            {
                copyTokenList.Add(tokenList[i]);
            }
            return copyTokenList;
        }


        public static bool AddVarTypeList(ref ConvertData convertData, List<Token> typeList, string thisFunctionName)
        {
            if (typeList.Count != 0)
            {
                for (int i = 0; i < typeList.Count; i++)
                {
                    convertData.GeneratedCode.Append(typeList[i].Text);
                }
                return true;
            }

            convertData.UnexpectedTypeError(null, "token type list was empty", thisFunctionName);

            return false;
        }

        public static string? JoinTextInListOfTokens(List<Token> tokenList)
        {
            if (tokenList == null) {
                return null;
            }
            if (tokenList.Count == 0) {
                return null;
            }
            StringBuilder outputSB = new();
            for (int i = 0; i < tokenList.Count; i++) {
                outputSB.Append(tokenList[i].Text);
            }
            return outputSB.ToString();
        }

        public static string? JoinTextInListOfTokensWithSpaces(List<Token> tokenList)
        {
            if (tokenList == null)
            {
                return null;
            }
            if (tokenList.Count == 0)
            {
                return null;
            }
            StringBuilder outputSB = new();
            for (int i = 0; i < tokenList.Count; i++)
            {
                outputSB.Append(tokenList[i].Text);
                outputSB.Append(' ');
            }
            return outputSB.ToString();
        }


        public static void PrintToken(Token? token)
        {
            if (token == null)
            {
                Fmt.Print("null");
                return;
            }
            Fmt.Print(token.Value.Text);
        }
        public static void PrintlnToken(Token? token)
        {
            if (token == null)
            {
                Fmt.Println("null");
                return;
            }
            Fmt.Println(token.Value.Text);
        }

        public static bool IsValidToken(ref ConvertData convertData, Token? token, TokenType expectedType)
        {
            if (token == null) {
                return false;
            }
            if (token.Value.Type != expectedType) {
                return false;
            }

            return true;
        }

        public static bool IsVarTypePart(Token token)
        {
            switch (token.Type)
            {
                case TokenType.LeftSquareBracket:
                case TokenType.Multiply:
                    return true;
                default:
                    return false;
            }
        }

        public static bool IsVarType(Token token)
        {
            switch (token.Type)
            {
                case TokenType.Int:
                case TokenType.Int8:
                case TokenType.Int16:
                case TokenType.Int32:
                case TokenType.Int64:
                case TokenType.Uint:
                case TokenType.Uint8:
                case TokenType.Uint16:
                case TokenType.Uint32:
                case TokenType.Uint64:
                case TokenType.Float32:
                case TokenType.Float64:
                case TokenType.String:
                case TokenType.Byte:
                case TokenType.Rune:
                case TokenType.Bool:
                case TokenType.Error:
                    return true;
                default:
                    return false;
            }
        }
        public static bool IsVarType(TokenType tokenType)
        {
            switch (tokenType)
            {
                case TokenType.Int:
                case TokenType.Int8:
                case TokenType.Int16:
                case TokenType.Int32:
                case TokenType.Int64:
                case TokenType.Uint:
                case TokenType.Uint8:
                case TokenType.Uint16:
                case TokenType.Uint32:
                case TokenType.Uint64:
                case TokenType.Float32:
                case TokenType.Float64:
                case TokenType.String:
                case TokenType.Byte:
                case TokenType.Rune:
                case TokenType.Bool:
                case TokenType.Error:
                case TokenType.Map:
                    return true;
                default:
                    return false;
            }
        }
    
        public static bool IsOperator(TokenType tokenType)
        {
            return tokenType == TokenType.Plus ||
                tokenType == TokenType.Minus;
        }

        public static bool IsAssignmentOperator(TokenType type)
        {
            return type == TokenType.Equals ||
                type == TokenType.PlusEquals ||
                type == TokenType.DivideEquals ||
                type == TokenType.MultiplyEquals ||
                type == TokenType.MinusEquals ||
                type == TokenType.ModulusEquals;
        }
    }
}
