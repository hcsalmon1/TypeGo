using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class ConvertSwitch
    {
        public static void Convert(ConvertData convertData, BlockData blockData, int nestCount, bool newLine)
        {
            if (blockData.Tokens.Count == 0)
            {
                convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData", "PrintTokens");
                return;
            }
            TokenType lastType = TokenType.NA;
            bool addedSpace = false;

            for (int i = 0; i < blockData.Tokens.Count; i++)
            {

                Token token = blockData.Tokens[i];

                if (token.Type == TokenType.NewLine)
                {
                    convertData.NewLineWithTabs(nestCount);
                    lastType = token.Type;
                    continue;
                }
                ConvertBlock.AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace);
                HandleToken(convertData, token, ref nestCount);
                addedSpace = ConvertBlock.AddSpaceAfter(convertData, token.Type, lastType, i);
                lastType = token.Type;
            }

            if (lastType != TokenType.NewLine)
            {
                convertData.NewLineWithTabs(nestCount);
            }
            if (newLine == true)
            {
                convertData.NewLineWithTabs(nestCount);
            }

            static void HandleToken(ConvertData convertData, Token token, ref int nestCount)
            {
                if (token.Type == TokenType.Semicolon)
                {

                    int codeLength = convertData.GeneratedCode.Length;
                    char lastChar = convertData.GeneratedCode[codeLength - 1];
                    if (lastChar == ' ')
                    {
                        convertData.GeneratedCode[codeLength - 1] = ';';
                        return;
                    }
                }
                if (token.Text == "switch") {
                    nestCount++;
                }
                if (token.Text == "}") {
                    nestCount--;
                    convertData.AppendChar('\r');
                    convertData.AddTabs(nestCount);
                }
                convertData.AppendString(token.Text);
            }
        }
    }
}
