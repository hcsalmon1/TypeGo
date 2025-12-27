
package formatting

import . "TypeGo/core"



func ProcessPackage(formatData *FormatData, globalBlock *CodeBlock, packageToken Token) {
	
	formatData.SetErrorFunction("ProcessPackage"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.NodeType = NodeType.Package; 
	blockData.StartingToken = packageToken; 
	
	blockData.Tokens = append(blockData.Tokens, packageToken)
	
	
	var packageNameToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(packageNameToken) == false {
		formatData.EndOfFileError(packageToken); 
		return; 
	}
	if packageNameToken.Type != TokenType.Identifier {
		formatData.MissingExpectedTypeError(packageNameToken, "Missing identifier in package"); 
		return; 
	}
	
	var newLineToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(newLineToken) == false {
		formatData.EndOfFileError(packageToken); 
		return; 
	}
	blockData.Tokens = append(blockData.Tokens, packageNameToken)
	
	blockData.Tokens = append(blockData.Tokens, newLineToken)
	
	formatData.Increment(); 
	
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}

func MultiImportInnerLoop(formatData *FormatData, blockData *BlockData, token Token, lastTokenType *IntTokenType) IntLoopAction {
	
	formatData.SetErrorFunction("MultiImportInnerLoop"); 
	
	var isValid bool
	
	switch token.Type {
	
	case TokenType.StringValue:
	
	isValid = 
	* lastTokenType == TokenType.Identifier || 
	* lastTokenType == TokenType.FullStop || 
	* lastTokenType == TokenType.NewLine || 
	* lastTokenType == TokenType.Comment || 
	* lastTokenType == TokenType.EndComment; 
	
	if isValid {
	blockData.Tokens = append(blockData.Tokens, token); 
	break; 
		}
	formatData.UnexpectedTypeError(token, "invalid import lasttype" +lastTokenType.ToString()+" token: " +token.Text); 
	return LoopAction.Return; 
	
	case TokenType.Identifier:
	isValid = 
	* lastTokenType == TokenType.NewLine || 
	* lastTokenType == TokenType.EndComment; 
	
	if isValid {
	blockData.Tokens = append(blockData.Tokens, token); 
	break; 
		}
	formatData.UnexpectedTypeError(token, "invalid import lasttype " +lastTokenType.ToString()+" token: " +token.Text); 
	return LoopAction.Return; 
	
	case TokenType.Comment:
	isValid = 
	* lastTokenType == TokenType.NewLine || 
	* lastTokenType == TokenType.EndComment; 
	
	if isValid {
	blockData.Tokens = append(blockData.Tokens, token); 
	break; 
		}
	formatData.UnexpectedTypeError(token, "invalid import lasttype " +lastTokenType.ToString()+" token: " +token.Text); 
	break; 
	
	case TokenType.EndComment:
	break; 
	
	case TokenType.NewLine:
	isValid = 
	* lastTokenType == TokenType.StringValue || 
	* lastTokenType == TokenType.LeftParenthesis || 
	* lastTokenType == TokenType.NewLine || 
	* lastTokenType == TokenType.Comment || 
	* lastTokenType == TokenType.EndComment; 
	
	if isValid {
	blockData.Tokens = append(blockData.Tokens, token); 
	break; 
		}
	formatData.UnexpectedTypeError(token, "invalid import lasttype " +lastTokenType.ToString()+" token: " +token.Text); 
	return LoopAction.Return; 
	
	case TokenType.FullStop:
	blockData.Tokens = append(blockData.Tokens, token); 
	return LoopAction.Continue; 
	
	case TokenType.RightParenthesis:
	blockData.Tokens = append(blockData.Tokens, token); 
	formatData.Increment(); 
	return LoopAction.Break; 
	
	default:
	formatData.UnexpectedTypeError(token, "invalid import lasttype " +lastTokenType.ToString()+" token: " +token.Text); 
	return LoopAction.Return; 
		}
	
	return LoopAction.Continue; 
}

func ProcessMultiImport(formatData *FormatData, globalBlock *CodeBlock, blockData *BlockData, leftParenthToken Token) {
	
	formatData.SetErrorFunction("ProcessMultiImport"); 
	
	blockData.Tokens = append(blockData.Tokens, leftParenthToken)
	
	formatData.Increment(); 
	var lastTokenType IntTokenType  = TokenType.LeftParenthesis; 
	blockData.NodeType = NodeType.Multi_Line_Import; 
	
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var tempToken Token  = formatData.GetToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(leftParenthToken); 
			return; 
		}
		var token Token  = tempToken; 
		
		var result IntLoopAction  = MultiImportInnerLoop(formatData, blockData, token, &lastTokenType); 
		if result == LoopAction.Return {
			return; 
		}
		if result == LoopAction.Break {
			break
			
		}
		
		lastTokenType = token.Type; 
		formatData.IncrementIfSame(previousIndex); 
			}
	
	
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, *blockData)
	
	
}

func ProcessImport(formatData *FormatData, globalBlock *CodeBlock, packageToken Token) {
	formatData.SetErrorFunction("ProcessImport"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.StartingToken = packageToken; 
	blockData.Tokens = append(blockData.Tokens, packageToken)
	
	
	var tempToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(packageToken); 
		return; 
	}
	var nextToken Token  = tempToken; 
	
	if nextToken.IsType(TokenType.LeftParenthesis) {
		ProcessMultiImport(formatData, globalBlock, &blockData, nextToken); 
		return; 
	}
	ProcessSingleImport(formatData, globalBlock, &blockData, &nextToken); 
	
}

func ProcessSingleImport(formatData *FormatData, globalBlock *CodeBlock, blockData *BlockData, nextToken *Token) {
	
	formatData.SetErrorFunction("ProcessSingleImport"); 
	
	var tempToken Token  = EmptyToken()
	
	
	blockData.NodeType = NodeType.Single_Import; 
	//Is Alias
	if nextToken.Type == TokenType.Identifier {
		blockData.NodeType = NodeType.Single_Import_With_Alias; 
		blockData.Tokens = append(blockData.Tokens, *nextToken)
		
		
		tempToken = formatData.GetNextToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(* nextToken); 
			return; 
		}
		* nextToken = tempToken; 
		
	}
	if nextToken.Type == TokenType.FullStop {
		blockData.Tokens = append(blockData.Tokens, *nextToken)
		
		var spaceToken Token  = Token{
		Text:" ", 
		Type:TokenType.StringValue, 
		LineNumber:0, 
		CharNumber:0, 
		}
		
		blockData.Tokens = append(blockData.Tokens, spaceToken)
		
		
		tempToken = formatData.GetNextToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(* nextToken); 
			return; 
		}
		* nextToken = tempToken; 
		
	}
	//Import name
	if nextToken.Type == TokenType.StringValue {
		blockData.Tokens = append(blockData.Tokens, *nextToken)
		
		tempToken = formatData.GetNextToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(* nextToken); 
			return; 
		}
		* nextToken = tempToken; 
		
	}
	//new line
	if (* nextToken).Type == TokenType.NewLine {
		blockData.Tokens = append(blockData.Tokens, *nextToken)
		
		globalBlock.BlockDataList = append(globalBlock.BlockDataList, *blockData)
		
		return; 
	}
	formatData.UnexpectedTypeError(* nextToken, "unexpected type in import"); 
	
}

func ProcessStruct(formatData *FormatData, globalBlock *CodeBlock, structToken Token) {
	formatData.SetErrorFunction("ProcessStruct"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.StartingToken = structToken; 
	blockData.NodeType = NodeType.Struct_Declaration; 
	
	//Add struct name
	var tempToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var structNameToken Token  = tempToken; 
	blockData.Tokens = append(blockData.Tokens, structNameToken)
	
	
	if formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in struct declaration") == false {
		return; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	formatData.AddToProcessLog("Struct body " +structNameToken.Text)
	formatData.IncreaseLogIndent()
	
	var structBlock CodeBlock  = FillStructBody(formatData); 
	
	formatData.DecreaseLogIndent()
	formatData.AddToProcessLog("} " +structNameToken.Text)
	
	if formatData.IsError() {
		return; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in struct declaration") == false {
		return; 
	}
	formatData.Increment(); 
	blockData.Block = &structBlock; 
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}

func ProcessEnum(formatData *FormatData, globalBlock *CodeBlock, structToken Token) {
	formatData.SetErrorFunction("ProcessEnum"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.StartingToken = structToken; 
	blockData.NodeType = NodeType.Enum_Declaration; 
	
	//Add struct name
	var tempToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var enumNameToken Token  = tempToken; 
	blockData.Tokens = append(blockData.Tokens, enumNameToken)
	
	
	if formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in enum declaration") == false {
		return; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var enumBlock CodeBlock  = FillEnumBody(formatData); 
	if formatData.IsError() {
		return; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in enum declaration") == false {
		return; 
	}
	formatData.Increment(); 
	blockData.Block = &enumBlock; 
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}

func ProcessEnumstruct(formatData *FormatData, globalBlock *CodeBlock, structToken Token) {
	formatData.SetErrorFunction("ProcessEnumStruct"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.StartingToken = structToken; 
	blockData.NodeType = NodeType.Enum_Struct_Declaration; 
	
	//Add struct name
	var tempToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	
	var enumNameToken Token  = tempToken; 
	blockData.Tokens = append(blockData.Tokens, enumNameToken)
	
	
	tempToken = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	
	if tempToken.Type == TokenType.Colon {
		blockData.NodeType = NodeType.Enum_Struct_Declaration_With_Alias; 
		
		tempToken = formatData.GetNextToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			return; 
		}
		if tempToken.Type != TokenType.Identifier {
			formatData.MissingExpectedTypeError(tempToken, "Missing alias identifier in enumstruct"); 
			return; 
		}
		blockData.VarName = tempToken.Text; 
		formatData.Increment(); 
		
	}
	if formatData.ExpectType(TokenType.LeftBrace, "missing expected '{' in enum declaration") == false {
		return; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var enumBlock CodeBlock  = FillEnumBody(formatData); 
	if formatData.IsError() {
		return; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in enum declaration") == false {
		return; 
	}
	formatData.Increment(); 
	blockData.Block = &enumBlock; 
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}

func FillEnumBody(formatData *FormatData) CodeBlock {
	
	formatData.SetErrorFunction("FillEnumBody"); 
	
	var block CodeBlock  = CodeBlock{
	BlockDataList:make([]BlockData, 0), 
	}; 
	
	var tempBlockData BlockData
	tempBlockData.NodeType = NodeType.Enum_Variable
	
	var whileCount int  = 0; 
	
	for formatData.IndexInBounds() {
	
		whileCount += 1; 
		if whileCount > 10000 {
			formatData.Result = FormatResult.Infinite_While_Loop; 
			return block; 
		}
		
		var token Token  = formatData.GetToken(); 
		if formatData.IsValidToken(token) == false {
			formatData.EndOfFileError(token); 
			return block; 
		}
		if token.Type == TokenType.Fn {
			ProcessFunction(formatData, &block.MethodList, token); 
			formatData.Increment(); 
			continue 
			
		}
		if token.Type == TokenType.Comment {
			formatData.Increment()
			continue 
			
		}
		if token.Type == TokenType.Comma {
			formatData.Increment()
			continue 
			
		}
		if token.Type == TokenType.NewLine {
			block.BlockDataList = append(block.BlockDataList, tempBlockData)
			
			tempBlockData = BlockData {}
			tempBlockData.NodeType = NodeType.Enum_Variable
			formatData.Increment(); 
			continue 
			
			
		}
		if token.Type == TokenType.RightBrace {
			break
			
		}
		
		tempBlockData.Tokens = append(tempBlockData.Tokens, token)
		
		
		formatData.Increment(); 
			}
	
	
	return block; 
}

func ProcessConstant(formatData *FormatData, globalBlock *CodeBlock, constToken Token) {
	formatData.SetErrorFunction("ProcessConstant"); 
	
	var blockData BlockData  = BlockData{
	Block:nil, 
	NodeType:NodeType.Constant_Global_Variable, 
	StartingToken:constToken, 
	Tokens:make([]Token, 0), 
	Variables:make([]Variable, 0), 
	}; 
	
	formatData.Increment(); 
	var tempToken Token  = formatData.GetToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	if tempToken.Type == TokenType.Identifier {
		blockData.Tokens = append(blockData.Tokens, tempToken)
		
		formatData.Increment(); 
		LoopValue(formatData, globalBlock, blockData); 
		return; 
	}
	if IsVarTypeEnum(tempToken.Type) == true {
		blockData.NodeType = NodeType.Constant_Global_Variable_With_Type; 
		var variable Variable  = Variable{
		TypeList:make([]Token, 0), 
		NameToken:make([]Token, 0), 
		}
		
		variable.TypeList = append(variable.TypeList, tempToken)
		
		formatData.Increment(); 
		tempToken = formatData.GetToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(tempToken); 
			return; 
		}
		if tempToken.Type != TokenType.Identifier {
			formatData.UnexpectedTypeError(tempToken, "missing expended identifier in constant variable"); 
			
		}
		
		variable.NameToken = append(variable.NameToken, tempToken)
		
		blockData.Variables = append(blockData.Variables, variable)
		
		formatData.Increment(); 
		
		LoopValue(formatData, globalBlock, blockData); 
		
		
	}
	
}

func LoopValue(formatData *FormatData, globalBlock *CodeBlock, blockData BlockData) {
	formatData.SetErrorFunction("LoopValue"); 
	
	for formatData.IndexInBounds() {
	
		var token Token  = formatData.GetToken(); 
		if formatData.IsValidToken(token) == false {
			formatData.EndOfFileError(token); 
			return; 
		}
		if token.Type == TokenType.NewLine {
			break
			
		}
		blockData.Tokens = append(blockData.Tokens, token)
		
		formatData.Increment(); 
			}
	
	
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}

func Process(formatData *FormatData, globalBlock *CodeBlock, interfaceToken Token) {
	formatData.SetErrorFunction("ProcessInterface"); 
	
	var blockData BlockData
	blockData.Tokens = make([]Token, 0)
	blockData.StartingToken = interfaceToken; 
	blockData.NodeType = NodeType.Interface_Declaration; 
	
	//Add struct name
	var tempToken Token  = formatData.GetNextToken(); 
	if formatData.IsValidToken(tempToken) == false {
		formatData.EndOfFileError(tempToken); 
		return; 
	}
	var structNameToken Token  = tempToken; 
	blockData.Tokens = append(blockData.Tokens, structNameToken)
	
	
	if formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in interface declaration") == false {
		return; 
	}
	formatData.Increment(); 
	formatData.IncrementIfNewLine(); 
	
	var structBlock CodeBlock  = FillInterfaceBody(formatData); 
	if formatData.IsError() {
		return; 
	}
	if formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in interface declaration") == false {
		return; 
	}
	formatData.Increment(); 
	blockData.Block = &structBlock; 
	globalBlock.BlockDataList = append(globalBlock.BlockDataList, blockData)
	
	
}
