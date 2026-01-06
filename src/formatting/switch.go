
package formatting

import . "TypeGo/core"



func ProcessSwitch(formatData *FormatData) BlockData {
	
	formatData.SetErrorFunction("ProcessSwitch"); 
	
	var blockData BlockData
	blockData.NodeType = NodeType.Switch; 
	AddIfCondition(formatData, &blockData); 
	
	if formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after switch") == false {
		return blockData; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var switchStatementBlock CodeBlock  = FillSwitchBody(formatData); 
	if formatData.IsError() {
		return blockData; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after switch") == false {
		return blockData; 
	}
	formatData.Increment(); 
	
	blockData.Block = &switchStatementBlock; 
	
	return blockData; 
}
