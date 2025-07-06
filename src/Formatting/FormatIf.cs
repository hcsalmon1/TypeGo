using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class FormatIf
    {
        public static BlockData? ProcessIf(FormatData formatData, Token token)
        {
            const string thisFunctionName = "ProcessIf";

            BlockData blockData = new BlockData {
                NodeType = NodeType.If_Statement,
                StartingToken = token,
            };

            AddIfCondition(formatData, blockData);

            if (formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after if", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();
            formatData.IncrementIfNewLine();

            CodeBlock? ifStatementBlock = FormatBody.FillBody(formatData);
            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return null;
            }
            if (formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after if", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();

            blockData.Block = ifStatementBlock;

            return blockData;
        }

        static bool IsDeclarationIf(FormatData formatData)
        {
            const string thisFunctionName = "IsDeclarationIf";
            int index = formatData.TokenIndex;

            for (int i = 0; i < formatData.TokenList.Count; i++) {

                Token? tempToken = formatData.GetTokenByIndex(index);
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return false;
                }
                Token token = tempToken.Value;
                if (token.Type == TokenType.Semicolon) {
                    return true;
                }
                if (token.Type == TokenType.LeftBrace) {
                    break;
                }
                index += 1;
            }
            return false;
        }

        static void AddIfCondition(FormatData formatData, BlockData blockData)
        {
            const string thisFunctionName = "IsDeclarationIf";

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                Token? tempToken = formatData.GetToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }
                Token token = tempToken.Value;
                if (token.Type == TokenType.LeftBrace) {
                    break;
                }
                blockData.Tokens.Add(token);
                formatData.Increment();
            }
        }
    
        public static BlockData? ProcessElse(FormatData formatData, Token token)
        {
            const string thisFunctionName = "ProcessElse";

            int index = formatData.TokenIndex + 1;

            BlockData blockData = new BlockData {
                NodeType = NodeType.Else_Statement,
                StartingToken = token,
            };

            formatData.Increment();
            Token? nextToken = formatData.GetTokenByIndex(index);
            if (nextToken == null) {
                formatData.EndOfFileError(token, thisFunctionName);
                return null;
            }
            if (nextToken.Value.Type == TokenType.If) {
                bool hasDeclaration = IsDeclarationIf(formatData);
                if (hasDeclaration) {
                    formatData.UnsupportedFeatureError(token, thisFunctionName);
                    return null;
                }

                AddIfCondition(formatData, blockData);
            }

            if (formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after else", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();

            CodeBlock? ifStatementBlock = FormatBody.FillBody(formatData);
            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return null;
            }
            if (formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after else", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();

            blockData.Block = ifStatementBlock;

            return blockData;
        }
    }
}
