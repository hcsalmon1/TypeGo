using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class FormatComment
    {
        public static BlockData? ProcessComment(FormatData formatData)
        {
            BlockData blockData = new();

            blockData.NodeType = NodeType.Comment;

            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
            return blockData;
        }
    }
}
