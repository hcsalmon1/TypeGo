
package formatting

import . "TypeGo/core"



func ProcessIf(formatData *FormatData, token Token) BlockData {
	formatData.SetErrorFunction("ProcessIf"); 
	
	var blockData BlockData  = BlockData{
		NodeType:NodeType.If_Statement, 
		StartingToken:token, 
	}; 
	AddIfCondition(formatData, &blockData); 
	
	if formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after if") == false {
		return blockData; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	//string condition = JoinTextInListOfTokens(blockData.Tokens);
	//if (formatData.IsError()) {
	//formatData.SetError("if condition was null", "getting if condition", token, FormatResult.Internal_Error);
	//return blockData;
	//}
	var ifStatementBlock CodeBlock  = FillBody(formatData); 
	if formatData.IsError() {
		return blockData; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after if") == false {
		return blockData; 
	}
	formatData.Increment(); 
	
	blockData.Block = &ifStatementBlock; 
	
	return blockData; 
}

func ProcessErrorCheck(formatData *FormatData, token Token) BlockData {
	
	formatData.SetErrorFunction("ProcessErrorCheck"); 
	
	var blockData BlockData  = BlockData{
		NodeType:NodeType.Err_Check, 
		StartingToken:token, 
	}; formatData.Increment(); 
	
	if formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after if") == false {
		return blockData; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var errCheckBlock CodeBlock  = FillBody(formatData); if formatData.IsError() {
		return blockData; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after if") == false {
		return blockData; 
	}
	formatData.Increment(); 
	
	blockData.Block = &errCheckBlock; 
	
	return blockData; 
}

func IsDeclarationIf(formatData *FormatData) bool {
	formatData.SetErrorFunction("IsDeclarationIf"); 
	
	var index int  = formatData.TokenIndex; 
	var token_count int  = len(formatData.TokenList)
	
	for i := 0; i < token_count; i++ {
	
		var tempToken Token  = formatData.GetTokenByIndex(index); if tempToken.Type == TokenType.NA {
			formatData.EndOfFileError(tempToken); 
			return false; 
		}
		var token Token  = tempToken; if token.Type == TokenType.Semicolon {
			return true; 
		}
		if token.Type == TokenType.LeftBrace {
			break
			
		}
		index += 1; 
		
	}
	
	return false; 
}

func AddIfCondition(formatData *FormatData, blockData *BlockData) {
	formatData.SetErrorFunction("AddIfCondition"); 
	
	for formatData.IndexInBounds() {
	
		var tempToken Token  = formatData.GetToken(); if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			return; 
		}
		var token Token  = tempToken; if token.Type == TokenType.LeftBrace {
			break
			
		}
		blockData.Tokens = append(blockData.Tokens, token)
		
		formatData.Increment(); 
		
	}
	
	
}

func ProcessElse(formatData *FormatData, token Token) BlockData {
	formatData.SetErrorFunction("ProcessElse"); 
	
	var index int  = formatData.TokenIndex + 1; 
	var blockData BlockData  = BlockData{
		NodeType:NodeType.Else_Statement, 
		StartingToken:token, 
	}; 
	formatData.Increment(); 
	var nextToken Token  = formatData.GetTokenByIndex(index); if formatData.IsValidToken(nextToken) == false {
		formatData.EndOfFileError(token); 
		return blockData; 
	}
	if nextToken.Type == TokenType.If {
		var hasDeclaration bool  = IsDeclarationIf(formatData); if hasDeclaration {
			formatData.UnsupportedFeatureError(token, ""); 
			return blockData; 
		}
		
		AddIfCondition(formatData, &blockData); 
		
	}
	
	if formatData.ExpectType(TokenType.LeftBrace, "Missing expected '{' after else") == false {
		return blockData; 
	}
	formatData.Increment(); 
	
	var ifStatementBlock CodeBlock  = FillBody(formatData); if formatData.IsError() {
		return blockData; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "Missing expected '}' after else") == false {
		return blockData; 
	}
	formatData.Increment(); 
	
	blockData.Block = &ifStatementBlock; 
	
	return blockData; 
}
