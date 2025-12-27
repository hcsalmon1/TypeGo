
package formatting

import . "TypeGo/core"



func ProcessFor(formatData *FormatData, firstToken Token) BlockData {
	formatData.SetErrorFunction("ProcessFor"); 
	
	var blockData BlockData
	
	blockData.NodeType = NodeType.For_Loop; 
	blockData.StartingToken = firstToken; 
	
	for formatData.IndexInBounds() {
	
		var token Token  = formatData.GetToken(); 
		
		if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			return blockData; 
		}
		
		if token.Type == TokenType.LeftBrace {
			break
			
		}
		blockData.Tokens = append(blockData.Tokens, token)
		
		
		formatData.Increment(); 
			}
	
	
	if formatData.ExpectType(TokenType.LeftBrace, "missing '{' after for loop") == false {
		return blockData; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var forBody CodeBlock  = FillBody(formatData); 
	
	if formatData.IsError() {
		return blockData; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "missing '}' after for loop block") == false {
		return blockData; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	blockData.Block = &forBody; 
	
	return blockData; 
}
