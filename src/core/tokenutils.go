
package core


func JoinTextInListOfTokens(tokenList *[]Token) string {
	if tokenList == nil {
		return ""; 
	}
	var count int  = len(* tokenList)
	
	if count == 0 {
		return ""; 
	}
	var outputSB []string  = make([]string, 0)
	
	for i := 0; i < count; i++ {
	
		outputSB = append(outputSB, (*tokenList)[i].Text)
		
		
	}
	
	var result string  = ""
	
	for i := 0; i < len(outputSB); i++ {
	
		result += outputSB[i]
		
	}
	
	return result
}

func JoinTextInListOfTokensWithSpaces(tokenList []Token) string {
	if tokenList == nil {
		return ""; 
	}
	var count int  = len(tokenList)
	
	if count == 0 {
		return ""; 
	}
	var outputSB []byte  = make([]byte, 0); 
	for i := 0; i < count; i++ {
	
		var text string  = tokenList[i].Text
		
		var length int  = len(text)
		
		for j := 0; i < length; j++ {
		
			outputSB = append(outputSB, text[j])
			
			
		}
		
		outputSB = append(outputSB, ' ')
		
		
	}
	
	return string(outputSB)
}

func AppendStringToSlice(slice *[]byte, input string) {
	var length int  = len(input)
	
	if length == 0 {
		return 
	}
	for i := 0; i < length; i++ {
	
		var c byte  = input[i]
		
		* slice = append(* slice, c)
		
	}
	
	
}

func CopyTokenList(tokenList []Token) []Token {
	
	var copyTokenList []Token  = make([]Token, 0)
	
	var count int  = len(tokenList)
	
	if (count == 0) {
		return copyTokenList; 
	}
	for i := 0; i < count; i++ {
	
		copyTokenList = append(copyTokenList, tokenList[i])
		
		
	}
	
	return copyTokenList; 
}

func IsAssignmentOperator(token_type IntTokenType) bool {
	return token_type == TokenType.Equals || token_type == TokenType.PlusEquals || token_type == TokenType.DivideEquals || token_type == TokenType.MultiplyEquals || token_type == TokenType.MinusEquals || token_type == TokenType.ModulusEquals || token_type == TokenType.PlusPlus || token_type == TokenType.MinusMinus; 
}

func IsVarTypePart(token Token) bool {
	switch token.Type {
	case TokenType.LeftSquareBracket, TokenType.Multiply:
	return true; 
	default:
	return false; 
	
	}
	
	
}

func IsVarType(token Token) bool {
	
	switch token.Type {
	case TokenType.Int, TokenType.Int8, TokenType.Int16, TokenType.Int32, 
	TokenType.Int64, TokenType.Uint, TokenType.Uint8, TokenType.Uint16, 
	TokenType.Uint32, TokenType.Uint64, TokenType.Float32, TokenType.Float64, 
	TokenType.String, TokenType.Byte, TokenType.Rune, TokenType.Bool, 
	TokenType.Error, TokenType.Struct, TokenType.Map:
	
	return true; 
	
	default:
	return false; 
	
	}
	
	
}

func IsVarTypeEnum(tokenType IntTokenType) bool {
	
	switch tokenType {
	
	case TokenType.Int, TokenType.Int8, TokenType.Int16, TokenType.Int32, TokenType.Int64, 
	TokenType.Uint, TokenType.Uint8, TokenType.Uint16, TokenType.Uint32, TokenType.Uint64, 
	TokenType.Float32, 
	TokenType.Float64, 
	TokenType.String, 
	TokenType.Byte, 
	TokenType.Rune, 
	TokenType.Bool, 
	TokenType.Error, 
	TokenType.Map, 
	TokenType.Struct, 
	TokenType.Interface:
	return true; 
	default:
	return false; 
	
	}
	
	
}

func IsOperator(tokenType IntTokenType) bool {
	return tokenType == TokenType.Plus || tokenType == TokenType.Minus; 
}
