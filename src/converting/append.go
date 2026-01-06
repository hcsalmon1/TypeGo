
package converting

import . "TypeGo/core"



func PrintAppend(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Variables) == 0 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "variables count in PrintAppend"; 
		return; 
	}
	
	var variable Variable  = blockData.Variables[0]; 
	if variable.NameToken[0].Text == "" {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "NameToken is null in PrintAppend"; 
		return; 
	}
	
	var variableName string  = variable.NameToken[0].Text; 
	convertData.AppendString(variableName + " = append(" +variableName + ", "); 
	
	if len(blockData.Tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "tokens count is 0 in PrintAppend"; 
		return; 
	}
	
	for i := 0; i < len(blockData.Tokens); i++ {
	
		convertData.AppendToken(blockData.Tokens[i]); 
		
	}
	
	
	convertData.AppendChar(')'); 
	
}
