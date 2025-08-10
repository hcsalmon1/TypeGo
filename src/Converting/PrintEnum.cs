

namespace TypeGo
{
    public static class PrintEnum
    {
        public static void ProcessEnum(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Tokens == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Tokens was null in ProcessEnum, PrintEnum";
                return;
            }
            if (blockData.Tokens.Count != 1) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Tokens count wasn't 1 in ProcessEnum, PrintEnum";
                return;
            }
            string enumNameText = blockData.Tokens[0].Text;

            convertData.AppendString($"type {enumNameText} int\n\n");


            CodeBlock? enumVariableBlock = blockData.Block;
            if (enumVariableBlock == null)
            {
                convertData.AppendChar('\n');
                EndEnum(convertData);
                return;
            }

            if (enumVariableBlock.BlockDataList == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "BlockDataList was null in ProcessEnum, PrintEnum";
                return;
            }

            List<BlockData> blockDataList = enumVariableBlock.BlockDataList;

            if (blockDataList.Count == 0)
            {
                convertData.AppendChar('\n');
                EndEnum(convertData);
                return;
            }

            convertData.AppendString("const (\n\t");

            int enumCount = 0;

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                int tokenCount = varBlock.Tokens.Count;

                if (tokenCount == 0) {
                    continue;
                }

                if (tokenCount == 1 || tokenCount == 2) {

                    Token onlyToken = varBlock.Tokens[0];

                    convertData.AppendString(enumNameText);
                    convertData.AppendString(onlyToken.Text);
                    convertData.AppendString($" = {enumCount}");
                    enumCount += 1;
                    convertData.NewLineWithTabs(nestCount + 1);
                    continue;
                }

                if (tokenCount == 3)
                {

                    Token firstToken = varBlock.Tokens[0];
                    Token secondToken = varBlock.Tokens[1];
                    Token thirdToken = varBlock.Tokens[2];

                    if (secondToken.Text != "=")
                    {
                        convertData.MissingTypeError(secondToken, "missing expect '=' in enumstruct", "ProcessEnumstruct");
                        return;
                    }

                    if (!int.TryParse(thirdToken.Text, out int value))
                    {
                        convertData.MissingTypeError(secondToken, "missing expected integer in enumstruct", "ProcessEnumstruct");
                        return;
                    }

                    convertData.AppendString(enumNameText);
                    convertData.AppendString(firstToken.Text);
                    convertData.AppendString($" = {value}");
                    enumCount = value + 1;
                    convertData.NewLineWithTabs(nestCount + 1);
                }
            }

            convertData.AppendChar('\r');
            EndEnum(convertData);

            string enumNameLower = enumNameText.ToLower();
            convertData.AppendString($"func {enumNameText}ToString({enumNameLower} int) string {{\n\t");
            convertData.AppendString($"switch {enumNameLower} {{\n\t");

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                for (int i = 0; i < varBlock.Tokens.Count; i++)
                {
                    Token token = varBlock.Tokens[i];

                    convertData.AppendString($"case {enumNameText}{token.Text}:\n\t\t");
                    convertData.AppendString($"return \"{token.Text}\"\n\t");
                    break;
                }
            }
            convertData.AppendString($"default:\n\t\t");
            convertData.AppendString($"return \"Unknown\"\n\t");
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('\n');
        }

        public static void ProcessEnumstruct(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Tokens == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Tokens null in ProcessEnumstruct";
                return;
            }
            if (blockData.Tokens.Count != 1) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Tokens count != 1 in ProcessEnumstruct";
                return;
            }
            string enumNameText = blockData.Tokens[0].Text;

            convertData.AppendString($"var {enumNameText} = struct {{");


            CodeBlock? enumVariableBlock = blockData.Block;
            if (enumVariableBlock == null)
            {
                convertData.AppendChar('\n');
                EndEnum(convertData);
                return;
            }

            if (enumVariableBlock.BlockDataList == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "BlockDataList is null in ProcessEnumstruct";
                return;
            }

            List<BlockData> blockDataList = enumVariableBlock.BlockDataList;

            if (blockDataList.Count == 0) {
                convertData.AppendChar('\n');
                EndEnum(convertData);
                return;
            }

            convertData.AppendString("\n\t");

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                for (int i = 0; i < varBlock.Tokens.Count; i++)
                {
                    Token token = varBlock.Tokens[i];

                    convertData.AppendString(token.Text);
                    convertData.AppendString(" int");
                    convertData.NewLineWithTabs(nestCount + 1);
                    break;
                }
            }

            convertData.AppendChar('\r');
            convertData.AppendChar('}');
            convertData.AppendChar('{');
            convertData.NewLineWithTabs(nestCount + 1);

            int enumCount = 0;

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                int tokenCount = varBlock.Tokens.Count;

                if (tokenCount == 0) {
                    continue;
                }

                if (tokenCount == 1 || tokenCount == 2) {

                    Token onlyToken = varBlock.Tokens[0];

                    convertData.AppendString(onlyToken.Text);
                    convertData.AppendString($": {enumCount},");
                    enumCount += 1;
                    convertData.NewLineWithTabs(nestCount + 1);
                    continue;
                }

                if (tokenCount == 3) {

                    Token firstToken = varBlock.Tokens[0];
                    Token secondToken = varBlock.Tokens[1];
                    Token thirdToken = varBlock.Tokens[2];

                    if (secondToken.Text != "=") {
                        convertData.MissingTypeError(secondToken, "missing expect '=' in enumstruct", "ProcessEnumstruct");
                        return;
                    }

                    if (!int.TryParse(thirdToken.Text, out int value))
                    {
                        convertData.MissingTypeError(secondToken, "missing expected integer in enumstruct", "ProcessEnumstruct");
                        return;
                    }

                    convertData.AppendString(firstToken.Text);
                    convertData.AppendString($": {value},");
                    enumCount = value + 1;
                    convertData.NewLineWithTabs(nestCount + 1);
                }
            }

            convertData.AppendChar('\r');
            EndEnumStruct(convertData);

            string enumNameLower = enumNameText.ToLower();
            convertData.AppendString($"func {enumNameText}ToString({enumNameLower} int) string {{\n\t");
            convertData.AppendString($"switch {enumNameLower} {{\n\t");

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                for (int i = 0; i < varBlock.Tokens.Count; i++)
                {
                    Token token = varBlock.Tokens[i];

                    convertData.AppendString($"case {enumNameText}.{token.Text}:\n\t\t");
                    convertData.AppendString($"return \"{token.Text}\"\n\t");
                    break;
                }
            }
            convertData.AppendString($"default:\n\t\t");
            convertData.AppendString($"return \"Unknown\"\n");
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('\r');
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('\n');
        }

        private static void EndEnum(ConvertData convertData)
        {
            convertData.AppendChar(')');
            convertData.AppendChar('\n');
            convertData.AppendChar('\n');
        }

        private static void EndEnumStruct(ConvertData convertData)
        {
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('\n');
        }
    }
}
