
package formatting

import . "TypeGo/core"



func ProcessGo(formatData *FormatData, token Token) BlockData {
	
	formatData.SetErrorFunction("ProcessGo"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.NodeType = NodeType.Other; 
	
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}
