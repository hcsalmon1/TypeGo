using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class FormatSwitch
    {
        public static BlockData? ProcessSwitch(FormatData formatData)
        {
            BlockData blockData = new();
            blockData.NodeType = NodeType.Switch;

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: false);
            return blockData;
        }
    }
}
