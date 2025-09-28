using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class FormatReturn
    {
        public static BlockData? ProcessOther(FormatData formatData)
        {
            BlockData blockData = new BlockData();

            blockData.NodeType = NodeType.Other;

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);

            return blockData;
        }
    }
}
