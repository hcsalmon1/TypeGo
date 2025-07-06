using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class FormatFor
    {
        public static BlockData? ProcessFor(FormatData formatData, Token firstToken)
        {
            const string thisFunctionName = "ProcessFor";

            BlockData blockData = new BlockData();

            blockData.NodeType = NodeType.For_Loop;
            blockData.StartingToken = firstToken;

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                Token? token = formatData.GetToken();

                if (token == null) {
                    formatData.EndOfFileError(token, thisFunctionName);
                    return null;
                }

                if (token.Value.Type == TokenType.LeftBrace) {
                    break;
                }
                blockData.Tokens.Add(token.Value);

                formatData.Increment();
            }

            if (formatData.ExpectType(TokenType.LeftBrace, "missing '{' after for loop", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();
            formatData.IncrementIfNewLine();
            CodeBlock? forBody = FormatBody.FillBody(formatData);
            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return null;
            }
            if (forBody == null) {
                formatData.Result = FormatResult.Internal_Error;
                formatData.ErrorDetail = "forBody is null in ProcessFor but no prior error";
                formatData.ErrorToken = firstToken;
                return null;
            }
            if (formatData.ExpectType(TokenType.RightBrace, "missing '}' after for loop block", thisFunctionName) == false) {
                return null;
            }
            formatData.Increment();
            formatData.IncrementIfNewLine();

            blockData.Block = forBody;

            return blockData;
        }
    }
}
