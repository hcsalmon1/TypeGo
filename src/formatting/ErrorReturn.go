
package formatting

import . "TypeGo/core"



func ProcessErrReturn(formatData *FormatData, token Token) BlockData {
	
	formatData.SetErrorFunction("ProcessErrReturn"); 
	
	var blockData BlockData
	
	blockData.NodeType = NodeType.Err_Return; 
	formatData.Increment(); 
	
	LoopTokensUntilLineEnd(formatData, &blockData, true); 
	
	return blockData; 
}
