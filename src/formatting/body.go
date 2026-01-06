
package formatting

import . "TypeGo/core"



func ProcessTokenInBody(formatData *FormatData, block *CodeBlock, token Token) {
	
	formatData.AddToFunctionLog("ENTER ProcessTokenInBody")
	formatData.ErrorFunction = "ProcessTokenInBody"
	
	var blockData BlockData
	
	switch token.Type {
		
		
		case TokenType.Chan:
			blockData = ProcessChannel(formatData, token); 
			
			
		case TokenType.Int, TokenType.Int8, TokenType.Int16, TokenType.Int32, TokenType.Int64, TokenType.Uint, 
		TokenType.Uint8, TokenType.Uint16, TokenType.Uint32, TokenType.Uint64, TokenType.Float32, TokenType.Float64, 
		TokenType.String, TokenType.Byte, TokenType.Rune, TokenType.Bool, TokenType.LeftSquareBracket, TokenType.Error, 
		TokenType.Multiply, TokenType.Map, TokenType.Interface:
			
			blockData = ProcessDeclaration(formatData, token); 
			
			
		case TokenType.If:
			blockData = ProcessIf(formatData, token); 
			
			
		case TokenType.Else:
			blockData = ProcessElse(formatData, token); 
			
			
		case TokenType.For:
			blockData = ProcessFor(formatData, token); 
			
			
		case TokenType.Return:
			blockData = ProcessReturn(formatData); 
			
			
		case TokenType.Break:
			blockData = BlockData {}; 
			blockData.NodeType = NodeType.Break; 
			formatData.IncrementTwice(); 
			
			
		case TokenType.Continue:
			blockData = BlockData {}; 
			blockData.NodeType = NodeType.Continue; 
			blockData.Tokens = append(blockData.Tokens, token); 
			blockData.StartingToken = token
			formatData.IncrementTwice(); 
			
			
		case TokenType.Defer:
			blockData = ProcessDefer(formatData, &token); 
			
			
			
		case TokenType.Goto:
			blockData = ProcessDefer(formatData, &token); 
			
			
			
		case TokenType.Identifier:
			blockData = ProcessIdentifier(formatData, token); 
			
			
		case TokenType.LeftParenthesis, TokenType.Semicolon:
			return; 
		case TokenType.Tab:
			return; 
		case TokenType.MultiLineStart, TokenType.MultiLineEnd, TokenType.Comment:
			blockData = ProcessComment(formatData); 
			
			
		case TokenType.EndComment:
			return; 
		case TokenType.NewLine:
			blockData = BlockData {}; 
			blockData.NodeType = NodeType.NewLine; 
			
			
		case TokenType.ErrReturn:
			blockData = ProcessErrReturn(formatData, token); 
			
			
		case TokenType.Var:
			formatData.Result = FormatResult.UnexpectedType; 
			formatData.ErrorDetail = "'var' isn't ever used in TypeGo. Correct syntax: 'int number = 10', not 'var number int = 10'"; 
			formatData.ErrorToken = token; 
			formatData.AddToFunctionLog("ERROR ProcessTokenInBody")
			return; 
		case TokenType.Go:
			blockData = ProcessGo(formatData, token); 
			
			
		case TokenType.Switch:
			blockData = ProcessSwitch(formatData); 
			
			
		case TokenType.ErrCheck:
			blockData = ProcessErrorCheck(formatData, token); 
			
			
		default:
			formatData.UnsupportedFeatureError(token, "unexpected token " +token.Text); 
			formatData.AddToFunctionLog("ERROR ProcessTokenInBody")
			return; 
	}
	
	
	if formatData.IsError() {
		return; 
	}
	
	block.BlockDataList = append(block.BlockDataList, blockData)
	
	formatData.AddToFunctionLog("EXIT ProcessTokenInBody")
	
	
}

func FillBody(formatData *FormatData) CodeBlock {
	
	formatData.AddToFunctionLog("ENTER FillBody")
	formatData.ErrorFunction = "FillBody"; 
	
	var block CodeBlock  = CodeBlock{
		BlockDataList:make([]BlockData, 0), 
	}; 
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			formatData.AddToFunctionLog("ERROR ProcessTokenInBody")
			return block; 
		}
		
		if token.Type == TokenType.RightBrace {
			break
		}
		
		ProcessTokenInBody(formatData, &block, token); 
		if formatData.IsError() {
			return block; 
		}
		
		formatData.IncrementIfSame(previousIndex); 
			}
	
	formatData.AddToFunctionLog("EXIT ProcessTokenInBody")
	return block; 
}

func FillStructBody(formatData *FormatData) CodeBlock {
	
	formatData.AddToFunctionLog("ENTER FillStructBody")
	formatData.ErrorFunction = "FillStructBody"; 
	
	var block CodeBlock  = CodeBlock{
		BlockDataList:make([]BlockData, 0), 
		MethodList:make([]Function, 0), 
	}; 
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			formatData.AddToFunctionLog("ERROR FillStructBody")
			return block; 
		}
		if token.Type == TokenType.Equals {
			formatData.UnexpectedTypeError(token, "Can't set a value in a struct definition")
			return block
		}
		if token.Type == TokenType.RightBrace {
			break
		}
		if token.Type == TokenType.Fn {
			ProcessFunction(formatData, &block.MethodList, token); 
			formatData.Increment(); 
			continue
			
		}
		
		ProcessTokenInBody(formatData, &block, token); 
		if formatData.IsError() {
			return block; 
		}
		
		formatData.IncrementIfSame(previousIndex); 
			}
	
	formatData.AddToFunctionLog("EXIT FillStructBody")
	return block; 
}

func FillInterfaceBody(formatData *FormatData) CodeBlock {
	
	formatData.AddToFunctionLog("ENTER FillInterfaceBody")
	formatData.ErrorFunction = "FillInterfaceBody"; 
	
	var block CodeBlock  = CodeBlock{
		BlockDataList:make([]BlockData, 0), 
		MethodList:make([]Function, 0), 
	}; 
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			formatData.AddToFunctionLog("ERROR FillStructBody")
			return block; 
		}
		
		if token.Type == TokenType.RightBrace {
			break
		}
		if token.Type == TokenType.NewLine {
			formatData.Increment(); 
			continue
			
		}
		ProcessInterfaceFunction(formatData, &block.MethodList, token); 
		formatData.IncrementIfSame(previousIndex); 
			}
	
	formatData.AddToFunctionLog("EXIT FillStructBody")
	return block; 
}

func FillSwitchBody(formatData *FormatData) CodeBlock {
	
	formatData.ErrorFunction = "FillSwitchBody"; 
	
	var block CodeBlock  = CodeBlock{
		BlockDataList:make([]BlockData, 0), 
	}; 
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			return block; 
		}
		
		if token.Type == TokenType.RightBrace {
			break
		}
		
		FillSwitchCase(formatData, &block); 
		if formatData.IsError() {
			return block; 
		}
		
		formatData.IncrementIfSame(previousIndex); 
			}
	
	return block; 
	
}

func FillSwitchCase(formatData *FormatData, block *CodeBlock) {
	//Codeblock:
	//  []BlockData BlockDataList
	//  []Function MethodList
	
	//BlockData:
	
	//  IntNodeType NodeType
	//  []Token Tokens
	//  Token StartingToken
	//  *CodeBlock Block
	//  []Variable Variables
	//  string VarName
	
	var first_token Token  = formatData.GetToken(); if first_token.Type == TokenType.NA {
		formatData.EndOfFileError(first_token)
		return 
	}
	
	var blockData BlockData  = BlockData{
		NodeType:NodeType.Switch_Case, 
		StartingToken:first_token, 
		Tokens:make([]Token, 0), 
	}
	
	//finding case condition
	for formatData.IndexInBounds() {
	
		var index_before int  = formatData.TokenIndex
		
		var token Token  = formatData.GetToken()
		if first_token.Type == TokenType.NA {
			formatData.EndOfFileError(first_token)
			return 
		}
		
		blockData.Tokens = append(blockData.Tokens, token)
		
		
		formatData.IncrementIfSame(index_before)
		if token.Type == TokenType.Colon {
			break
		}
			}
	
	
	var case_body_block CodeBlock  = FillCaseBody(formatData)
	blockData.Block = &case_body_block
	
	block.BlockDataList = append(block.BlockDataList, blockData)
	
}

func FillCaseBody(formatData *FormatData) CodeBlock {
	
	formatData.ErrorFunction = "FillCaseBody"; 
	
	var block CodeBlock  = CodeBlock{
		BlockDataList:make([]BlockData, 0), 
	}; 
	for formatData.IndexInBounds() {
	
		var previousIndex int  = formatData.TokenIndex; 
		var token Token  = formatData.GetToken(); if token.Type == TokenType.NA {
			formatData.EndOfFileError(token); 
			return block; 
		}
		
		var is_ending_token bool  = 
		token.Type == TokenType.RightBrace || 
		token.Type == TokenType.Default || 
		token.Type == TokenType.Case
		
		if is_ending_token {
			break
		}
		
		ProcessTokenInBody(formatData, &block, token); 
		if formatData.IsError() {
			return block; 
		}
		
		formatData.IncrementIfSame(previousIndex); 
			}
	
	return block; 
}
