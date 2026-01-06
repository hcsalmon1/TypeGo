
package core

import (
"fmt"
"strings"
)


func PrintTokenList(token_list []Token) {
	fmt.Println(CYAN_TEXT, "\nPrinting Token List:", RESET_TEXT); 
	var list_count int  = len(token_list); if list_count == 0 {
		fmt.Println("Token list is zero"); 
		return; 
	}
	for i := 0; i < list_count; i++ {
	
		fmt.Printf("\t%sToken:%s %s'%s'%s, %stype:%s %s'%s'%s, %sline number:%s %s%d%s, %schar count:%s %s%d%s\n", GREY_TEXT, RESET_TEXT, ORANGE_TEXT, token_list[i].Text, RESET_TEXT, GREY_TEXT, RESET_TEXT, CREAM_TEXT, token_list[i].Type.ToString(), RESET_TEXT, GREY_TEXT, RESET_TEXT, LIGHT_GREEN_TEXT, token_list[i].LineNumber, RESET_TEXT, GREY_TEXT, RESET_TEXT, LIGHT_GREEN_TEXT, token_list[i].CharNumber, RESET_TEXT, ); 
			}
	
	fmt.Println(); 
	
}

func PrintFormatError(formatData *FormatData, code string) {
	
	var error_token Token  = formatData.ErrorToken; if error_token.Type == TokenType.NA {
		if formatData.Result == FormatResult.EndOfFile {
			error_token = FindLastToken(formatData); 
			
		}
		
		
	}
	var line_number int  = GetLineNumber(error_token); var char_number int  = GetCharNumber(error_token); 
	var codeLines []string  = strings.Split(code, "\n")
	
	var code_line string  = GetCodeLine(line_number, codeLines); 
	fmt.Printf("\t%sError on line %d, %d: %s%s\n", CREAM_TEXT, line_number, char_number, formatData.ErrorDetail, RESET_TEXT); 
	
	PrintCodeLines(line_number, codeLines, code_line); 
	pointToErrorPosition(char_number)
	PrintErrorToken(error_token); 
	
	fmt.Printf("\tError Function: %s%s\n", CREAM_TEXT, formatData.ErrorFunction); 
	fmt.Printf("\tError Process: %s%s\n", formatData.ErrorProcess, RESET_TEXT); 
	//formatData.PrintProcessLog()
	
}

func pointToErrorPosition(char_number int) {
	var temp_curly_lines []byte  = make([]byte, 0)
	for i := 0; i < char_number; i++ {
	
		temp_curly_lines = append(temp_curly_lines, '~')
		
			}
	
	var curly_lines string  = string(temp_curly_lines)
	fmt.Printf("\t%s%s^%s\n", GREEN_TEXT, curly_lines, RESET_TEXT)
	
}

func PrintConvertError(convertData *ConvertData, code string) {
	
	var error_token Token  = convertData.ErrorToken; var line_number int  = GetLineNumber(error_token); var char_number int  = GetCharNumber(error_token); 
	var codeLines []string  = strings.Split(code, "\n")
	
	var code_line string  = GetCodeLine(line_number, codeLines); 
	fmt.Printf("\t%sError on line %d, %d: %s%s\n", CREAM_TEXT, line_number, char_number, convertData.ErrorDetail, RESET_TEXT); 
	
	PrintCodeLines(line_number, codeLines, code_line); 
	pointToErrorPosition(char_number)
	PrintErrorToken(error_token); 
	
	fmt.Printf("\tError Function: %s%s\n", CREAM_TEXT, convertData.ErrorFunction); 
	fmt.Printf("\tError Process: %s%s\n", convertData.ErrorProcess, RESET_TEXT); 
	//convertData.PrintProcessLog()
	
}

func GetCodeLine(line_number int, codeLines []string) string {
	if line_number >= len(codeLines) || line_number < 0 {
		return "Invalid"; 
	} else  {
		return codeLines[line_number]; 
	}
	
	
}

func PrintCodeLines(line_number int, codeLines []string, code_line string) {
	if line_number - 1 >= 0 && line_number < len(codeLines) {
		fmt.Printf("\t%s\n", codeLines[line_number - 1]); 
		
	}
	fmt.Printf("\t%s\n", code_line); 
	
}

func PrintErrorToken(error_token Token) {
	fmt.Print("\tToken: "); 
	if error_token.Type != TokenType.NA {
		fmt.Printf("%s'%s'%s\n", ORANGE_TEXT, error_token.Text, RESET_TEXT); 
		
	} else  {
		fmt.Println("Error token not set"); 
		
	}
	
	
}

func GetLineNumber(error_token Token) int {
	if error_token.Type == TokenType.NA {
		return -1; 
	}
	return error_token.LineNumber; 
}

func GetCharNumber(error_token Token) int {
	if error_token.Type == TokenType.NA {
		return -1; 
	}
	return error_token.CharNumber; 
}

func FindLastToken(formatData *FormatData) Token {
	
	var count int  = len(formatData.TokenList); if count == 0 {
		return EmptyToken(); 
	}
	
	return formatData.GetTokenByIndex(count - 1); 
}
