using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class FormatIdentifier
    {

        enum IdentifierLoopAction
        {
            Continue,
            Break,
            Declaration,
            Other,
            Append,
            Error,
        }

        public static BlockData? ProcessIdentifier(FormatData formatData, in Token firstToken)
        {
            const string thisFunctionName = "ProcessIdentifier";

            int initialIndex = formatData.TokenIndex;

            BlockData blockData = new BlockData {
                Block = null,
                NodeType = NodeType.Other,
                StartingToken = firstToken,
                Tokens = new List<Token>(),
                Variables = new List<Variable>(),
            };

            TokenType lastTokenType = TokenType.NA;
            IdentifierLoopAction loopAction = IdentifierLoopAction.Continue;

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                Token? tempToken = formatData.GetToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(firstToken, thisFunctionName);
                    return null;
                }
                Token thisToken = tempToken.Value;

                loopAction = IdentifierLoopCode(formatData, thisToken, lastTokenType);
                if (loopAction != IdentifierLoopAction.Continue) {
                    break;
                }

                formatData.Increment();
                lastTokenType = thisToken.Type;
            }

            if (loopAction == IdentifierLoopAction.Declaration) {
                formatData.TokenIndex = initialIndex;
                return FormatDeclarations.DeclarationLoop(formatData, firstToken);
            }
            if (loopAction == IdentifierLoopAction.Other) {
                formatData.TokenIndex = initialIndex;
                FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
            }
            if (loopAction == IdentifierLoopAction.Append) {
                formatData.TokenIndex = initialIndex;
                FillAppend(formatData, blockData);
            }

            return blockData;
        }

        static IdentifierLoopAction IdentifierLoopCode(FormatData formatData, Token thisToken, TokenType lastTokenType)
        {

            if (TokenUtils.IsVarType(thisToken.Type)) {
                return IdentifierLoopAction.Declaration;
            }

            switch (thisToken.Type)
            {
                case TokenType.Identifier: //custom type     Object object = ...

                    if (lastTokenType == TokenType.Identifier) {
                        return IdentifierLoopAction.Declaration;
                    }

                    return IdentifierLoopAction.Continue;

                case TokenType.FullStop:
                    return IdentifierLoopAction.Continue;

                case TokenType.Comma:
                    return IdentifierLoopAction.Continue;

                case TokenType.Equals:
                case TokenType.PlusPlus:
                case TokenType.PlusEquals:
                case TokenType.MinusEquals:
                case TokenType.MultiplyEquals:
                case TokenType.DivideEquals:
                case TokenType.ModulusEquals:
                case TokenType.LeftParenthesis:
                case TokenType.RightParenthesis:
                case TokenType.LeftBrace:
                case TokenType.LeftSquareBracket:
                case TokenType.RightBrace:
                case TokenType.NewLine:
                case TokenType.Channel_Setter:
                case TokenType.Colon:
                    return IdentifierLoopAction.Other;

                case TokenType.ColonEquals:
                    formatData.UnexpectedTypeError(thisToken, "':=' unsupported outside of for loops, write: 'int value = 10' and not 'value := 10'", "IdentifierLoopCode");
                    return IdentifierLoopAction.Error;

                case TokenType.Append:
                    return IdentifierLoopAction.Append;



                default:
                    formatData.UnexpectedTypeError(thisToken, $"type: {thisToken.Type}", "IdentifierLoopCode");
                    return IdentifierLoopAction.Error;
            }
        }

        static void FillAppend(FormatData formatData, BlockData blockData)
        {
            const string FUNCTION_NAME = "FillAppend";

            blockData.NodeType = NodeType.Append;

            Token? tempToken = null;

            int startingIndex = formatData.TokenIndex;
            int? appendIndex = null;

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                tempToken = formatData.GetToken();

                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, FUNCTION_NAME);
                    return;
                }

                if (tempToken.Value.Type == TokenType.Append) {
                    appendIndex = formatData.TokenIndex;
                    break;
                }
                formatData.Increment();
            }

            if (appendIndex == null) {
                formatData.UnexpectedTypeError(tempToken, "Missing 'append'", FUNCTION_NAME);
                return;
            }

            Variable nameVariable = new Variable();
            nameVariable.TypeList = new List<Token>();

            StringBuilder nameTextBuilder = new();

            for (int i = startingIndex; i < appendIndex.Value; i++) {

                if (i + 1 == appendIndex.Value)
                {
                    break;
                }

                tempToken = formatData.GetTokenByIndex(i);

                nameTextBuilder.Append(tempToken.Value.Text);


            }

            nameVariable.NameToken.Add(new Token(nameTextBuilder.ToString(), TokenType.Identifier, 0, 0));

            blockData.Variables.Add(nameVariable);

            formatData.TokenIndex = appendIndex.Value + 1;
            formatData.ExpectType(TokenType.LeftParenthesis, "expecting a '(' after append", FUNCTION_NAME);
            formatData.Increment();

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                tempToken = formatData.GetToken();

                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, FUNCTION_NAME);
                    return;
                }

                if (tempToken.Value.Type == TokenType.RightParenthesis) {
                    break;
                }

                blockData.Tokens.Add(tempToken.Value);

                formatData.Increment();
            }

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {

                tempToken = formatData.GetToken();

                if (tempToken == null)
                {
                    formatData.EndOfFileError(tempToken, FUNCTION_NAME);
                    return;
                }

                if (tempToken.Value.Type == TokenType.Semicolon) {
                    break;
                }
                if (tempToken.Value.Type == TokenType.NewLine) {
                    break;
                }

                formatData.Increment();
            }
        }
    }
}
