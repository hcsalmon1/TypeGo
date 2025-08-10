

namespace TypeGo
{
    public static class ConvertErrReturn
    {
        public static void Process(ConvertData convertData, BlockData blockData, int nestCount)
        {
            convertData.AppendString($"if err != nil {{");
            convertData.NewLineWithTabs(nestCount + 1);
            convertData.AppendString("return ");

            List<Token> tokenList = blockData.Tokens;
            if (tokenList.Count != 0)
            {
                ConvertBlock.PrintTokens(convertData, blockData, nestCount, newLine: false);
            }
            convertData.NewLineWithTabs(nestCount);
            convertData.AppendChar('}');
            convertData.NewLineWithTabs(nestCount);
        }
    }
}
