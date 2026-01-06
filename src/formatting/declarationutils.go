
package formatting

import . "TypeGo/core"



func ConvertToNodetypeWithValue(blockData *BlockData) {
	switch blockData.NodeType {
		
		case NodeType.Single_Declaration_No_Value:
			blockData.NodeType = NodeType.Single_Declaration_With_Value; 
			
			
		case NodeType.Multiple_Declarations_No_Value:
			blockData.NodeType = NodeType.Multiple_Declarations_Same_Type_With_Value; 
			
			
		case NodeType.Multiple_Declarations_Same_Type_No_Value:
			blockData.NodeType = NodeType.Multiple_Declarations_Same_Type_With_Value; 
			
			
		default:
			
			
	}
	
	
}

func ProcessAfterNull(formatData *FormatData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	
	formatData.AddToFunctionLog("ENTER ProcessAfterNull")
	if IsVarTypeEnum(token.Type) {
		* lastType = LastTokenType.Vartype; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Multiply {
		* lastType = LastTokenType.Pointer; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.LeftSquareBracket {
		* lastType = LastTokenType.LeftSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Identifier {
		* lastType = LastTokenType.Identifier; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Map {
		* lastType = LastTokenType.Map; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Interface {
		* lastType = LastTokenType.Interface; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterNull")
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unexpected type:" +token.Type.ToString()+" in declaration"); 
	formatData.AddToFunctionLog("ERROR ProcessAfterNull")
	return LoopAction.Error; 
	
}

func ProcessAfterVarType(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	
	formatData.AddToFunctionLog("ENTER ProcessAfterVarType")
	
	if IsVarTypeEnum(token.Type) {
		formatData.UnexpectedTypeError(token, "unexpected vartype in declaration"); 
		formatData.AddToFunctionLog("ERROR ProcessAfterVarType")
		return LoopAction.Return; 
	}
	
	if token.Type == TokenType.Identifier {
		tempVariable.NameToken = append(tempVariable.NameToken, token)
		
		var variable Variable  = Variable{
			NameToken:CopyTokenList(tempVariable.NameToken), 
			TypeList:CopyTokenList(tempVariable.TypeList), 
		}
		blockData.Variables = append(blockData.Variables, variable)
		
		* lastType = LastTokenType.Identifier; 
		tempVariable.SetToDefaults(); 
		formatData.Increment(); 
		formatData.AddToFunctionLog("EXIT ProcessAfterVarType")
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.RightSquareBracket {
		* lastType = LastTokenType.RightSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterVarType")
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.LeftSquareBracket {
		* lastType = LastTokenType.LeftSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterVarType")
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unsupported type:" +token.Type.ToString()+"after variable"); 
	formatData.AddToFunctionLog("ERROR ProcessAfterVarType")
	return LoopAction.Error; 
}

func ProcessAfterIdentifier(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	
	formatData.AddToFunctionLog("ENTER ProcessAfterIdentifier")
	
	if token.Type == TokenType.Semicolon {
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Break; 
	}
	
	var name_count int  = len(tempVariable.NameToken)
	var type_count int  = len(tempVariable.TypeList)
	if token.Type == TokenType.Equals {
		ConvertToNodetypeWithValue(blockData); 
		
		if len(tempVariable.TypeList) != 0 {
			var variable Variable  = Variable{
				NameToken:CopyTokenList(tempVariable.TypeList), 
				TypeList:make([]Token, 0), 
			}
			blockData.Variables = append(blockData.Variables, variable)
			
			tempVariable.TypeList = nil; 
			blockData.NodeType = NodeType.Multiple_Declarations_One_Type_One_Set_Value; 
			
		}
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Break; 
	}
	
	if token.Type == TokenType.Comma {
		if name_count == 0 && type_count != 0 {
			var variable Variable  = Variable{
				NameToken:CopyTokenList(tempVariable.TypeList), 
				TypeList:make([]Token, 0), 
			}
			blockData.Variables = append(blockData.Variables, variable)
			
			tempVariable.SetToDefaults(); 
			
		}
		
		blockData.NodeType = NodeType.Multiple_Declarations_No_Value; 
		* lastType = LastTokenType.Comma; 
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.NewLine {
		if type_count != 0 && name_count == 0 {
			if len(blockData.Variables) != 0 {
				blockData.Variables[0].NameToken = append(blockData.Variables[0].NameToken, tempVariable.TypeList[0]); 
				
			}
			
			tempVariable.TypeList = nil; 
			
		}
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Break; 
	}
	
	if token.Type == TokenType.Identifier {
		tempVariable.NameToken = append(tempVariable.NameToken, token)
		
		var variable Variable  = Variable{
			NameToken:CopyTokenList(tempVariable.NameToken), 
			TypeList:CopyTokenList(tempVariable.TypeList), 
		}
		blockData.Variables = append(blockData.Variables, variable)
		
		* lastType = LastTokenType.Identifier; 
		tempVariable.SetToDefaults(); 
		formatData.Increment(); 
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.FullStop {
		* lastType = LastTokenType.FullStop; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Continue; 
	}
	
	if token.Type == TokenType.RightSquareBracket {
		* lastType = LastTokenType.RightSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		formatData.AddToFunctionLog("EXIT ProcessAfterIdentifier")
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unexpected typed after identifier"); 
	formatData.AddToFunctionLog("ERROR ProcessAfterIdentifier")
	return LoopAction.Error; 
}

func ProcessAfterPointer(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if IsVarTypeEnum(token.Type) {
		* lastType = LastTokenType.Vartype; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.LeftSquareBracket {
		* lastType = LastTokenType.LeftSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Identifier {
		* lastType = LastTokenType.Identifier; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after '*'"); 
	return LoopAction.Error; 
}

func ProcessAfterLeftSqBracket(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if token.Type == TokenType.RightSquareBracket {
		* lastType = LastTokenType.RightSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.IntegerValue {
		* lastType = LastTokenType.IntegerValue; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if IsVarTypeEnum(token.Type) {
		* lastType = LastTokenType.Vartype; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Identifier {
		* lastType = LastTokenType.Identifier; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.FullStop {
		* lastType = LastTokenType.FullStop; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after '['"); 
	return LoopAction.Error; 
}

func ProcessAfterRightSqBracket(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if IsVarTypeEnum(token.Type) {
		* lastType = LastTokenType.Vartype; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.LeftSquareBracket {
		* lastType = LastTokenType.LeftSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Identifier {
		* lastType = LastTokenType.Identifier; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Multiply {
		* lastType = LastTokenType.Pointer; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Map {
		* lastType = LastTokenType.Map; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Interface {
		* lastType = LastTokenType.Interface; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after ']'"); 
	return LoopAction.Error; 
}

func ProcessAfterComma(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if IsVarTypeEnum(token.Type) {
		* lastType = LastTokenType.Vartype; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	if token.Type == TokenType.Identifier {
		* lastType = LastTokenType.Identifier; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after ','"); 
	return LoopAction.Error; 
}

func convertTokenTypeToLastTokenType(token_type IntTokenType) IntLastTokenType {
	
	if IsVarTypeEnum(token_type) {
		return LastTokenType.Vartype
	}
	
	switch token_type {
		
		case TokenType.Identifier:
			return LastTokenType.Identifier
		case TokenType.LeftSquareBracket:
			return LastTokenType.LeftSqBracket
		case TokenType.RightSquareBracket:
			return LastTokenType.RightSqBracket
		case TokenType.LeftParenthesis:
			return LastTokenType.LeftParenth
		case TokenType.RightParenthesis:
			return LastTokenType.RightParenth
		case TokenType.Comma:
			return LastTokenType.Comma
		case TokenType.Semicolon:
			return LastTokenType.Semicolon
		case TokenType.NewLine:
			return LastTokenType.Newline
		case TokenType.IntegerValue:
			return LastTokenType.IntegerValue
		case TokenType.Map:
			return LastTokenType.Map
		case TokenType.Multiply:
			return LastTokenType.Pointer
		case TokenType.FullStop:
			return LastTokenType.FullStop
		case TokenType.Interface:
			return LastTokenType.Interface
		case TokenType.LeftBrace:
			return LastTokenType.LeftBrace
		case TokenType.RightBrace:
			return LastTokenType.RightBrace
	}
	
	return LastTokenType.Null
}

func ProcessAfterFullStop(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	
	var is_valid bool  = 
	token.Type == TokenType.Identifier || 
	token.Type == TokenType.FullStop || 
	token.Type == TokenType.RightSquareBracket
	
	if is_valid == true {
		* lastType = convertTokenTypeToLastTokenType(token.Type); 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after '.'"); 
	return LoopAction.Error; 
}

func ProcessAfterIntegerValue(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if token.Type == TokenType.RightSquareBracket {
		* lastType = LastTokenType.RightSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		//Fmt.PrintlnColor($"\tlast type was ',', add '{token.Text}' to typeList", ConsoleColor.Cyan);
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after integer value"); 
	return LoopAction.Error; 
}

func ProcessAfterMap(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if token.Type == TokenType.LeftSquareBracket {
		* lastType = LastTokenType.LeftSqBracket; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after 'map'"); 
	return LoopAction.Error; 
}

func ProcessAfterInterface(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if token.Type == TokenType.LeftBrace {
		* lastType = LastTokenType.LeftBrace; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after 'interface'"); 
	return LoopAction.Error; 
}

func ProcessAfterLeftBrace(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	if token.Type == TokenType.RightBrace {
		* lastType = LastTokenType.RightBrace; 
		tempVariable.TypeList = append(tempVariable.TypeList, token)
		
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after 'interface'"); 
	return LoopAction.Error; 
}

func ProcessAfterRightBrace(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	
	formatData.AddToFunctionLog("ENTER ProcessAfterRightBrace")
	
	if token.Type == TokenType.Identifier {
		tempVariable.NameToken = append(tempVariable.NameToken, token)
		
		var variable Variable  = Variable{
			NameToken:CopyTokenList(tempVariable.NameToken), 
			TypeList:CopyTokenList(tempVariable.TypeList), 
		}
		blockData.Variables = append(blockData.Variables, variable)
		
		* lastType = LastTokenType.Identifier; 
		tempVariable.SetToDefaults(); 
		formatData.Increment(); 
		formatData.AddToFunctionLog("EXIT ProcessAfterRightBrace")
		return LoopAction.Continue; 
	}
	formatData.UnexpectedTypeError(token, "unsupported type: " +token.Type.ToString()+" after 'interface'"); 
	formatData.AddToFunctionLog("ERROR ProcessAfterRightBrace")
	return LoopAction.Error; 
}

func ProcessDeclarationToken(formatData *FormatData, blockData *BlockData, token Token, lastType *IntLastTokenType, tempVariable *Variable) IntLoopAction {
	switch *lastType {
		
		case LastTokenType.Null:
			return ProcessAfterNull(formatData, token, lastType, tempVariable); 
		case LastTokenType.Identifier:
			return ProcessAfterIdentifier(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.Vartype:
			return ProcessAfterVarType(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.Pointer:
			return ProcessAfterPointer(formatData, blockData, token, lastType, tempVariable); 
			
		case LastTokenType.LeftSqBracket:
			return ProcessAfterLeftSqBracket(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.RightSqBracket:
			return ProcessAfterRightSqBracket(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.Comma:
			return ProcessAfterComma(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.IntegerValue:
			return ProcessAfterIntegerValue(formatData, blockData, token, lastType, tempVariable); 
		case LastTokenType.FullStop:
			return ProcessAfterFullStop(formatData, blockData, token, lastType, tempVariable); 
			
		case LastTokenType.Map:
			return ProcessAfterMap(formatData, blockData, token, lastType, tempVariable); 
			
		case LastTokenType.Interface:
			return ProcessAfterInterface(formatData, blockData, token, lastType, tempVariable); 
			
		case LastTokenType.LeftBrace:
			return ProcessAfterLeftBrace(formatData, blockData, token, lastType, tempVariable); 
			
		case LastTokenType.RightBrace:
			return ProcessAfterRightBrace(formatData, blockData, token, lastType, tempVariable); 
			
		default:
			formatData.UnexpectedTypeError(token, "unsupported last type: " +lastType.ToString()+", this type: " +token.Type.ToString())
			return LoopAction.Error; 
	}
	
	
}

func WriteTokens(formatData *FormatData, blockData *BlockData) {
	
	formatData.AddToFunctionLog("ENTER WriteTokens")
	
	switch blockData.NodeType {
		
		case NodeType.Invalid:
			break
			
		case NodeType.Single_Declaration_With_Value:
			LoopTokensUntilLineEnd(formatData, blockData, true); 
			if blockData.Validate(formatData) == false {
				formatData.AddToFunctionLog("ERROR WriteTokens")
				return 
			}
			
			
		case NodeType.Single_Declaration_No_Value:
			LoopTokensUntilLineEnd(formatData, blockData, true); 
			
			
		case NodeType.Multiple_Declarations_One_Type_One_Set_Value:
			LoopTokensUntilLineEnd(formatData, blockData, true); 
			
			
		case NodeType.Multiple_Declarations_No_Value:
			LoopTokensUntilLineEnd(formatData, blockData, true); 
			
			
		case NodeType.Multiple_Declarations_With_Value:
			formatData.UnsupportedFeatureError(blockData.StartingToken, "Multiple declarations with value, not supported yet"); 
			formatData.AddToFunctionLog("ERROR WriteTokens")
			return 
			
		case NodeType.Multiple_Declarations_Same_Type_No_Value:
			formatData.UnsupportedFeatureError(blockData.StartingToken, "Multiple same type declarations with no value, not supported yet"); 
			formatData.AddToFunctionLog("ERROR WriteTokens")
			return 
			
		case NodeType.Multiple_Declarations_Same_Type_With_Value:
			LoopTokensUntilLineEnd(formatData, blockData, true); 
			
			
		default:
			break
			
	}
	
	formatData.AddToFunctionLog("EXIT WriteTokens")
	
}
