
package converting

import (
."TypeGo/core"
"fmt"
)


func InfiniteWhileCheck(whileCount *int, MAX int) bool {
	* whileCount += 1
	if * whileCount >= MAX {
		return true
	}
	return false
}

func ShouldReturn(convertData *ConvertData, whileCount *int, WHILE_CAP int) bool {
	
	if InfiniteWhileCheck(whileCount, WHILE_CAP) {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "InfiniteWhileLoop in Should Return"; 
		return true; 
	}
	if convertData.IsError() {
		return true; 
	}
	return false; 
}

func PrintCodeFormat(convertData *ConvertData) {
	convertData.SetErrorFunction("LoopTokens")
	
	var globalBlock CodeBlock  = convertData.CodeFormat.GlobalBlock; 
	var nest_count int  = 0
	var is_global bool  = true
	ProcessBlock(convertData, &globalBlock, nest_count, is_global); 
	if convertData.IsError() {
		return 
	}
	
	var functions []Function  = convertData.CodeFormat.Functions; 
	for functionIndex := 0; functionIndex < len(functions); functionIndex++ {
	
		ProcessFunction(convertData, &functions[functionIndex]); 
		
		if convertData.IsError() {
			return; 
		}
		
	}
	
	
}

func ConvertToGo(codeFormat *CodeFormat, convertResult *IntConvertResult, code string) string {
	//fmt.Printf("\t%sConverting:%s\t\t\t", GREY_TEXT, RESET_TEXT);
	var convertData ConvertData  = ConvertData{
		CodeFormat:* codeFormat, 
		ErrorDetail:"", 
		ConvertResult:ConvertResult.Ok, 
		GeneratedCode:make([]byte, 0), 
	}; 
	PrintCodeFormat( &convertData); 
	if convertData.IsError() {
		fmt.Printf("Error %s%s%s\n", RED_TEXT, convertData.ConvertResult.ToString(), RESET_TEXT); 
		PrintConvertError( &convertData, code); 
		return ""; 
	}
	
	//fmt.Printf("%sDone%s\n", CYAN_TEXT, RESET_TEXT);
	
	return string(convertData.GeneratedCode); 
}
