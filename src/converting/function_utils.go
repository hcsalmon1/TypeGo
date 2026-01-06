
package converting

import . "TypeGo/core"



func AddParameter(parameters []Variable, tempParameter *Variable) {
	if tempParameter.NameToken == nil {
		return; 
	}
	if len(tempParameter.TypeList) == 0 {
		return; 
	}
	var parameterCopy Variable  = Variable{}; parameterCopy.TypeList = make([]Token, 0)
	for i := 0; i < len(tempParameter.TypeList); i++ {
	
		parameterCopy.TypeList = append(parameterCopy.TypeList, tempParameter.TypeList[i])
		
		
	}
	
	parameterCopy.NameToken = CopyTokenList(tempParameter.NameToken); 
	
	parameters = append(parameters, parameterCopy)
	
	tempParameter.SetToDefaults(); 
	
}

func HandleTokenParameter(formatData *FormatData, token Token, tempParameter *Variable, lastType *IntTokenType) {
	
	if token.Type == TokenType.Identifier {
		if * lastType == TokenType.Identifier || * lastType == TokenType.RightBrace {
			tempParameter.NameToken = append(tempParameter.NameToken, token)
			
			return; 
		}
		if IsVarTypeEnum(* lastType) {
			tempParameter.NameToken = append(tempParameter.NameToken, token)
			
			return; 
		}
		
		var wasPartOfType bool  = 
		* lastType == TokenType.FullStop || 
		* lastType == TokenType.RightSquareBracket || 
		* lastType == TokenType.Multiply || 
		* lastType == TokenType.Comma || 
		* lastType == TokenType.LeftParenthesis; 
		if wasPartOfType {
			tempParameter.TypeList = append(tempParameter.TypeList, token)
			
			return; 
		}
		
		formatData.UnexpectedTypeError(token, "Found unexpected identifier in parameters"); 
		return; 
	}
	var isSkippableToken bool  = 
	token.Type == TokenType.Tab || 
	token.Type == TokenType.NewLine || 
	token.Type == TokenType.RightParenthesis || 
	token.Type == TokenType.Comma; 
	if isSkippableToken {
		return; 
	}
	tempParameter.TypeList = append(tempParameter.TypeList, token)
	
	
}

func ParameterInnerLoop(formatData *FormatData, parameters *[]Variable, tempParameter *Variable, whileCount *int, lastType *IntTokenType) IntLoopAction {
	
	formatData.SetErrorFunction("ParameterInnerLoop"); 
	var MAX int  = 10000; 
	if * whileCount >= MAX {
		formatData.Result = FormatResult.Internal_Error; 
		formatData.ErrorDetail = "infinite while loop in ParameterInnerLoop, FunctionUtils"; 
		return LoopAction.Error; 
	}
	* whileCount += 1
	var indexBefore int  = formatData.TokenIndex; 
	var token Token  = formatData.GetToken(); if token.Text == "" {
		formatData.EndOfFileError(token); 
		return LoopAction.Return; 
	}
	
	HandleTokenParameter(formatData, token, tempParameter, lastType); 
	
	if token.Type == TokenType.Comma {
		AddParameter(* parameters, tempParameter); 
		
	}
	if token.Type == TokenType.RightParenthesis {
		AddParameter(* parameters, tempParameter); 
		return LoopAction.Break; 
	}
	
	formatData.IncrementIfSame(indexBefore); 
	* lastType = token.Type; 
	return LoopAction.Continue; 
}

func FindParameters(formatData *FormatData, parameters []Variable) {
	
	formatData.SetErrorFunction("FindParameters")
	
	var whileCount int  = 0; 
	var tempParameter Variable  = Variable{}; tempParameter.SetToDefaults(); 
	var lastType IntTokenType  = TokenType.LeftParenthesis; 
	//Get to ')' and then don't increment
	for formatData.TokenIndex < len(formatData.TokenList) {
	
		var iterationResult IntLoopAction  = ParameterInnerLoop(formatData, &parameters, &tempParameter, &whileCount, &lastType); if iterationResult == LoopAction.Break {
			break
			
		}
		if iterationResult == LoopAction.Return {
			return; 
		}
		
	}
	
	
}
