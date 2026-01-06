
package formatting

import . "TypeGo/core"



func DeclarationLoop(format_data *FormatData, firstToken Token) BlockData {
	
	format_data.AddToFunctionLog("ENTER DeclarationLoop")
	format_data.ErrorFunction = "DeclarationLoop"; 
	
	var whileCount int  = 0; 
	var blockData BlockData  = BlockData{
		Block:nil, 
		NodeType:NodeType.Single_Declaration_No_Value, 
		StartingToken:firstToken, 
		Tokens:make([]Token, 0), 
		Variables:make([]Variable, 0), 
	}; var tempVariable Variable  = Variable{
		NameToken:make([]Token, 0), 
		TypeList:make([]Token, 0), 
	}; var lastType IntLastTokenType  = LastTokenType.Null; 
	for format_data.IndexInBounds() {
	
		var result IntLoopAction  = DeclarationInnerLoop(format_data, &whileCount, &blockData, &tempVariable, &lastType); if result == LoopAction.Break {
			break
			
		}
		if result == LoopAction.Return {
			return blockData; 
		}
			}
	
	
	WriteTokens(format_data, &blockData); 
	format_data.AddToFunctionLog("EXIT DeclarationLoop")
	return blockData; 
}

func DeclarationInnerLoop(format_data *FormatData, whileCount *int, blockData *BlockData, tempVariable *Variable, lastType *IntLastTokenType) IntLoopAction {
	
	format_data.AddToFunctionLog("ENTER DeclarationInnerLoop")
	
	if IsInfiniteWhile(whileCount, 10000) == true {
		format_data.Result = FormatResult.Internal_Error; 
		format_data.ErrorDetail = "Infinite white loop in DeclarationInnerLoop"; 
		return LoopAction.Error; 
	}
	
	var previousIndex int  = format_data.TokenIndex; 
	var tempToken Token  = format_data.GetToken(); if tempToken.Type == TokenType.NA {
		format_data.EndOfFileError(tempToken); 
		format_data.AddToFunctionLog("ERROR DeclarationInnerLoop")
		return LoopAction.Return; 
	}
	var loopResult IntLoopAction  = ProcessDeclarationToken(format_data, blockData, tempToken, lastType, tempVariable); if loopResult == LoopAction.Break {
		format_data.AddToFunctionLog("EXIT DeclarationInnerLoop")
		return LoopAction.Break; 
	}
	if loopResult == LoopAction.Return {
		format_data.AddToFunctionLog("EXIT DeclarationInnerLoop")
		return LoopAction.Return; 
	}
	
	format_data.IncrementIfSame(previousIndex); 
	format_data.AddToFunctionLog("EXIT DeclarationInnerLoop")
	return LoopAction.Continue; 
}

func ProcessDeclaration(format_data *FormatData, firstToken Token) BlockData {
	
	format_data.AddToFunctionLog("ENTER ProcessDeclaration")
	
	var isDeclaration bool  = 
	IsVarTypeEnum(firstToken.Type) || 
	firstToken.Type == TokenType.LeftSquareBracket; 
	if isDeclaration {
		format_data.AddToFunctionLog("EXIT ProcessDeclaration")
		return DeclarationLoop(format_data, firstToken); 
	}
	
	if firstToken.Type == TokenType.Multiply {
		if IsPointerDeclaration(format_data) {
			format_data.AddToFunctionLog("EXIT ProcessDeclaration")
			return DeclarationLoop(format_data, firstToken); 
		}
		format_data.AddToFunctionLog("EXIT ProcessDeclaration")
		return FillNonDeclaration(format_data, firstToken); 
	}
	format_data.AddToFunctionLog("EXIT ProcessDeclaration")
	return EmptyBlockData(); 
}

func FillNonDeclaration(format_data *FormatData, firstToken Token) BlockData {
	format_data.AddToFunctionLog("ENTER FillNonDeclaration")
	var blockData BlockData  = BlockData{
		Block:nil, 
		NodeType:NodeType.Other, 
		StartingToken:firstToken, 
		Tokens:make([]Token, 0), 
		Variables:make([]Variable, 0), 
	}; LoopTokensUntilLineEnd(format_data, &blockData, true); 
	format_data.AddToFunctionLog("EXIT FillNonDeclaration")
	return blockData; 
}
