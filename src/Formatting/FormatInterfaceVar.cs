

namespace TypeGo
{
    public static class FormatInterface
    {
        public static BlockData? FormatInterfaceVar(FormatData formatData, Token firstToken)
        {
            const string FUNCTION_NAME = "FormatInterfaceVar";

            BlockData blockData = new BlockData {
                Block = null,
                NodeType = NodeType.Single_Declaration_No_Value,
                StartingToken = firstToken,
                Tokens = new List<Token>(),
                Variables = new List<Variable>(),
            };

            Variable variable = new();
            variable.TypeList.Add(firstToken);

            formatData.Increment();
            Token? tempToken = formatData.GetToken();

            if (tempToken == null) {
                formatData.EndOfFileError(firstToken, "FormatInterfaceVar");
                return null;
            }

            if (tempToken.Value.Type != TokenType.LeftBrace) {
                formatData.MissingExpectedTypeError(tempToken, "Missing expected '{' in interface declaration", FUNCTION_NAME);
                return null;
            }

            variable.TypeList.Add(tempToken.Value);

            formatData.Increment();
            tempToken = formatData.GetToken();

            if (tempToken == null) {
                formatData.EndOfFileError(firstToken, "FormatInterfaceVar");
                return null;
            }

            if (tempToken.Value.Type != TokenType.RightBrace) {
                formatData.MissingExpectedTypeError(tempToken, "Missing expected '}' in interface declaration, advanced local interfaces not supported yet", FUNCTION_NAME);
                return null;
            }

            variable.TypeList.Add(tempToken.Value);

            formatData.Increment();

            tempToken = formatData.GetToken();

            if (tempToken == null) {
                formatData.EndOfFileError(firstToken, "FormatInterfaceVar");
                return null;
            }

            if (tempToken.Value.Type != TokenType.Identifier) {
                formatData.MissingExpectedTypeError(tempToken, "Missing expected identifier in interface declaration", FUNCTION_NAME);
                return null;
            }

            variable.NameToken = tempToken.Value;

            formatData.Increment();

            tempToken = formatData.GetToken();

            if (tempToken == null) {
                formatData.EndOfFileError(firstToken, "FormatInterfaceVar");
                return null;
            }

            if (tempToken.Value.Type == TokenType.Equals)
            {
                blockData.NodeType = NodeType.Single_Declaration_With_Value;
                FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
            }

            blockData.Variables.Add(variable);

            return blockData;
        }
    }
}
