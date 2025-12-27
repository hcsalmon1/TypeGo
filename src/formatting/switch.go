
package formatting

import . "TypeGo/core"



func ProcessSwitch(formatData *FormatData) BlockData {
	
	formatData.SetErrorFunction("ProcessSwitch"); 
	
	var blockData BlockData
	blockData.NodeType = NodeType.Switch; 
	
	LoopTokensUntilLineEnd(formatData, &blockData, false); 
	return blockData; 
}
