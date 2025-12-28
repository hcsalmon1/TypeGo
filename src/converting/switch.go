
package converting

import . "TypeGo/core"



func ConvertSwitch(convertData *ConvertData, blockData *BlockData, nestCount int, newLine bool) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; 
	var addedSpace bool  = false; 
	
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		
		if token.Type == TokenType.NewLine {
			convertData.NewLineWithTabs(); 
			lastType = token.Type; 
			continue 
			
			
		}
		AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
		HandleTokenSwitch(convertData, token, &nestCount); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		
	}
	
	
	if lastType != TokenType.NewLine {
		convertData.NewLineWithTabs(); 
		
	}
	if newLine == true {
		convertData.NewLineWithTabs(); 
		
	}
	
	
	
}

func HandleTokenSwitch(convertData *ConvertData, token Token, nestCount *int) {
	
	if token.Type == TokenType.Semicolon {
		var codeLength int  = len(convertData.GeneratedCode)
		
		var lastChar byte  = convertData.GeneratedCode[codeLength - 1]
		
		if lastChar == ' ' {
			convertData.GeneratedCode[codeLength - 1] = ';'
			return 
		}
		
	}
	if token.Text == "switch" {
		* nestCount++; 
		
	}
	if token.Text == "}" {
		* nestCount--; 
		convertData.AppendChar('\r'); 
		convertData.AddTabs(); 
		
	}
	convertData.AppendString(token.Text); 
	
}
