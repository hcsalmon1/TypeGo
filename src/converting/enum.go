
package converting

import (
."TypeGo/core"
"strconv"
"strings"
)


func ProcessEnum(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if blockData.Tokens == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens was null in ProcessEnum, PrintEnum"; 
		return; 
	}
	if len(blockData.Tokens) != 1 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens count wasn't 1 in ProcessEnum, PrintEnum"; 
		return; 
	}
	var enumNameText string  = blockData.Tokens[0].Text; 
	
	convertData.AppendString("type " +enumNameText + " int\n\n"); 
	
	var enumVariableBlock *CodeBlock  = blockData.Block; 
	if enumVariableBlock == nil {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	if enumVariableBlock.BlockDataList == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "BlockDataList was null in ProcessEnum, PrintEnum"; 
		return; 
	}
	
	var blockDataList []BlockData  = enumVariableBlock.BlockDataList; 
	
	if len(blockDataList) == 0 {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	convertData.AppendString("const (\n\t"); 
	
	var enumCount int  = 0; 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		var tokenCount int  = len(varBlock.Tokens); 
		
		if tokenCount == 0 {
			continue 
			
			
		}
		
		if tokenCount == 1 || tokenCount == 2 {
			var onlyToken Token  = varBlock.Tokens[0]; 
			
			convertData.AppendString(enumNameText); 
			convertData.AppendString(onlyToken.Text); 
			convertData.AppendString(" = " +strconv.Itoa(enumCount)); 
			enumCount += 1; 
			convertData.NewLineWithTabs(); 
			continue 
			
			
		}
		
		if tokenCount == 3 {
			var firstToken Token  = varBlock.Tokens[0]; 
			var secondToken Token  = varBlock.Tokens[1]; 
			//Token thirdToken = varBlock.Tokens[2];
			
			if secondToken.Text != "=" {
				convertData.MissingTypeError(secondToken, "missing expect '=' in enumstruct"); 
				return; 
			}
			
			var value int; var err error
			value, err  = strconv.Atoi(secondToken.Text)
			
			if err != nil {
			convertData.MissingTypeError(secondToken, "missing expected integer in enumstruct"); 
			return; 
			}
			
			
			convertData.AppendString(enumNameText); 
			convertData.AppendString(firstToken.Text); 
			convertData.AppendString(" = " +strconv.Itoa(value))
			enumCount = value + 1; 
			convertData.NewLineWithTabs(); 
			
		}
			}
	
	
	convertData.AppendChar('\r'); 
	EndEnum(convertData); 
	
	var enumNameLower string  = strings.ToLower(enumNameText); 
	convertData.AppendString("func " +enumNameText + "ToString(" +enumNameLower + " int) string {\n\t"); 
	convertData.AppendString("switch self {\n\t"); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		for i := 0; i < len(varBlock.Tokens); i++ {
		
			var token Token  = varBlock.Tokens[i]; 
			
			convertData.AppendString("case " +enumNameText + token.Text + ":\n\t\t"); 
			convertData.AppendString("return \"" +token.Text + "\"\n\t"); 
			break
					}
		
			}
	
	convertData.AppendString("default:\n\t\t"); 
	convertData.AppendString("return \"Unknown\"\n\t"); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	
}

func ProcessEnumstruct(convertData *ConvertData, blockData *BlockData, nestCount int) {
	if blockData.Tokens == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens null in ProcessEnumstruct"; 
		return; 
	}
	if len(blockData.Tokens) != 1 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens count != 1 in ProcessEnumstruct"; 
		return; 
	}
	var enumNameText string  = blockData.Tokens[0].Text; 
	
	var alias_name string  = "Int" +enumNameText
	
	
	convertData.AppendString("type " +alias_name + " int\n")
	
	convertData.AppendString("var " +enumNameText + " = struct {"); 
	
	
	var enumVariableBlock *CodeBlock  = blockData.Block; 
	if enumVariableBlock == nil {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	if enumVariableBlock.BlockDataList == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "BlockDataList is null in ProcessEnumstruct"; 
		return; 
	}
	
	var blockDataList []BlockData  = enumVariableBlock.BlockDataList; 
	
	if len(blockDataList) == 0 {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	convertData.AppendString("\n\t"); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		for i := 0; i < len(varBlock.Tokens); i++ {
		
			var token Token  = varBlock.Tokens[i]; 
			
			convertData.AppendString(token.Text); 
			convertData.AppendString(" " +alias_name); 
			convertData.NewLineWithTabs(); 
			break
					}
		
			}
	
	
	convertData.AppendChar('\r'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('{'); 
	convertData.NewLineWithTabs(); 
	
	var enumCount int  = 0; 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		var tokenCount int  = len(varBlock.Tokens); 
		
		if tokenCount == 0 {
			continue 
			
			
		}
		
		if tokenCount == 1 || tokenCount == 2 {
			var onlyToken Token  = varBlock.Tokens[0]; 
			
			convertData.AppendString(onlyToken.Text); 
			convertData.AppendString(": " +strconv.Itoa(enumCount)+",")
			enumCount += 1; 
			convertData.NewLineWithTabs(); 
			continue 
			
			
		}
		
		if tokenCount == 3 {
			var firstToken Token  = varBlock.Tokens[0]; 
			var secondToken Token  = varBlock.Tokens[1]; 
			var thirdToken Token  = varBlock.Tokens[2]
			
			
			if secondToken.Text != "=" {
				convertData.MissingTypeError(secondToken, "missing expect '=' in enumstruct"); 
				return; 
			}
			
			
			var value int; var err error
			value, err  = strconv.Atoi(thirdToken.Text)
			
			if err != nil {
			convertData.MissingTypeError(thirdToken, "missing expected integer in enumstruct"); 
			return; 
			}
			
			
			convertData.AppendString(firstToken.Text); 
			convertData.AppendString(": " +strconv.Itoa(value)+",")
			enumCount = value + 1; 
			convertData.NewLineWithTabs(); 
			
		}
			}
	
	
	convertData.AppendChar('\r'); 
	EndEnumStruct(convertData); 
	
	convertData.AppendString("func (self " +alias_name + ") ToString() string {\n\t"); 
	convertData.AppendString("switch self {\n\t"); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		for i := 0; i < len(varBlock.Tokens); i++ {
		
			var token Token  = varBlock.Tokens[i]; 
			
			convertData.AppendString("case " +enumNameText + "." +token.Text + ":\n\t\t"); 
			convertData.AppendString("return \"" +token.Text + "\"\n\t"); 
			break
					}
		
			}
	
	convertData.AppendString("default:\n\t\t"); 
	convertData.AppendString("return \"Unknown\"\n"); 
	convertData.AppendChar('\t'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\r'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	PrintEnumMethods(convertData, enumVariableBlock, nestCount, alias_name)
	
}

func ProcessEnumstructWithAlias(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if blockData.Tokens == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens null in ProcessEnumstruct"; 
		return; 
	}
	if len(blockData.Tokens) != 1 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Tokens count != 1 in ProcessEnumstruct"; 
		return; 
	}
	var enumNameText string  = blockData.Tokens[0].Text; 
	
	if blockData.VarName == "" {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "Alias name is null in enumstruct"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	
	var var_type string  = blockData.VarName; 
	
	convertData.AppendString("type " +var_type + " int\n"); 
	
	convertData.AppendString("var " +enumNameText + " = struct {"); 
	
	
	var enumVariableBlock *CodeBlock  = blockData.Block; 
	if enumVariableBlock == nil {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	if enumVariableBlock.BlockDataList == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "BlockDataList is null in ProcessEnumstruct"; 
		return; 
	}
	
	var blockDataList []BlockData  = enumVariableBlock.BlockDataList; 
	
	if len(blockDataList) == 0 {
		convertData.AppendChar('\n'); 
		EndEnum(convertData); 
		return; 
	}
	
	convertData.AppendString("\n\t"); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		for i := 0; i < len(varBlock.Tokens); i++ {
		
			var token Token  = varBlock.Tokens[i]; 
			
			convertData.AppendString(token.Text); 
			convertData.AppendString(" " +var_type); 
			convertData.NewLineWithTabs(); 
			break
					}
		
			}
	
	
	convertData.AppendChar('\r'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('{'); 
	convertData.NewLineWithTabs(); 
	
	var enumCount int  = 0; 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		var tokenCount int  = len(varBlock.Tokens); 
		
		if tokenCount == 0 {
			continue 
			
			
		}
		
		if tokenCount == 1 || tokenCount == 2 {
			var onlyToken Token  = varBlock.Tokens[0]; 
			
			convertData.AppendString(onlyToken.Text); 
			convertData.AppendString(": " +strconv.Itoa(enumCount)+",")
			enumCount += 1; 
			convertData.NewLineWithTabs(); 
			continue 
			
			
		}
		
		if tokenCount == 3 {
			var firstToken Token  = varBlock.Tokens[0]; 
			var secondToken Token  = varBlock.Tokens[1]; 
			var thirdToken Token  = varBlock.Tokens[2]; 
			
			if secondToken.Text != "=" {
				convertData.MissingTypeError(secondToken, "missing expect '=' in enumstruct"); 
				return; 
			}
			
			var value int; var err error
			value, err  = strconv.Atoi(thirdToken.Text)
			
			if err != nil {
			convertData.MissingTypeError(thirdToken, "missing expected integer in enumstruct"); 
			return; 
			}
			
			
			convertData.AppendString(firstToken.Text); 
			convertData.AppendString(": " +strconv.Itoa(value)+",")
			enumCount = value + 1; 
			convertData.NewLineWithTabs(); 
			
		}
			}
	
	
	convertData.AppendChar('\r'); 
	EndEnumStruct(convertData); 
	
	convertData.AppendString("func (self " +var_type + ") ToString() string {\n\t"); 
	convertData.AppendString("switch self {\n\t"); 
	
	for blockIndex := 0; blockIndex < len(blockDataList); blockIndex++ {
	
		var varBlock BlockData  = blockDataList[blockIndex]; 
		
		for i := 0; i < len(varBlock.Tokens); i++ {
		
			var token Token  = varBlock.Tokens[i]; 
			
			convertData.AppendString("case " +enumNameText + "." +token.Text + ":\n\t\t"); 
			convertData.AppendString("return \"" +token.Text + "\"\n\t"); 
			break
					}
		
			}
	
	convertData.AppendString("default:\n\t\t"); 
	convertData.AppendString("return \"Unknown\"\n"); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\r'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	PrintEnumMethods(convertData, enumVariableBlock, nestCount, var_type)
	
}

func PrintEnumMethods(convertData *ConvertData, enumBlock *CodeBlock, nestCount int, enumName string) {
	
	var functions []Function  = enumBlock.MethodList; 
	if functions == nil {
		return; 
	}
	if len(functions) == 0 {
		return; 
	}
	
	convertData.MethodType = MethodType.Enum; 
	convertData.StructName = enumName; 
	
	for i := 0; i < len(functions); i++ {
	
		ProcessFunction(convertData, &functions[i]); 
			}
	
	convertData.MethodType = MethodType.None; 
	convertData.MethodVarNames = nil
	
}

func EndEnum(convertData *ConvertData) {
	convertData.AppendChar(')'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	
}

func EndEnumStruct(convertData *ConvertData) {
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	convertData.AppendChar('\n'); 
	
}
