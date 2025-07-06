using System;
using System.Collections.Generic;


namespace TypeGo
{
    public static class FormatChan
    {
        public static BlockData? ProcessChannel(FormatData formatData, Token constToken)
        {
            formatData.Increment();
            Token? nextToken = formatData.GetToken();
            if (nextToken == null) {
                formatData.EndOfFileError(constToken, "ProcessChannel");
                return null;
            }
            BlockData? blockData = FormatDeclarations.DeclarationLoop(formatData, nextToken.Value);
            if (blockData == null) {
                return null;
            }

            SetNodeType(blockData);
            blockData.StartingToken = constToken;

            return blockData;
        }

        static void SetNodeType(BlockData blockData)
        {
            switch (blockData.NodeType)
            {
                case NodeType.Single_Declaration_With_Value:
                case NodeType.Multiple_Declarations_With_Value:
                case NodeType.Multiple_Declarations_Same_Type_With_Value:
                    blockData.NodeType = NodeType.Channel_Declaration_With_Value;
                    break;

                case NodeType.Single_Declaration_No_Value:
                case NodeType.Multiple_Declarations_No_Value:
                case NodeType.Multiple_Declarations_Same_Type_No_Value:
                    blockData.NodeType = NodeType.Channel_Declaration;
                    break;

                default:
                    blockData.NodeType = NodeType.Channel_Declaration;
                    break;
            }
        }
    }
}
