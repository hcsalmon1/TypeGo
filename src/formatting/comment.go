
package formatting

import . "TypeGo/core"



func ProcessComment(format_data *FormatData) BlockData {
	format_data.AddToFunctionLog("ENTER ProcessComment")
	var blockData BlockData
	
	blockData.NodeType = NodeType.Comment; 
	
	LoopTokensUntilLineEnd(format_data, &blockData, true); 
	format_data.AddToFunctionLog("EXIT ProcessComment")
	return blockData; 
}
