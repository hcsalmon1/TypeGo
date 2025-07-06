

namespace TypeGo
{
    public static class ConvertIf
    {
        public static void PrintIfStatement(ConvertData convertData, BlockData blockData, int nestCount) 
        {

            //public NodeType NodeType = NodeType.Invalid;
            //public List<Token> Tokens = new List<Token>();
            //public Token? StartingToken = null;
            //public CodeBlock? Block = null;
            //public List<Variable> Variables = new List<Variable>();

            if (blockData.Tokens == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "tokens is null in PrintIfStatement, ConvertIf";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (blockData.Tokens.Count == 0) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "No Tokens in PrintIfStatement, ConvertIf";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }

            ConvertBlock.PrintTokensNoNL(convertData, blockData, nestCount);
            convertData.AppendChar(' ');
            convertData.AppendChar('{');
            if (blockData.Block != null) {
                ConvertBlock.ProcessBlock(convertData, blockData.Block, nestCount + 1);
            }
            convertData.AppendChar('\r');
            convertData.AddTabs(nestCount);
            convertData.AppendChar('}');

            convertData.lastNodeType = NodeType.If_Statement;

            //convertData.NewLineWithTabs(nestCount);
        }
        public static void PrintElseStatement(ConvertData convertData, BlockData blockData, int nestCount)
        {

            //public NodeType NodeType = NodeType.Invalid;
            //public List<Token> Tokens = new List<Token>();
            //public Token? StartingToken = null;
            //public CodeBlock? Block = null;
            //public List<Variable> Variables = new List<Variable>();

            // Walk backwards deleting until we hit something that's not space/tab/newline


            if (blockData.Tokens == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "No tokens in Else statement";
                return;
            }

            convertData.AppendString(" else ");
            if (blockData.Tokens.Count != 0) {
                ConvertBlock.PrintTokensNoNL(convertData, blockData, nestCount);
            }

            convertData.AppendChar(' ');
            convertData.AppendChar('{');
            if (blockData.Block != null)
            {
                ConvertBlock.ProcessBlock(convertData, blockData.Block, nestCount + 1);
            }
            convertData.AppendChar('\r');
            convertData.AddTabs(nestCount);
            convertData.AppendChar('}');

            convertData.lastNodeType = NodeType.Else_Statement;
            //   convertData.NewLineWithTabs(nestCount);
        }
    }
}
