

namespace TypeGo
{
    public static class FormatGo
    {
        public static BlockData? ProcessGo(FormatData formatData, ref readonly Token token)
        {
            //const string THIS_FUNCTION = "ProcessDefer";

            BlockData blockData = new BlockData();

            blockData.NodeType = NodeType.Other;

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);

            return blockData;
        }
    }
}
