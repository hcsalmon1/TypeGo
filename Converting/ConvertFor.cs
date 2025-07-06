

namespace TypeGo
{
    public static class ConvertFor
    {
        public static void PrintFor(ConvertData convertData, BlockData blockData, int nestCount)
        {
            const string thisFunctionName = "PrintFor";

            ConvertBlock.PrintTokensNoNL(convertData, blockData, nestCount);
            if (convertData.IsError()) {
                convertData.AddTrace(thisFunctionName);
                return;
            }
            convertData.AppendChar(' ');
            convertData.AppendChar('{');
            convertData.NewLineWithTabs(nestCount + 1);

            CodeBlock? forBlock = blockData.Block;
            if (forBlock != null) {
                ConvertBlock.ProcessBlock(convertData, forBlock, nestCount + 1);
            }
            convertData.AppendChar('\r');
            convertData.AddTabs(nestCount);
            convertData.AppendChar('}');
            convertData.NewLineWithTabs(nestCount);
        }
    }
}
