
package formatting

import . "TypeGo/core"



func FormatInterfaceVar(formatData *FormatData, firstToken Token) BlockData {
	
	formatData.SetErrorFunction("FormatInterfaceVar"); 
	
	var blockData BlockData  = BlockData{
	Block:nil, 
	NodeType:NodeType.Single_Declaration_No_Value, 
	StartingToken:firstToken, 
	Tokens:make([]Token, 0), 
	Variables:make([]Variable, 0), 
	}; 
	
	var variable Variable
	variable.TypeList = make([]Token, 0)
	variable.TypeList = append(variable.TypeList, firstToken)
	
	
	formatData.Increment(); 
	var tempToken Token  = formatData.GetToken(); 
	
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(firstToken); 
		return blockData; 
	}
	
	if tempToken.Type != TokenType.LeftBrace {
		formatData.MissingExpectedTypeError(tempToken, "Missing expected '{' in interface declaration"); 
		return blockData; 
	}
	
	variable.TypeList = append(variable.TypeList, tempToken)
	
	
	formatData.Increment(); 
	tempToken = formatData.GetToken(); 
	
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(firstToken); 
		return blockData; 
	}
	
	if tempToken.Type != TokenType.RightBrace {
		formatData.MissingExpectedTypeError(tempToken, "Missing expected '}' in interface declaration, advanced local interfaces not supported yet"); 
		return blockData; 
	}
	
	variable.TypeList = append(variable.TypeList, tempToken)
	
	
	formatData.Increment(); 
	
	tempToken = formatData.GetToken(); 
	
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(firstToken); 
		return blockData; 
	}
	
	if tempToken.Type != TokenType.Identifier {
		formatData.MissingExpectedTypeError(tempToken, "Missing expected identifier in interface declaration"); 
		return blockData; 
	}
	
	variable.NameToken = append(variable.NameToken, tempToken)
	
	
	formatData.Increment(); 
	
	tempToken = formatData.GetToken(); 
	
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(firstToken); 
		return blockData; 
	}
	
	if tempToken.Type == TokenType.Equals {
		blockData.NodeType = NodeType.Single_Declaration_With_Value; 
		LoopTokensUntilLineEnd(formatData, &blockData, true); 
		
	}
	
	blockData.Variables = append(blockData.Variables, variable)
	
	
	return blockData; 
}
