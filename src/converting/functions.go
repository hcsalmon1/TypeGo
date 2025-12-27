
package converting

import . "TypeGo/core"



func ProcessFunction(convertData *ConvertData, function *Function) {
	
	convertData.NewLineWithTabs(); 
	PrintFunctionNameAndParameters(convertData, function); 
	
	if function.ReturnType != "" {
		convertData.AppendString(function.ReturnType)
		convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
		
		
	}
	convertData.GeneratedCode = append(convertData.GeneratedCode, '{')
	
	
	if function.InnerBlock == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "inner block is null in ProcessFunction"; 
		return; 
	}
	
	convertData.IncrementNestCount()
	ProcessBlock(convertData, function.InnerBlock, 1, false); 
	if convertData.IsError() {
		return 
	}
	convertData.DecrementNestCount()
	convertData.NewLine(); 
	convertData.GeneratedCode = append(convertData.GeneratedCode, '}')
	
	convertData.NewLine(); 
	
}

func PrintFunctionNameAndParameters(convertData *ConvertData, function *Function) {
	switch convertData.MethodType {
	case MethodType.Struct:
	convertData.AppendString("func " +"(self *" +convertData.StructName + ") " +function.Name + "(")
	
	case MethodType.Enum:
	convertData.AppendString("func " +"(self " +convertData.StructName + ") " +function.Name + "(")
	
	default:
	convertData.AppendString("func " +function.Name + "(")
	
		}
	
	
	PrintParameters(convertData, function); 
	convertData.GeneratedCode = append(convertData.GeneratedCode, ')')
	
	convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
	
	
}

func PrintParameters(convertData *ConvertData, function *Function) {
	var parameters []Variable  = function.Parameters; 
	if len(parameters) == 0 {
		return; 
	}
	
	for parameterIndex := 0; parameterIndex < len(parameters); parameterIndex++ {
	
		var parameter Variable  = parameters[parameterIndex]; 
		if parameter.NameToken == nil {
			convertData.ConvertResult = ConvertResult.Internal_Error; 
			convertData.ErrorDetail = "name token is null in PrintParameters"; 
			return; 
		}
		var typeAsText string  = JoinTextInListOfTokens( &parameter.TypeList); 
		if typeAsText == "" {
			convertData.ConvertResult = ConvertResult.Internal_Error; 
			convertData.ErrorDetail = "var type text is null in PrintParameters"; 
			return; 
		}
		if parameterIndex != 0 {
			convertData.GeneratedCode = append(convertData.GeneratedCode, ',')
			
			convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
			
			
		}
		if len(parameter.NameToken) == 0 {
			convertData.ConvertResult = ConvertResult.Internal_Error; 
			convertData.ErrorDetail = "parameter name invalid in PrintParameters"; 
			return; 
		}
		convertData.AppendString(parameter.NameToken[0].Text + " " +typeAsText); 
			}
	
	
}
