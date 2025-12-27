
package formatting

import (
."TypeGo/core"
)


func ProcessIdentifier(formatData *FormatData, firstToken Token) BlockData {
	
	formatData.AddToFunctionLog("ENTER ProcessIdentifier")
	formatData.SetErrorFunction("ProcessIdentifier"); 
	
	var initialIndex int  = formatData.TokenIndex; 
	
	var blockData BlockData  = BlockData{
	Block:nil, 
	NodeType:NodeType.Other, 
	StartingToken:firstToken, 
	Tokens:make([]Token, 0), 
	Variables:make([]Variable, 0), 
	}; 
	
	var lastTokenType IntTokenType  = TokenType.NA; 
	var loopAction IntIdentLoopAction  = IdentLoopAction.Continue; 
	
	for formatData.IndexInBounds() {
	
		var tempToken Token  = formatData.GetToken(); 
		if formatData.IsValidToken(tempToken) == false {
			formatData.EndOfFileError(firstToken); 
			return blockData; 
		}
		var thisToken Token  = tempToken; 
		
		loopAction = IdentifierLoopCode(formatData, thisToken, lastTokenType); 
		if loopAction != IdentLoopAction.Continue {
			break
			
		}
		
		formatData.Increment(); 
		lastTokenType = thisToken.Type; 
			}
	
	
	if loopAction == IdentLoopAction.Error {
		return blockData
	}
	
	if loopAction == IdentLoopAction.Declaration {
		formatData.TokenIndex = initialIndex; 
		formatData.AddToFunctionLog("EXIT ProcessFunction")
		return DeclarationLoop(formatData, firstToken); 
	}
	if loopAction == IdentLoopAction.Other {
		formatData.TokenIndex = initialIndex; 
		formatData.SetErrorProcess("Loop tokens until line end, other, ProcessIdentifier"); 
		LoopTokensUntilLineEnd(formatData, &blockData, true); 
		
	}
	if loopAction == IdentLoopAction.Append {
		formatData.TokenIndex = initialIndex; 
		FillAppend(formatData, &blockData); 
		
	}
	formatData.AddToFunctionLog("EXIT ProcessFunction")
	return blockData; 
}

func IdentifierLoopCode(formatData *FormatData, thisToken Token, lastTokenType IntTokenType) IntIdentLoopAction {
	formatData.SetErrorFunction("IdentifierLoopCode"); 
	
	if IsVarTypeEnum(thisToken.Type) {
		return IdentLoopAction.Declaration; 
	}
	
	switch thisToken.Type {
	
	case TokenType.Identifier:
	
	if lastTokenType == TokenType.Identifier {
	return IdentLoopAction.Declaration; 
		}
	
	return IdentLoopAction.Continue; 
	
	case TokenType.FullStop:
	return IdentLoopAction.Continue; 
	
	case TokenType.Comma:
	return IdentLoopAction.Continue; 
	
	case TokenType.Equals, TokenType.PlusPlus, TokenType.PlusEquals, TokenType.MinusEquals, TokenType.MultiplyEquals, TokenType.DivideEquals, 
	TokenType.ModulusEquals, TokenType.LeftParenthesis, TokenType.RightParenthesis, TokenType.LeftBrace, TokenType.LeftSquareBracket, 
	TokenType.RightBrace, TokenType.NewLine, TokenType.Channel_Setter, TokenType.Colon:
	return IdentLoopAction.Other; 
	
	case TokenType.ColonEquals:
	formatData.UnexpectedTypeError(thisToken, "':=' unsupported outside of for loops, write: 'int value = 10' and not 'value := 10'"); 
	return IdentLoopAction.Error; 
	
	case TokenType.Append:
	return IdentLoopAction.Append; 
	
	default:
	formatData.UnexpectedTypeError(thisToken, "type: " +thisToken.Type.ToString()); 
	return IdentLoopAction.Error; 
		}
	
	
	formatData.SetError(FormatResult.UnexpectedType, "'" +thisToken.Text + "' after identifier")
	return IdentLoopAction.Error
}
