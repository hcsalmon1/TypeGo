
package converting

import . "TypeGo/core"



func ProcessStruct(convertData *ConvertData, blockData *BlockData, nestCount int) {
	if blockData.Tokens == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens is null in ProcessStruct"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(blockData.Tokens) != 1 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Token count is 1 in ProcessStruct"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	var structNameText string  = blockData.Tokens[0].Text; 
	
	convertData.AppendString("type " +structNameText + " struct {"); 
	convertData.NewLineWithTabs(); 
	
	var structVariableBlock *CodeBlock  = blockData.Block; 
	if structVariableBlock == nil {
		convertData.AppendChar('\n'); 
		EndStruct(convertData); 
		return; 
	}
	
	if structVariableBlock.BlockDataList == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "BlockDataList is null in ProcessStruct"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	
	var blockDataList []BlockData  = structVariableBlock.BlockDataList; 
	
	if len(blockDataList) == 0 {
		convertData.AppendChar('\n'); 
		EndStruct(convertData); 
		return; 
	}
	
	var varNames []string  = make([]string, 0); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		if varBlock.NodeType != NodeType.Single_Declaration_No_Value {
			PrintOther(convertData, &varBlock, nestCount); 
			if convertData.IsError() {
				return 
			}
			continue 
			
			
		}
		
		for i := 0; i < len(varBlock.Variables); i++ {
		
			var variable Variable  = varBlock.Variables[i]; 
			if variable.NameToken == nil {
				continue 
				
				
			}
			if len(variable.NameToken) == 0 {
				continue 
				
				
			}
			varNames = append(varNames, variable.NameToken[0].Text)
			
			var varTypeAsText string  = JoinTextInListOfTokens( &variable.TypeList); 
			if varTypeAsText == "" {
				continue 
				
				
			}
			convertData.AppendString(variable.NameToken[0].Text + " " +varTypeAsText); 
			convertData.NewLineWithTabs(); 
			
		}
		
		
	}
	
	
	convertData.AppendChar('\r'); 
	EndStruct(convertData); 
	PrintStructMethods(convertData, structVariableBlock, nestCount, &varNames, structNameText); 
	
}

func PrintStructMethods(convertData *ConvertData, structBlock *CodeBlock, nestCount int, varName *[]string, structName string) {
	
	var functions []Function  = structBlock.MethodList; 
	if functions == nil {
		return; 
	}
	if len(functions) == 0 {
		return; 
	}
	
	convertData.MethodType = MethodType.Struct; 
	convertData.StructName = structName; 
	
	for i := 0; i < len(functions); i++ {
	
		ProcessFunction(convertData, &functions[i]); 
		
	}
	
	convertData.MethodType = MethodType.None; 
	convertData.MethodVarNames = nil
	
}

func PrintOther(convertData *ConvertData, varBlock *BlockData, nestCount int) {
	
	if varBlock.NodeType == NodeType.Single_Declaration_With_Value {
		convertData.ErrorToken = varBlock.StartingToken
		convertData.ErrorDetail = "Can't set a value in a struct definition"
		convertData.ConvertResult = ConvertResult.Unexpected_Type
		return 
	}
	
	if varBlock.NodeType != NodeType.Other {
		return; 
	}
	
	if len(varBlock.Tokens) == 0 {
		return; 
	}
	
	for i := 0; i < len(varBlock.Tokens); i++ {
	
		var token Token  = varBlock.Tokens[i]; 
		
		if token.Type == TokenType.NewLine {
			continue 
			
			
		}
		convertData.AppendString(varBlock.Tokens[i].Text); 
		
	}
	
	
}

func EndStruct(convertData *ConvertData) {
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	
}
