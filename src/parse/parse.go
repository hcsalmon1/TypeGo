
package parse

import (
."TypeGo/core"
 //"fmt"
)


func ParseToTokens(err *IntParseResult, code string) []Token {
	
	//fmt.Print("\tParsing\t\t\t")
	
	var parse_data ParseData
	parse_data.Code = code; 
	
	parse_data.CodeLength = len(code)
	if parse_data.CodeLength == 0 {
		* err = ParseResult.String_Length_Zero; 
		return parse_data.TokenList; 
	}
	
	var while_count int  = 0
	var MAX_WHILE_ITERATIONS int  = parse_data.CodeLength + 100
	
	parse_data.CharacterIndex = 0
	for parse_data.CharacterIndex < parse_data.CodeLength {
	
		if while_count >= MAX_WHILE_ITERATIONS {
			* err = ParseResult.Infinite_While_Loop
			return parse_data.TokenList
		}
		while_count += 1
		
		processCharacter( &parse_data); 
		if parse_data.ParseResult != ParseResult.Ok {
			* err = parse_data.ParseResult; 
			return parse_data.TokenList; 
		}
		
			}
	
	//fmt.Println("Done")
	return parse_data.TokenList; 
}

func processCharacter(parse_data *ParseData) {
	
	if shouldSkip(parse_data) == true {
		return 
	}
	
	var previousCharacterIndex int  = parse_data.CharacterIndex; 
	var success bool  = true
	var token Token  = getToken( &success, parse_data); 
	if previousCharacterIndex == parse_data.CharacterIndex {
		parse_data.CharacterIndex += 1; 
		
	}
	if isInvalid(token, success, parse_data) == true {
		return; 
	}
	
	parse_data.TokenList = append(parse_data.TokenList, token)
	
	parse_data.LastToken = token; 
	
}

func shouldSkip(parse_data *ParseData) bool {
	
	var current_char byte  = parse_data.Code[parse_data.CharacterIndex]; 
	if current_char == '\n' {
		var token Token  = Token{
			Text:"\\n", 
			Type:TokenType.NewLine, 
			LineNumber:parse_data.LineCount, 
			CharNumber:parse_data.CharCount, 
		}
		
		parse_data.TokenList = append(parse_data.TokenList, token)
		
		
		parse_data.LineCount += 1; 
		parse_data.CharCount = 0; 
		parse_data.CharacterIndex += 1; 
		return true; 
	}
	var is_special_char bool  = 
	current_char == '\r' || 
	current_char == '\t' || 
	current_char == ' ' || 
	current_char == '\\'; 
	if current_char == '\t' {
		parse_data.CharCount += 8; 
		
	} else if current_char == ' ' || current_char == '\\' {
		parse_data.CharCount += 1; 
		
	}
	
	
	if is_special_char == true {
		parse_data.CharacterIndex += 1; 
		return true; 
	}
	return false; 
}

func isInvalid(token Token, success bool, parseData *ParseData) bool {
	
	if success == false {
		return true; 
	}
	if parseData.IsError() {
		return true; 
	}
	return false; 
}

func getToken(success *bool, parse_data *ParseData) Token {
	
	if parse_data.CharacterIndex + 1 < parse_data.CodeLength {
		var c1 byte  = parse_data.Code[parse_data.CharacterIndex]; var c2 byte  = parse_data.Code[parse_data.CharacterIndex + 1]; 
		if c1 == '/' && c2 == '/' {
			return readLineComment(parse_data); 
		}
		if c1 == '/' && c2 == '*' {
			return readBlockComment(parse_data); 
		}
		
	}
	
	var current_char byte  = parse_data.Code[parse_data.CharacterIndex]; 
	if current_char == '"' {
		return readString(success, parse_data); 
	}
	if current_char == '\'' {
		return readChar(success, parse_data); 
	}
	if isOperator(current_char) {
		return readOperator(success, parse_data); 
	}
	if isSeparator(current_char) {
		return readSeparator(success, parse_data); 
	}
	if current_char == '`' {
		return ReadMultilineString(parse_data); 
	}
	
	return readWord(success, parse_data); 
}

func readString(success *bool, parseData *ParseData) Token {
	
	var string_builder []byte  = make([]byte, 0)
	
	string_builder = append(string_builder, parseData.Code[parseData.CharacterIndex])
	
	parseData.CharacterIndex += 1; 
	
	var lastChar byte  = ' '; 
	for parseData.CharacterIndex < parseData.CodeLength {
	
		var currentChar byte  = parseData.Code[parseData.CharacterIndex]; 
		var isStringEnd bool  = 
		currentChar == '"' && 
		lastChar != '\\'; 
		if isStringEnd == true {
			string_builder = append(string_builder, currentChar)
			
			parseData.CharacterIndex++; 
			var token Token  = Token{
				Text:string(string_builder), 
				Type:TokenType.StringValue, 
				LineNumber:parseData.LineCount, 
				CharNumber:parseData.CharCount, 
			}
			parseData.CharCount += len(token.Text); 
			return token
		}
		string_builder = append(string_builder, currentChar)
		
		if currentChar == '\\' && lastChar == '\\' {
			lastChar = ' '; 
			
		} else  {
			lastChar = currentChar; 
			
		}
		
		parseData.CharacterIndex++; 
			}
	
	
	* success = false
	parseData.ParseResult = ParseResult.Unterminated_String; 
	return EmptyToken()
}

func readSeparator(success *bool, parseData *ParseData) Token {
	var c byte  = parseData.Code[parseData.CharacterIndex]; parseData.CharacterIndex += 1; 
	var tokenText string  = string(c)
	var token Token  = Token{
		Text:tokenText, 
		Type:getTokenType(tokenText), 
		LineNumber:parseData.LineCount, 
		CharNumber:parseData.CharCount, 
	}
	parseData.CharCount += 1; 
	return token
}

func readChar(success *bool, parseData *ParseData) Token {
	
	parseData.CharacterIndex += 1; 
	if parseData.CharacterIndex >= parseData.CodeLength {
		parseData.ParseResult = ParseResult.Unexpected_Value; 
		* success = false
		return EmptyToken()
	}
	
	var charValue byte  = parseData.Code[parseData.CharacterIndex]; var charValueSecondPart byte  = ' '; parseData.CharacterIndex += 1; 
	
	if charValue == '\\' {
		charValueSecondPart = parseData.Code[parseData.CharacterIndex]; 
		parseData.CharacterIndex += 1; 
		
	}
	
	if parseData.CharacterIndex >= parseData.CodeLength {
		parseData.ParseResult = ParseResult.Unexpected_Value; 
		* success = false
		return EmptyToken()
	}
	if parseData.Code[parseData.CharacterIndex] != '\'' {
		parseData.ParseResult = ParseResult.Unterminated_Char; 
		* success = false
		return EmptyToken()
	}
	parseData.CharacterIndex += 1; 
	
	if charValueSecondPart != ' ' {
		var token_text string  = "'" +string([]byte {charValue, charValueSecondPart})+"'"
		parseData.CharCount += len(token_text); 
		return Token {Text:token_text, Type:TokenType.CharValue, LineNumber:parseData.LineCount, CharNumber:parseData.CharCount}
	}
	parseData.CharCount += 3; 
	return Token {Text:"'" +string(charValue)+"'", Type:TokenType.CharValue, LineNumber:parseData.LineCount, CharNumber:parseData.CharCount}
}

func readOperator(success *bool, parseData *ParseData) Token {
	
	var string_builder []byte  = make([]byte, 0)
	var c byte  = parseData.Code[parseData.CharacterIndex]; string_builder = append(string_builder, c)
	
	
	parseData.CharacterIndex += 1; 
	
	// Lookahead for compound operators like "==", "!="
	if parseData.CharacterIndex < parseData.CodeLength {
		var next byte  = parseData.Code[parseData.CharacterIndex]; if isOperator(next) {
			string_builder = append(string_builder, next)
			
			parseData.CharacterIndex += 1; 
			
		}
		
	}
	
	var tokenText string  = string(string_builder)
	var token Token  = Token{
		Text:tokenText, 
		Type:getTokenType(tokenText), 
		LineNumber:parseData.LineCount, 
		CharNumber:parseData.CharCount, 
	}
	parseData.CharCount += len(tokenText); 
	return token
}

func readWord(success *bool, parseData *ParseData) Token {
	
	var string_builder []byte  = make([]byte, 0)
	
	for parseData.CharacterIndex < parseData.CodeLength {
	
		var c byte  = parseData.Code[parseData.CharacterIndex]; 
		if isLetterOrDigit(c) || c == '_' {
			string_builder = append(string_builder, c)
			
			parseData.CharacterIndex++; 
			
		} else  {
			break
			
		}
		
			}
	
	
	var word string  = string(string_builder)
	var token Token  = Token{
		Text:word, 
		Type:getTokenType(word), 
		LineNumber:parseData.LineCount, 
		CharNumber:parseData.CharCount, 
	}
	parseData.CharCount += len(word); 
	return token
}

func readLineComment(parseData *ParseData) Token {
	
	var string_builder []byte  = make([]byte, 0)
	
	string_builder = append(string_builder, '/')
	
	string_builder = append(string_builder, '/')
	
	parseData.CharacterIndex += 2; 
	
	for parseData.CharacterIndex < parseData.CodeLength {
	
		var c byte  = parseData.Code[parseData.CharacterIndex]; if c == '\n' {
			break
			
		}
		string_builder = append(string_builder, c)
		
		parseData.CharacterIndex++; 
			}
	
	var token Token  = Token{
		Text:string(string_builder), 
		Type:TokenType.Comment, 
		LineNumber:parseData.LineCount, 
		CharNumber:parseData.CharCount, 
	}
	parseData.CharCount += len(token.Text); 
	return token
}

func readBlockComment(parseData *ParseData) Token {
	
	var string_builder []byte  = make([]byte, 0)
	string_builder = append(string_builder, '/')
	
	string_builder = append(string_builder, '*')
	
	parseData.CharacterIndex += 2; 
	
	for parseData.CharacterIndex + 1 < parseData.CodeLength {
	
		var c1 byte  = parseData.Code[parseData.CharacterIndex]; var c2 byte  = parseData.Code[parseData.CharacterIndex + 1]; string_builder = append(string_builder, c1)
		
		parseData.CharacterIndex++; 
		
		if (c1 == '*' && c2 == '/') {
			string_builder = append(string_builder, '/')
			
			parseData.CharacterIndex++; 
			return Token {Text:string(string_builder), Type:TokenType.Comment, LineNumber:parseData.LineCount, CharNumber:parseData.CharCount}
		}
		
		if c1 == '\n' {
			parseData.LineCount++; 
			parseData.CharCount = 0; 
			
		}
			}
	
	
	parseData.ParseResult = ParseResult.Unterminated_Comment; 
	return EmptyToken()
}

func ReadMultilineString(parseData *ParseData) Token {
	
	var sb []byte  = make([]byte, 0); sb = append(sb, parseData.Code[parseData.CharacterIndex])
	
	parseData.CharacterIndex += 1; 
	
	var lastChar byte  = ' '; 
	for parseData.CharacterIndex < len(parseData.Code) {
	
		var currentChar byte  = parseData.Code[parseData.CharacterIndex]; 
		var isStringEnd bool  = 
		currentChar == '`' && lastChar != '\\'; 
		if isStringEnd == true {
			sb = append(sb, currentChar)
			
			parseData.CharacterIndex++; 
			var token Token  = Token{
				Text:string(sb), 
				Type:TokenType.StringValue, 
				LineNumber:parseData.LineCount, 
				CharNumber:parseData.CharCount, 
			}
			return token; 
		}
		sb = append(sb, currentChar)
		
		if currentChar == '\\' && lastChar == '\\' {
			lastChar = ' '; 
			
		} else  {
			lastChar = currentChar; 
			
		}
		
		parseData.CharacterIndex++; 
			}
	
	
	parseData.ParseResult = ParseResult.Unterminated_String; 
	return EmptyToken(); 
}
