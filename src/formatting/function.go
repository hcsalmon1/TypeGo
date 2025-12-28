
package formatting

import (
."TypeGo/core"
)


func TypeAndNameLeftParenthesis(formatData *FormatData, returnType *string, functionName *string, firstToken Token) {
	
	formatData.SetErrorFunction("TypeAndNameLeftParenthesis"); 
	
	var returnTypeTokenList []Token  = make([]Token, 0)
	
	returnTypeTokenList = append(returnTypeTokenList, firstToken)
	
	formatData.TokenIndex += 1
	
	LoopUntilRightParenthesis(formatData, &returnTypeTokenList); 
	
	* returnType = JoinTextInListOfTokens( &returnTypeTokenList); 
	
	var nextToken Token  = formatData.GetToken(); 
	if nextToken.Type == TokenType.NA {
		return 
	}
	if nextToken.Type != TokenType.Identifier {
		formatData.MissingExpectedTypeError(nextToken, "missing identifier in function name"); 
		return; 
	}
	* functionName = nextToken.Text; 
	formatData.Increment(); 
	
}

func TypeAndNameOther(formatData *FormatData, returnType *string, functionName *string, firstToken Token) {
	
	formatData.SetErrorFunction("TypeAndNameOther"); 
	
	var returnTypeTokenList []Token  = make([]Token, 0); 
	FillVarType(formatData, &returnTypeTokenList); 
	
	* returnType = JoinTextInListOfTokens( &returnTypeTokenList); 
	
	var nextToken Token  = formatData.GetToken(); 
	if formatData.IsValidToken(nextToken) == false {
		formatData.EndOfFileError(firstToken); 
		return; 
	}
	if nextToken.Type != TokenType.Identifier {
		formatData.MissingExpectedTypeError(nextToken, "missing identifier in function name"); 
		return; 
	}
	* functionName = nextToken.Text; 
	formatData.Increment(); 
	
}

func TypeAndNameIdentifier(formatData *FormatData, ReturnType *string, functionName *string, firstToken Token) {
	
	formatData.SetErrorFunction("TypeAndNameIdentifier"); 
	
	var index int  = formatData.TokenIndex; 
	var tempToken Token  = formatData.GetNextToken(); 
	
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var nextToken Token  = tempToken; 
	//void type
	if nextToken.Type == TokenType.LeftParenthesis {
		* ReturnType = ""; 
		* functionName = firstToken.Text; 
		return; 
	}
	if nextToken.Type == TokenType.Identifier {
		tempToken = formatData.GetNextToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			return; 
		}
		if tempToken.Type != TokenType.LeftParenthesis {
			formatData.MissingExpectedTypeError(tempToken, "missing expected '(' in function"); 
			return; 
		}
		
		* ReturnType = firstToken.Text; 
		* functionName = nextToken.Text; 
		return; 
	}
	formatData.TokenIndex = index; 
	TypeAndNameOther(formatData, ReturnType, functionName, firstToken); 
	
}

func GetFunctionTypeAndName(formatData *FormatData, returnType *string, functionName *string) {
	formatData.SetErrorFunction("GetFunctionTypeAndName"); 
	
	var tempToken Token  = EmptyToken(); 
	tempToken = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var firstToken Token  = tempToken; 
	
	switch firstToken.Type {
	
	case TokenType.LeftParenthesis:
	TypeAndNameLeftParenthesis(formatData, returnType, functionName, firstToken)
	
	case TokenType.Identifier:
	TypeAndNameIdentifier(formatData, returnType, functionName, firstToken); 
	
	case TokenType.Void:
	formatData.UnexpectedTypeError(firstToken, "Can't use void as a return type")
	return 
	
	default:
	TypeAndNameOther(formatData, returnType, functionName, firstToken); 
	
	
	}
	
	
}

func GetInterfaceMethodTypeAndName(formatData *FormatData, returnType *string, functionName *string) {
	
	formatData.SetErrorFunction("GetFunctionTypeAndName"); 
	
	var tempToken Token  = EmptyToken(); 
	tempToken = formatData.GetToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var firstToken Token  = tempToken; 
	switch firstToken.Type {
	
	case TokenType.Identifier:
	TypeAndNameIdentifier(formatData, returnType, functionName, firstToken); 
	
	case TokenType.LeftParenthesis:
	TypeAndNameLeftParenthesis(formatData, returnType, functionName, firstToken); 
	
	default:
	TypeAndNameOther(formatData, returnType, functionName, firstToken); 
	
	}
	
	
	
}

func ProcessFunction(formatData *FormatData, functions *[]Function, fnToken Token) {
	
	formatData.AddToFunctionLog("ENTER ProcessFunction")
	formatData.SetErrorFunction("ProcessFunction"); 
	
	var parameters []Variable  = make([]Variable, 0)
	
	var returnType string  = ""; 
	var functionName string  = ""; 
	
	GetFunctionTypeAndName(formatData, &returnType, &functionName); 
	if formatData.IsError() {
		return; 
	}
	//Expect '('
	if formatData.ExpectType(TokenType.LeftParenthesis, "Missing expected '(' in fn") == false {
		formatData.AddToFunctionLog("ERROR ProcessFunction")
		return; 
	}
	formatData.Increment(); 
	
	FindParameters(formatData, &parameters); 
	if formatData.IsError() {
		return; 
	}
	
	//Expect ')'
	if formatData.ExpectType(TokenType.RightParenthesis, "Missing expected ')' in fn") == false {
		formatData.AddToFunctionLog("ERROR ProcessFunction")
		return; 
	}
	//Expect '{'
	if formatData.ExpectNextType(TokenType.LeftBrace, "Missing expected '{' in fn") == false {
		formatData.AddToFunctionLog("ERROR ProcessFunction")
		return; 
	}
	formatData.Increment(); 
	
	formatData.AddToProcessLog(functionName + "() {")
	formatData.IncreaseLogIndent()
	var innerBlock CodeBlock  = FillBody(formatData); 
	formatData.DecreaseLogIndent()
	formatData.AddToProcessLog("}")
	
	if formatData.IsError() {
		return 
	}
	
	//Expect '}'
	if formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' in fn") == false {
		formatData.AddToFunctionLog("ERROR ProcessFunction")
		return; 
	}
	
	var function Function  = Function{
	InnerBlock: &innerBlock, 
	Parameters:parameters, 
	Name:functionName, 
	ReturnType:returnType, 
	StartingToken:fnToken, 
	}; 
	formatData.AddToFunctionLog("EXIT ProcessFunction")
	* functions = append(* functions, function); 
	
}

func ProcessInterfaceFunction(formatData *FormatData, functions *[]Function, fnToken Token) {
	
	formatData.SetErrorFunction("ProcessInterfaceFunction"); 
	
	var parameters []Variable  = make([]Variable, 0); 
	var returnType string  = ""; 
	var functionName string  = ""; 
	
	GetInterfaceMethodTypeAndName(formatData, &returnType, &functionName); 
	if formatData.IsError() {
		return; 
	}
	//Expect '('
	if formatData.ExpectType(TokenType.LeftParenthesis, "Missing expected '(' in interface") == false {
		return; 
	}
	formatData.Increment(); 
	
	FindParameters(formatData, &parameters); 
	if formatData.IsError() {
		return; 
	}
	
	//Expect ')'
	if formatData.ExpectType(TokenType.RightParenthesis, "Missing expected ')' in interface") == false {
		return; 
	}
	formatData.Increment(); 
	
	var function Function  = Function{
	InnerBlock:nil, 
	Parameters:parameters, 
	Name:functionName, 
	ReturnType:returnType, 
	}; 
	* functions = append(* functions, function); 
	
}
