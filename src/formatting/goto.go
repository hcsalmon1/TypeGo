
package formatting

import . "TypeGo/core"



func ProcessGoto(formatData *FormatData, token Token) BlockData {
	formatData.SetErrorFunction("ProcessGoto"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.NodeType = NodeType.Other; 
	
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}
