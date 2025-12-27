
package formatting

import . "TypeGo/core"



func FillAppend(formatData *FormatData, blockData *BlockData) {
	
	formatData.AddToFunctionLog("ENTER FillAppend")
	formatData.SetErrorFunction("FillAppend"); 
	
	blockData.NodeType = NodeType.Append; 
	
	var tempToken Token  = EmptyToken(); 
	
	var startingIndex int  = formatData.TokenIndex; 
	var appendIndex int  = -1; 
	
	for formatData.IndexInBounds() {
	
		tempToken = formatData.GetToken(); 
		
		if tempToken.Type == TokenType.NA {
			formatData.EndOfFileError(tempToken); 
			formatData.AddToFunctionLog("ERROR FillAppend")
			return; 
		}
		
		if tempToken.Type == TokenType.Append {
			appendIndex = formatData.TokenIndex; 
			break
			
		}
		formatData.Increment(); 
			}
	
	
	if appendIndex == -1 {
		formatData.UnexpectedTypeError(tempToken, "Missing 'append'"); 
		formatData.AddToFunctionLog("ERROR FillAppend")
		return; 
	}
	
	var nameVariable Variable
	nameVariable.TypeList = make([]Token, 0); 
	
	var nameTextBuilder []string  = make([]string, 0)
	
	
	for i := startingIndex; i < appendIndex; i++ {
	
		if i + 1 == appendIndex {
			break
			
		}
		
		tempToken = formatData.GetTokenByIndex(i); 
		if tempToken.Type == TokenType.NA {
			formatData.EndOfFileError(tempToken); 
			formatData.AddToFunctionLog("ERROR FillAppend")
			return; 
		}
		nameTextBuilder = append(nameTextBuilder, tempToken.Text)
		
			}
	
	
	var name_token Token  = Token{
	Text:ConcatStrings(nameTextBuilder), 
	Type:TokenType.Identifier, 
	LineNumber:0, 
	CharNumber:0, 
	}
	
	nameVariable.NameToken = append(nameVariable.NameToken, name_token)
	
	
	blockData.Variables = append(blockData.Variables, nameVariable)
	
	
	formatData.TokenIndex = appendIndex + 1; 
	formatData.ExpectType(TokenType.LeftParenthesis, "expecting a '(' after append"); 
	formatData.Increment(); 
	
	var openParenthesisCount int  = 1; 
	
	for formatData.IndexInBounds() {
	
		tempToken = formatData.GetToken(); 
		
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			formatData.AddToFunctionLog("ERROR FillAppend")
			return; 
		}
		
		if tempToken.Type == TokenType.LeftParenthesis {
			openParenthesisCount += 1; 
			
			
		} else if tempToken.Type == TokenType.RightParenthesis {
		
		if openParenthesisCount != 1 {
			openParenthesisCount -= 1; 
			
		} else  {
		break
		
		}
		
		
		}
		
		
		blockData.Tokens = append(blockData.Tokens, tempToken)
		
		
		formatData.Increment(); 
			}
	
	
	
	for formatData.IndexInBounds() {
	
		tempToken = formatData.GetToken(); 
		
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			formatData.AddToFunctionLog("ERROR FillAppend")
			return; 
		}
		
		if tempToken.Type == TokenType.Semicolon {
			break
			
		}
		if tempToken.Type == TokenType.NewLine {
			break
			
		}
		
		formatData.Increment(); 
			}
	
	formatData.AddToFunctionLog("EXIT FillAppend")
	
}
