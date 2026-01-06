
package converting

import . "TypeGo/core"



func PrintErrorCheckStatement(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	convertData.AppendString("if err != nil {"); 
	if blockData.Block != nil {
		ProcessBlock(convertData, blockData.Block, nestCount + 1, false); 
		
	}
	
	convertData.NewLineWithTabs(); 
	convertData.AppendChar('}'); 
	
	convertData.LastNodeType = NodeType.If_Statement; 
	
	//convertData.NewLineWithTabs(nestCount);
	
}

func PrintIfStatement(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	
	if blockData.Tokens == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "tokens is null in PrintIfStatement, ConvertIf"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(blockData.Tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "No Tokens in PrintIfStatement, ConvertIf"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	
	PrintTokensNoNL(convertData, blockData, nestCount); 
	convertData.AppendChar(' '); 
	convertData.AppendChar('{'); 
	convertData.IncrementNestCount()
	if blockData.Block != nil {
		ProcessBlock(convertData, blockData.Block, nestCount + 1, false); 
		
	}
	convertData.DecrementNestCount()
	
	if convertData.WasNewLine() == false {
		convertData.NewLineWithTabs(); 
		
	}
	convertData.AppendChar('}'); 
	
	convertData.LastNodeType = NodeType.If_Statement; 
	
	
}

func PrintElseStatement(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	//public NodeType NodeType = NodeType.Invalid;
	//public List<Token> Tokens = new List<Token>();
	//public Token? StartingToken = null;
	//public CodeBlock? Block = null;
	//public List<Variable> Variables = new List<Variable>();
	
	// Walk backwards deleting until we hit something that's not space/tab/newline
	
	convertData.AppendString(" else "); 
	if len(blockData.Tokens) != 0 {
		PrintTokensNoNL(convertData, blockData, nestCount); 
		
	}
	
	convertData.AppendChar(' '); 
	convertData.AppendChar('{'); 
	convertData.IncrementNestCount()
	if blockData.Block != nil {
		ProcessBlock(convertData, blockData.Block, nestCount + 1, false); 
		
	}
	convertData.DecrementNestCount()
	convertData.NewLineWithTabs(); 
	convertData.AppendChar('}'); 
	
	convertData.LastNodeType = NodeType.Else_Statement; 
	//   convertData.NewLineWithTabs(nestCount);
	
}
