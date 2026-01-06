
package formatting

import (
."TypeGo/core"
"strings"
)


func LoopUntilEndInner(formatData *FormatData, blockData *BlockData, token Token, lastTokenType IntTokenType, bracket_counts *BracketCounts, stopAtSemicolon bool) IntLoopAction {
	
	formatData.SetErrorFunction("LoopUntilEndInner"); 
	
	if token.Type == TokenType.LeftSquareBracket {
		bracket_counts.OpenSqCount += 1; 
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.RightSquareBracket {
		bracket_counts.OpenSqCount -= 1; 
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.LeftParenthesis {
		bracket_counts.OpenBracketCount += 1; 
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.RightParenthesis {
		bracket_counts.OpenBracketCount -= 1; 
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.LeftBrace {
		bracket_counts.OpenBraceCount += 1; 
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.RightBrace {
		bracket_counts.OpenBraceCount -= 1; 
		return LoopAction.Continue; 
	}
	
	if bracket_counts.AreAllZero() == false {
		return LoopAction.Continue; 
	}
	
	if stopAtSemicolon == true {
		if token.Type == TokenType.Semicolon {
			blockData.Tokens = append(blockData.Tokens, token)
			
			return LoopAction.Break; 
		}
		
	}
	if token.Type == TokenType.NewLine {
		blockData.Tokens = append(blockData.Tokens, token)
		
		if IsLineContinuingToken(lastTokenType) == false {
			formatData.Increment(); 
			return LoopAction.Break; 
		}
		
	}
	if token.Type == TokenType.EndComment {
		return LoopAction.Break; 
	}
	return LoopAction.Continue; 
}

func LoopTokensUntilLineEnd(formatData *FormatData, blockData *BlockData, stopAtSemicolon bool) {
	
	formatData.SetErrorFunction("LoopTokensUntilLineEnd"); 
	
	var lastTokenType IntTokenType  = TokenType.NA; 
	var bracket_counts BracketCounts
	bracket_counts.OpenBraceCount = 0
	bracket_counts.OpenBracketCount = 0
	bracket_counts.OpenSqCount = 0
	
	for formatData.IndexInBounds() {
	
		var tempToken Token  = formatData.GetToken(); if formatData.IsValidToken(tempToken) == false {
			break
			
		}
		var token Token  = tempToken; var loopResult IntLoopAction  = LoopUntilEndInner(formatData, blockData, token, lastTokenType, &bracket_counts, stopAtSemicolon); if loopResult == LoopAction.Break {
			lastTokenType = token.Type; 
			break
			
		}
		
		formatData.Increment(); 
		blockData.Tokens = append(blockData.Tokens, token)
		
		lastTokenType = token.Type; 
			}
	
	
	if lastTokenType == TokenType.Semicolon {
		formatData.Increment(); 
		var tempToken Token  = formatData.GetToken(); if formatData.IsValidToken(tempToken) {
			if tempToken.Type == TokenType.NewLine {
				formatData.Increment(); 
				
			}
			
		}
		
	}
	
}

func IsLineContinuingToken(token_type IntTokenType) bool {
	switch token_type {
		
		
		case TokenType.Minus, TokenType.Plus, TokenType.Divide, TokenType.Multiply, TokenType.Equals, TokenType.And, 
		TokenType.AndAnd, TokenType.Or, TokenType.OrOr, TokenType.PlusEquals, TokenType.MinusEquals, TokenType.MultiplyEquals, 
		TokenType.DivideEquals, TokenType.GreaterThan, TokenType.LessThan, TokenType.EqualsEquals, TokenType.GreaterThanEquals, 
		TokenType.LessThanEquals, TokenType.Modulus, TokenType.ModulusEquals, TokenType.NotEquals, TokenType.LeftBrace, 
		TokenType.LeftSquareBracket, TokenType.Comma, TokenType.FullStop:
			return true; 
			
		default:
			return false; 
	}
	
	return false
}

func IsInfiniteWhile(count *int, max int) bool {
	* count += 1
	return * count >= max
}

func IsPointerDeclaration(formatData *FormatData) bool {
	
	formatData.SetErrorFunction("IsPointerDeclaration"); 
	
	var index int  = formatData.TokenIndex; var lastType IntTokenType  = TokenType.NA; var identifierFound bool  = false; 
	for formatData.IndexInBounds() {
	
		if index >= len(formatData.TokenList) {
			formatData.EndOfFileError(EmptyToken()); 
			return false; 
		}
		var token Token  = formatData.TokenList[index]; if token.Type == TokenType.LeftSquareBracket {
			if identifierFound == true {
				return false; 
			}
			return true; 
		}
		if token.Type == TokenType.Identifier {
			identifierFound = true; 
			if lastType == TokenType.Identifier {
				return true; 
			}
			
		}
		if IsVarTypeEnum(token.Type) {
			return true; 
		}
		if IsAssignmentOperator(token.Type) {
			return false; 
		}
		lastType = token.Type; 
		index += 1; 
			}
	
	return false; 
}

func HandleToken(formatData *FormatData, token Token, returnTypeTokenList *[]Token, lastType *IntTokenType) bool {
	formatData.SetErrorFunction("HandleToken"); 
	
	if IsVarTypeEnum(token.Type) {
		* returnTypeTokenList = append(* returnTypeTokenList, token); 
		return false; 
	}
	
	if token.Type == TokenType.Identifier {
		if * lastType == TokenType.Identifier || * lastType == TokenType.RightBrace {
			return true; 
		}
		if IsVarTypeEnum(* lastType) {
			return true; 
		}
		
		var wasPartOfType bool  = 
		* lastType == TokenType.FullStop || 
		* lastType == TokenType.RightSquareBracket || 
		* lastType == TokenType.Multiply || 
		* lastType == TokenType.LeftParenthesis || 
		* lastType == TokenType.NA; 
		if wasPartOfType {
			* returnTypeTokenList = append(* returnTypeTokenList, token); 
			return false; 
		}
		
		formatData.UnexpectedTypeError(token, "Found unexpected identifier in parameters"); 
		return true; 
	}
	var isSkippableToken bool  = 
	token.Type == TokenType.Tab || 
	token.Type == TokenType.NewLine || 
	token.Type == TokenType.RightParenthesis || 
	token.Type == TokenType.Comma; 
	if isSkippableToken {
		return false; 
	}
	* returnTypeTokenList = append(* returnTypeTokenList, token); 
	return false; 
}

func VarTypeInnerCode(formatData *FormatData, returnTypeTokenList *[]Token, whileCount *int, lastType *IntTokenType) IntLoopAction {
	formatData.SetErrorFunction("VarTypeInnerCode"); 
	
	var MAX int  = 10000; 
	if IsInfiniteWhile(whileCount, MAX) {
		formatData.Result = FormatResult.Internal_Error; 
		formatData.ErrorDetail = "Infinite while loop in VarTypeInnerLoop, FormatUtils"; 
		return LoopAction.Error; 
	}
	var indexBefore int  = formatData.TokenIndex; 
	var token Token  = formatData.GetToken(); if formatData.IsValidToken(token) == false {
		formatData.EndOfFileError(token); 
		return LoopAction.Return; 
	}
	
	var shouldBreak bool  = HandleToken(formatData, token, returnTypeTokenList, lastType); if shouldBreak == true {
		return LoopAction.Break; 
	}
	
	formatData.IncrementIfSame(indexBefore); 
	* lastType = token.Type; 
	return LoopAction.Continue; 
}

func FillVarType(formatData *FormatData, returnTypeTokenList *[]Token) {
	
	formatData.SetErrorFunction("FillVarType"); 
	
	var whileCount int  = 0; 
	var lastType IntTokenType  = TokenType.NA; 
	//Get to ')' and then don't increment
	for formatData.IndexInBounds() {
	
		var iterationResult IntLoopAction  = VarTypeInnerCode(formatData, returnTypeTokenList, &whileCount, &lastType); if iterationResult == LoopAction.Break {
			break
			
		}
		if iterationResult == LoopAction.Return {
			return; 
		}
			}
	
	
}

func LoopUntilRightParenthesis(formatData *FormatData, returnTypeTokenList *[]Token) {
	formatData.SetErrorFunction("LoopUntilRightParenthesis"); 
	
	var whileCount int  = 0; 
	var openParenthesisCount int  = 1; 
	for formatData.IndexInBounds() {
	
		var MAX int  = 10000; 
		if IsInfiniteWhile( &whileCount, MAX) {
			formatData.Result = FormatResult.Internal_Error; 
			formatData.ErrorDetail = "Infinite while loop in LoopUntilRightParenthesis, FormatUtils"; 
			return; 
		}
		var indexBefore int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if formatData.IsValidToken(token) == false {
			formatData.EndOfFileError(token); 
			return; 
		}
		
		* returnTypeTokenList = append(* returnTypeTokenList, token); 
		
		formatData.IncrementIfSame(indexBefore); 
		
		if token.Type == TokenType.LeftParenthesis {
			openParenthesisCount += 1; 
			
		}
		
		if token.Type == TokenType.RightParenthesis {
			if openParenthesisCount != 1 {
				openParenthesisCount -= 1; 
				continue
				
			}
			break
			
		}
			}
	
	
}

func FindParameters(formatData *FormatData, parameters *[]Variable) {
	formatData.SetErrorFunction("FindParameters"); 
	
	var whileCount int  = 0; 
	var parameter_data ParameterData
	
	parameter_data.TempParameter = Variable {}
	parameter_data.TempParameter.SetToDefaults()
	parameter_data.LastTokenType = TokenType.LeftParenthesis
	parameter_data.ParameterPhase = ParameterPhase.TypeOrName
	parameter_data.Parameters = parameters
	
	//Get to ')' and then don't increment
	for formatData.IndexInBounds() {
	
		var iterationResult IntLoopAction  = ParameterInnerLoop(formatData, &parameter_data, &whileCount); if iterationResult == LoopAction.Break {
			break
			
		}
		if iterationResult == LoopAction.Return {
			return; 
		}
			}
	
	
}

func ParameterInnerLoop(formatData *FormatData, parameter_data *ParameterData, whileCount *int) IntLoopAction {
	formatData.SetErrorFunction("ParameterInnerLoop"); 
	var MAX int  = 10000; 
	if IsInfiniteWhile(whileCount, MAX) {
		formatData.Result = FormatResult.Internal_Error; 
		formatData.ErrorDetail = "infinite while loop in ParameterInnerLoop, FunctionUtils"; 
		return LoopAction.Error; 
	}
	var indexBefore int  = formatData.TokenIndex; 
	var token Token  = formatData.GetToken(); if formatData.IsValidToken(token) == false {
		formatData.EndOfFileError(token); 
		return LoopAction.Return; 
	}
	
	HandleParameterToken(formatData, token, parameter_data); 
	
	if formatData.IsError() {
		return LoopAction.Return
	}
	
	if token.Type == TokenType.Comma {
		AddParameter(parameter_data); 
		parameter_data.ParameterPhase = ParameterPhase.TypeOrName
		
	}
	if token.Type == TokenType.RightParenthesis {
		AddParameter(parameter_data); 
		return LoopAction.Break; 
	}
	
	formatData.IncrementIfSame(indexBefore); 
	parameter_data.LastTokenType = token.Type; 
	return LoopAction.Continue; 
}

func HandleParameterToken(formatData *FormatData, token Token, parameter_data *ParameterData) {
	
	if token.Type == TokenType.Identifier {
		if parameter_data.ParameterPhase == ParameterPhase.End {
			formatData.UnexpectedTypeError(token, "expected comma or ) in parameters")
			return 
		}
		
		if parameter_data.LastTokenType == TokenType.Identifier || parameter_data.LastTokenType == TokenType.RightBrace {
			parameter_data.TempParameter.NameToken = append(parameter_data.TempParameter.NameToken, token)
			
			parameter_data.ParameterPhase = ParameterPhase.End
			return; 
		}
		if IsVarTypeEnum(parameter_data.LastTokenType) {
			parameter_data.TempParameter.NameToken = append(parameter_data.TempParameter.NameToken, token)
			
			return; 
		}
		
		var wasPartOfType bool  = 
		parameter_data.LastTokenType == TokenType.FullStop || 
		parameter_data.LastTokenType == TokenType.RightSquareBracket || 
		parameter_data.LastTokenType == TokenType.Multiply || 
		parameter_data.LastTokenType == TokenType.Comma || 
		parameter_data.LastTokenType == TokenType.LeftParenthesis; 
		if wasPartOfType {
			parameter_data.TempParameter.TypeList = append(parameter_data.TempParameter.TypeList, token)
			
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
	parameter_data.TempParameter.TypeList = append(parameter_data.TempParameter.TypeList, token)
	
	
}

func AddParameter(parameter_data *ParameterData) {
	
	if parameter_data.TempParameter.NameToken == nil {
		return; 
	}
	var parameter_count int  = len(parameter_data.TempParameter.TypeList)
	if parameter_count == 0 {
		return; 
	}
	var parameterCopy Variable
	parameterCopy.TypeList = make([]Token, 0)
	for i := 0; i < parameter_count; i++ {
	
		parameterCopy.TypeList = append(parameterCopy.TypeList, parameter_data.TempParameter.TypeList[i])
		
			}
	
	parameterCopy.NameToken = CopyTokenList(parameter_data.TempParameter.NameToken); 
	
	* parameter_data.Parameters = append(* parameter_data.Parameters, parameterCopy); 
	parameter_data.TempParameter.SetToDefaults(); 
	
}

func ConcatStrings(slice []string) string {
	var result strings.Builder
	for _, s := range slice {
	
		result.WriteString(s)
			}
	
	return result.String()
}
