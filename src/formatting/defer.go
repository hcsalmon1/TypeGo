
package formatting

import . "TypeGo/core"



func ProcessDefer(formatData *FormatData, token *Token) BlockData {
	
	formatData.SetErrorFunction("ProcessDefer"); 
	
	var blockData BlockData
	
	blockData.NodeType = NodeType.Other; 
	
	formatData.SetErrorProcess("Loop tokens until line end, defer"); 
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}
