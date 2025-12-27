
package formatting

import . "TypeGo/core"



func ProcessReturn(formatData *FormatData) BlockData {
	
	formatData.SetErrorFunction("ProcessReturn"); 
	
	var blockData BlockData
	
	blockData.NodeType = NodeType.Return; 
	
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}

func ProcessOther(formatData *FormatData) BlockData {
	
	formatData.SetErrorFunction("ProcessOther"); 
	
	var blockData BlockData
	
	blockData.NodeType = NodeType.Other; 
	
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}
