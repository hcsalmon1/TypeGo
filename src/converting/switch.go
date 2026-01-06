
package converting

import . "TypeGo/core"



func ConvertSwitch(convertData *ConvertData, blockData *BlockData, nestCount int, newLine bool) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if token.Type == TokenType.NewLine {
			convertData.NewLineWithTabs()
			continue
		}
		convertData.AppendString(token.Text)
		if token.Type == TokenType.Switch {
			convertData.AppendChar(' ')
			
		}
		
	}
	
	convertData.AppendChar(' ')
	convertData.AppendChar('{')
	convertData.IncrementNestCount()
	convertData.NewLineWithTabs()
	if blockData.Block != nil {
		ProcessBlock(convertData, blockData.Block, nestCount + 1, false); 
		
	}
	convertData.DecrementNestCount()
	convertData.RemoveLastTab()
	convertData.AppendChar('}')
	
	if newLine == true {
		convertData.NewLineWithTabs(); 
		
	}
	
	
}

func ConvertSwitchCase(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if token.Type == TokenType.NewLine {
			convertData.NewLineWithTabs()
			continue
		}
		
		convertData.AppendString(token.Text)
		if token.Type == TokenType.Case || token.Type == TokenType.Comma {
			convertData.AppendChar(' ')
			
		}
		
	}
	
	convertData.IncrementNestCount()
	if blockData.Block != nil {
		ProcessBlock(convertData, blockData.Block, nestCount + 1, false); 
		
	}
	convertData.DecrementNestCount()
	
	convertData.NewLineWithTabs(); 
	
}
