using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class FormatGoto
    {
        public static BlockData? ProcessGoto(FormatData formatData, ref readonly Token token)
        {
            //const string THIS_FUNCTION = "ProcessDefer";

            BlockData blockData = new BlockData();

            blockData.NodeType = NodeType.Other;

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);

            return blockData;
        }
    }
}
