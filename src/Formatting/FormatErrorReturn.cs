

namespace TypeGo
{
    public static class FormatErrorReturn
    {
        public static BlockData? ProcessErrReturn(FormatData formatData, in Token token)
        {
            BlockData blockData = new BlockData();

            blockData.NodeType = NodeType.Err_Return;
            formatData.Increment();

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);

            return blockData;
        }
    }
}
