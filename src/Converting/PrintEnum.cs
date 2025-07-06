

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

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                for (int i = 0; i < varBlock.Tokens.Count; i++)
                {
                    Token token = varBlock.Tokens[i];

                    convertData.AppendString($"{enumNameText}");
                    convertData.AppendString(token.Text);
                    if (blockIndex == 0) {
                        convertData.AppendString(" = iota");
                    }
                    convertData.NewLineWithTabs(nestCount + 1);
                    break;
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

            if (blockDataList.Count == 0)
            {
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

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++)
            {

                BlockData varBlock = blockDataList[blockIndex];

                for (int i = 0; i < varBlock.Tokens.Count; i++)
                {
                    Token token = varBlock.Tokens[i];

                    convertData.AppendString(token.Text);
                    convertData.AppendString($": {blockIndex},");
                    convertData.NewLineWithTabs(nestCount + 1);
                    break;
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
