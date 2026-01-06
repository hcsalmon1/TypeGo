
package formatting

import (
."TypeGo/core"
"fmt"
)


func FormatCode(tokenList []Token, result *IntFormatResult, code string) CodeFormat {
	
	//fmt.Printf("\t%sFormatting:%s\t\t\t", GREY_TEXT, RESET_TEXT);
	var codeFormat CodeFormat
	var globalBlock CodeBlock
	var functionList []Function  = make([]Function, 0); 
	var formatData FormatData  = FormatData{
		TokenList:tokenList, 
		ErrorToken:EmptyToken(), 
		ErrorDetail:"", 
		Result:FormatResult.Ok, 
		ErrorFunction:"", 
		FunctionLog:make([]string, 0), 
		ProcessLog:make([]string, 0), 
	}; 
	formatData.ErrorFunction = "FormatCode"; 
	
	if len(tokenList) == 0 {
		* result = FormatResult.NoTokens; 
		return codeFormat; 
	}
	
	for formatData.IndexInBounds() == true {
	
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			* result = formatData.Result
			return codeFormat; 
		}
		
		var previousIndex int  = formatData.TokenIndex; 
		var blockData BlockData
		
		switch token.Type {
			
			
			case TokenType.Chan:
				blockData = ProcessChannel( &formatData, token); 
				globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData); 
				
				
			case TokenType.Const:
				ProcessConstant( &formatData, &globalBlock, token); 
				
				
			case TokenType.Int, TokenType.Int8, TokenType.Int16, TokenType.Int32, TokenType.Int64, TokenType.Uint, 
			TokenType.Uint8, TokenType.Uint16, TokenType.Uint32, TokenType.Uint64, TokenType.Float32, TokenType.Float64, 
			TokenType.String, TokenType.Byte, TokenType.Rune, TokenType.Bool, TokenType.LeftSquareBracket, TokenType.Error, 
			TokenType.Multiply:
				blockData = ProcessDeclaration( &formatData, token); 
				globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData); 
				
				
			case TokenType.Type:
				blockData = ProcessOther( &formatData); 
				globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData); 
				
				
			case TokenType.Package:
				ProcessPackage( &formatData, &globalBlock, token); 
				
				
			case TokenType.Import:
				ProcessImport( &formatData, &globalBlock, token); 
				
				
			case TokenType.Struct:
				ProcessStruct( &formatData, &globalBlock, token); 
				
				
			case TokenType.Enum:
				ProcessEnum( &formatData, &globalBlock, token); 
				
				
			case TokenType.Enumstruct:
				ProcessEnumstruct( &formatData, &globalBlock, token); 
				
				
			case TokenType.Interface:
				Process( &formatData, &globalBlock, token); 
				
				
			case TokenType.Fn:
				ProcessFunction( &formatData, &functionList, token); 
				
				
			case TokenType.Var:
				formatData.Result = FormatResult.UnexpectedType; 
				formatData.ErrorDetail = "'var' isn't ever used in TypeGo. Correct syntax: 'int number = 10', not 'var number int = 10'"; 
				formatData.ErrorToken = token; 
				return codeFormat; 
				
			default:
				
				
		}
		
		
		if formatData.IsError() {
			fmt.Println("Error:", formatData.Result.ToString()); 
			PrintFormatError( &formatData, code); 
			* result = formatData.Result; 
			//formatData.PrintFunctionLog()
			return codeFormat; 
		}
		
		formatData.IncrementIfSame(previousIndex); 
			}
	
	
	codeFormat.Functions = functionList
	codeFormat.GlobalBlock = globalBlock
	
	//formatData.PrintFunctionLog()
	
	//fmt.Printf("%sDone%s\n", CYAN_TEXT, RESET_TEXT);
	
	return codeFormat; 
}
