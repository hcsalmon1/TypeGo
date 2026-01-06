
package converting

import (
."TypeGo/core"
)


func PrintConstant(convertData *ConvertData, blockData *BlockData, nestCount int) {
	convertData.AppendString("const "); 
	if len(blockData.Variables) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "no variable tokens in constant declaration"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	var variable Variable  = blockData.Variables[0]; if variable.NameToken == nil {
		convertData.ConvertResult = ConvertResult.Null_Token; 
		convertData.ErrorDetail = "null name token in constant declaration"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	convertData.AppendToken(variable.NameToken[0]); 
	convertData.AppendChar(' '); 
	if len(variable.TypeList) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "no variable type in constant declaration"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	convertData.AppendString(variable.TypeList[0].Text); 
	PrintTokens(convertData, blockData, nestCount); 
	
}

func ProcessSingleDeclarationNoValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	var string_builder []byte  = make([]byte, 0)
	
	var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "no variables in single declaration"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for varIndex := 0; varIndex < len(variables); varIndex++ {
	
		var varTypeTokenList []Token  = variables[varIndex].TypeList; 
		if varTypeTokenList == nil {
			convertData.UnexpectedTypeError(blockData.StartingToken, "Type is invalid in single declaration"); 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); if varTypeAsText == "" {
			convertData.UnexpectedTypeError(blockData.StartingToken, "Type is invalid in single declaration"); 
			return; 
		}
		
		var varNameList []Token  = variables[varIndex].NameToken; 
		for nameTokenIndex := 0; nameTokenIndex < len(varNameList); nameTokenIndex++ {
		
			var nameToken Token  = variables[varIndex].NameToken[nameTokenIndex]; if nameToken.Type == TokenType.NA {
				convertData.UnexpectedTypeError(blockData.StartingToken, "Variable name is invalid in single declaration"); 
				return; 
			}
			
			if nameTokenIndex != 0 {
				string_builder = append(string_builder, ',')
				
				string_builder = append(string_builder, ' ')
				
				
			}
			for i := 0; i < len(nameToken.Text); i++ {
			
				string_builder = append(string_builder, nameToken.Text[i])
				
				
			}
			
			
			
		}
		
		
		convertData.AppendString("var " +string(string_builder)+" " +varTypeAsText); 
		
	}
	
	
}

func ProcessChannelDeclarationNoValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "no variables in single declaration"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for i := 0; i < len(variables); i++ {
	
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Unexpected_Type; 
			convertData.ErrorDetail = "Type is invalid in single declaration"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); if varTypeAsText == "" {
			convertData.ConvertResult = ConvertResult.Unexpected_Type; 
			convertData.ErrorDetail = "Type is invalid in single declaration"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Unexpected_Type; 
			convertData.ErrorDetail = "Variable name is invalid in single declaration"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		convertData.AppendString("var " +nameToken.Text + " chan " +varTypeAsText); 
		
		
	}
	
	
}

func ProcessChannelDeclarationWithValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	var tokens []Token  = blockData.Tokens; var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No variables in single declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(tokens) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "No tokens in single declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for i := 0; i < len(variables); i++ {
	
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "invalid var type in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList)
		if varTypeAsText == "" {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "invalid var type text in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Text == "" {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "No variable name in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		convertData.AppendString("var " +nameToken.Text + " chan " +varTypeAsText + " "); 
		PrintTokensDeclaration(convertData, blockData, nestCount, false, varTypeAsText); 
		
	}
	
	
}

func ProcessSingleDeclarationWithValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	var tokens []Token  = blockData.Tokens; var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No variables in single declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(tokens) == 0 {
		convertData.ConvertResult = ConvertResult.No_Token_In_Node; 
		convertData.ErrorDetail = "No tokens in single declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for i := 0; i < len(variables); i++ {
	
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "invalid var type in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); if varTypeAsText == "" {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "invalid var type text in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "No variable name in single declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var should_add_type bool  = true
		
		if len(varTypeTokenList) > 2 {
			var is_inferred_array bool  = 
			varTypeTokenList[0].Type == TokenType.LeftSquareBracket && 
			varTypeTokenList[1].Type == TokenType.FullStop
			
			if is_inferred_array == true {
				should_add_type = false
				
			}
			
		}
		
		if should_add_type == true {
			convertData.AppendString("var " +nameToken.Text + " " +varTypeAsText + " "); 
			
		} else  {
			convertData.AppendString("var " +nameToken.Text + " "); 
			
		}
		
		PrintTokensDeclaration(convertData, blockData, nestCount, false, varTypeAsText); 
		
	}
	
	//convertData.NewLineWithTabs()
	
}

func ProcessMultipleDeclarationNoValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	var tokens []Token  = blockData.Tokens; var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No variables in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No Tokens in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for i := 0; i < len(variables); i++ {
	
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = " var type is invalid in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); if varTypeAsText == "" {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "var type text is invalid in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varNameTokenList []Token  = variables[i].NameToken; 
		convertData.AppendChar('v'); 
		convertData.AppendChar('a'); 
		convertData.AppendChar('r'); 
		convertData.AppendChar(' '); 
		
		for varNameIndex := 0; varNameIndex < len(varNameTokenList); varNameIndex++ {
		
			var nameToken Token  = varNameTokenList[varNameIndex]; if nameToken.Text == "" {
				convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
				convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
				convertData.ErrorToken = blockData.StartingToken; 
				return; 
			}
			
			if varNameIndex != 0 {
				convertData.AppendChar(','); 
				convertData.AppendChar(' '); 
				
			}
			convertData.AppendString(nameToken.Text); 
			
		}
		
		
		convertData.AppendString(" " +varTypeAsText + "\n\t"); 
		
	}
	
	
	convertData.AppendChar(' '); 
	PrintTokensDeclaration(convertData, blockData, nestCount, false, ""); 
	
}

func ProcessMultipleDeclarationWithValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	var tokens []Token  = blockData.Tokens; var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No variables in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No Tokens in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	for i := 0; i < len(variables); i++ {
	
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = " var type is invalid in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); if varTypeAsText == "" {
			//convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
			//convertData.ErrorDetail = "var type text is invalid in multiple declaration with value";
			//convertData.ErrorToken = blockData.StartingToken;
			continue
			
		}
		
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		convertData.AppendString("var " +nameToken.Text + " " +varTypeAsText + "\n\t"); 
		
	}
	
	
	for i := 0; i < len(variables); i++ {
	
		if len(variables[i].NameToken) == 0 {
			continue
			
		}
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		if i != 0 {
			convertData.AppendChar(','); 
			convertData.AppendChar(' '); 
			
		}
		convertData.AppendString(nameToken.Text); 
		
	}
	
	convertData.AppendChar(' '); 
	PrintTokensDeclaration(convertData, blockData, nestCount, false, ""); 
	
}

func ProcessMultipleDeclarationWithSetValue(convertData *ConvertData, blockData *BlockData, nestCount int) {
	var tokens []Token  = blockData.Tokens; var variables []Variable  = blockData.Variables; if len(variables) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No variables in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	if len(tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
		convertData.ErrorDetail = "No Tokens in multiple declaration with value"; 
		convertData.ErrorToken = blockData.StartingToken; 
		return; 
	}
	
	for i := 0; i < len(variables); i++ {
	
		if i != 0 {
			convertData.AppendChar(';'); 
			convertData.AppendChar(' '); 
			
		}
		
		var varTypeTokenList []Token  = variables[i].TypeList; 
		if varTypeTokenList == nil {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = " var type is invalid in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); 
		if len(variables[i].NameToken) == 0 {
			continue
			
		}
		
		var nameToken Token  = variables[i].NameToken[0]; if nameToken.Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		if varTypeAsText == "" {
			continue
			
		}
		
		convertData.AppendString("var " +nameToken.Text + " " +varTypeAsText); 
		
	}
	
	convertData.NewLineWithTabs(); 
	
	for i := 0; i < len(variables); i++ {
	
		var nameToken Token  = EmptyToken(); 
		if len(variables[i].NameToken) == 0 {
			var varTypeTokenList []Token  = variables[i].TypeList; var varTypeAsText string  = JoinTextInListOfTokens( &varTypeTokenList); 
			if varTypeAsText == "" {
				convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
				convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
				convertData.ErrorToken = blockData.StartingToken; 
				return; 
			}
			
			
		} else  {
			nameToken = variables[i].NameToken[0]; 
			
		}
		
		
		
		if nameToken.Text == "" {
			convertData.ConvertResult = ConvertResult.Missing_Expected_Type; 
			convertData.ErrorDetail = "name Token is null in multiple declaration with value"; 
			convertData.ErrorToken = blockData.StartingToken; 
			return; 
		}
		
		if i != 0 {
			convertData.AppendChar(','); 
			convertData.AppendChar(' '); 
			
		}
		convertData.AppendString(nameToken.Text); 
		
	}
	
	convertData.AppendChar(' '); 
	PrintTokensDeclaration(convertData, blockData, nestCount, false, ""); 
	
}

func WriteVarTypeText(convertData *ConvertData, varTypeAsText string) {
	
	var firstChar byte  = varTypeAsText[0]; 
	if firstChar == '*' {
		var charArray []byte  = []byte(varTypeAsText)
		charArray[0] = '&'; 
		convertData.AppendString(string(charArray)); 
		return; 
	}
	convertData.AppendString(varTypeAsText); 
	
}

func PrintTokensDeclaration(convertData *ConvertData, blockData *BlockData, nestCount int, newLine bool, varTypeAsText string) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; var inMake bool  = false; var addedSpace bool  = false; 
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if token.Type == TokenType.NewLine {
			if lastType != TokenType.NewLine {
				convertData.NewLineWithTabs(); 
				
			}
			lastType = token.Type; 
			continue
			
		}
		if i == 1 {
			if token.Type == TokenType.LeftBrace {
				WriteVarTypeText(convertData, varTypeAsText); 
				
			}
			
		}
		if i == 2 || i == 1 {
			if lastType == TokenType.Make {
				if token.Type == TokenType.LeftParenthesis {
					convertData.AppendChar('('); 
					convertData.AppendString(varTypeAsText); 
					inMake = true; 
					continue
					
				}
				
			}
			
		}
		if inMake == true {
			if token.Type != TokenType.RightParenthesis {
				convertData.AppendChar(','); 
				convertData.AppendChar(' '); 
				
			}
			inMake = false; 
			
		}
		
		if token.Type == TokenType.RightBrace {
			convertData.DecrementNestCount()
			convertData.RemoveLastTab()
			
		}
		
		AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
		convertData.AppendString(token.Text); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		if token.Type == TokenType.LeftBrace {
			convertData.IncrementNestCount()
			
		}
		
	}
	
	
}

func ProcessInterfaceDeclaration(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Tokens) == 0 {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "No tokens in interface declaration"; 
		return; 
	}
	
	//first token is the name of the interface
	var interfaceNameToken Token  = blockData.Tokens[0]; 
	if interfaceNameToken.Type == TokenType.NA {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "interface name token is null in interface declaration"; 
		return; 
	}
	
	//block has method list and that's all, 0 count is blank interface
	convertData.AppendString("type " +interfaceNameToken.Text + " interface {\n\t"); 
	
	if blockData.Block == nil {
		convertData.ConvertResult = ConvertResult.Internal_Error; 
		convertData.ErrorDetail = "Block is null in interface declaration"; 
		return; 
	}
	
	var methodList []Function  = blockData.Block.MethodList; 
	if len(methodList) == 0 {
		EndInterface(convertData); 
		return; 
	}
	
	for i := 0; i < len(methodList); i++ {
	
		var function Function  = methodList[i]; 
		convertData.AppendString(function.Name + "(")
		
		PrintParameters(convertData, &function); 
		convertData.GeneratedCode = append(convertData.GeneratedCode, ')')
		
		convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
		
		if function.ReturnType != "" {
			convertData.AppendString(function.ReturnType)
			convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
			
			
		}
		convertData.GeneratedCode = append(convertData.GeneratedCode, '\n')
		
		convertData.GeneratedCode = append(convertData.GeneratedCode, '\t')
		
		
	}
	
	
	EndInterface(convertData); 
	
}

func EndInterface(convertData *ConvertData) {
	convertData.AppendChar('\r'); 
	convertData.AppendChar('}'); 
	convertData.AppendChar('\n'); 
	
}

func PrintParametersInterface(convertData *ConvertData, function Function) {
	
	var parameters []Variable  = function.Parameters; if len(parameters) == 0 {
		return; 
	}
	
	for parameterIndex := 0; parameterIndex < len(parameters); parameterIndex++ {
	
		var parameter Variable  = parameters[parameterIndex]; if parameter.NameToken[0].Type == TokenType.NA {
			convertData.ConvertResult = ConvertResult.Internal_Error; 
			convertData.ErrorDetail = "name token is null in PrintParameters"; 
			return; 
		}
		var typeAsText string  = JoinTextInListOfTokens( &parameter.TypeList); if typeAsText == "" {
			convertData.ConvertResult = ConvertResult.Internal_Error; 
			convertData.ErrorDetail = "var type text is null in PrintParameters"; 
			return; 
		}
		if parameterIndex != 0 {
			convertData.GeneratedCode = append(convertData.GeneratedCode, ',')
			
			convertData.GeneratedCode = append(convertData.GeneratedCode, ' ')
			
			
		}
		convertData.AppendString(parameter.NameToken[0].Text + " " +typeAsText)
		
	}
	
	
}

func AddMethodPrefix(convertData *ConvertData, varName string) {
	
	if convertData.MethodType == MethodType.None {
		return; 
	}
	
	if len(convertData.MethodVarNames) == 0 {
		return; 
	}
	if IsMethodVar(convertData, varName) == false {
		return; 
	}
	
	var firstLetter byte  = convertData.StructName[0]; convertData.AppendChar(firstLetter); 
	convertData.AppendChar('.'); 
	return; 
}

func IsMethodVar(convertData *ConvertData, varName string) bool {
	
	for i := 0; i < len(convertData.MethodVarNames); i++ {
	
		var name string  = convertData.MethodVarNames[i]; if name == varName {
			return true; 
		}
		
	}
	
	return false; 
}

func PrintTokens(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; var addedSpace bool  = false; 
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
		HandleToken(convertData, token); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		
	}
	
	
}

func HandleToken(convertData *ConvertData, token Token) {
	
	if token.Text == "\r" {
		return; 
	}
	if token.Text == "\r\n" {
		return; 
	}
	if token.Type == TokenType.NewLine {
		return; 
	}
	if token.Type == TokenType.Semicolon {
		var codeLength int  = len(convertData.GeneratedCode); var lastChar byte  = convertData.GeneratedCode[codeLength - 1]; if lastChar == ' ' {
			convertData.GeneratedCode[codeLength - 1] = ';'; 
			return; 
		}
		
	}
	convertData.AppendString(token.Text); 
	
}

func PrintTokensNewLine(convertData *ConvertData, blockData *BlockData, nestCount int, newLine bool) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; var addedSpace bool  = false; 
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if token.Type == TokenType.NewLine {
			convertData.NewLineWithTabs(); 
			lastType = token.Type; 
			continue
			
		}
		AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
		HandleToken(convertData, token); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		
	}
	
	
	if lastType != TokenType.NewLine {
		convertData.NewLineWithTabs(); 
		
	}
	if newLine == true {
		convertData.NewLineWithTabs(); 
		
	}
	
	
	
}

func HandleTokenPrintTokensNL(convertData *ConvertData, token Token) {
	
	if token.Type == TokenType.Semicolon {
		var codeLength int  = len(convertData.GeneratedCode); var lastChar byte  = convertData.GeneratedCode[codeLength - 1]; if lastChar == ' ' {
			convertData.GeneratedCode[codeLength - 1] = ';'; 
			return; 
		}
		
	}
	convertData.AppendString(token.Text); 
	
}

func PrintComment(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; var addedSpace bool  = false; 
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if i != 0 {
			AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
			
		}
		
		HandleToken(convertData, token); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		
	}
	
	
}

func HandleTokenComment(convertData *ConvertData, token Token) {
	
	if token.Text == "\r" {
		return; 
	}
	if token.Text == "\r\n" {
		return; 
	}
	if token.Type == TokenType.NewLine {
		return; 
	}
	if token.Type == TokenType.Semicolon {
		var codeLength int  = len(convertData.GeneratedCode)
		var lastChar byte  = convertData.GeneratedCode[codeLength - 1]; if lastChar == ' ' {
			convertData.GeneratedCode[codeLength - 1] = ';'; 
			return; 
		}
		
	}
	convertData.AppendString(token.Text); 
	
}

func PrintTokensNoNL(convertData *ConvertData, blockData *BlockData, nestCount int) {
	
	if len(blockData.Tokens) == 0 {
		convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData"); 
		return; 
	}
	var lastType IntTokenType  = TokenType.NA; var addedSpace bool  = false; 
	for i := 0; i < len(blockData.Tokens); i++ {
	
		var token Token  = blockData.Tokens[i]; 
		if token.Type == TokenType.NewLine {
			convertData.NewLineWithTabs(); 
			lastType = token.Type; 
			continue
			
		}
		AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace); 
		convertData.AppendString(token.Text); 
		addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i); 
		lastType = token.Type; 
		
	}
	
	
}

func AddSpaceBefore(convertData *ConvertData, thisType IntTokenType, lastType IntTokenType, tokenIndex int, addedSpace bool) {
	
	var codeLength int  = len(convertData.GeneratedCode)
	
	if codeLength == 0 {
		var lastCharAdded byte  = convertData.GeneratedCode[codeLength - 1]; 
		if lastCharAdded == ' ' {
			return; 
		}
		
	}
	
	if addedSpace == true {
		return; 
	}
	
	var addSpace bool  = false; var shouldHaveSpaceBeforePlus bool  = false
	
	switch thisType {
		
		
		case TokenType.NotEquals, TokenType.And, TokenType.AndAnd, TokenType.Or, TokenType.OrOr, TokenType.PlusEquals, TokenType.MinusEquals, 
		TokenType.MultiplyEquals, TokenType.DivideEquals, TokenType.GreaterThan, TokenType.LessThan, TokenType.EqualsEquals, TokenType.GreaterThanEquals, 
		TokenType.LessThanEquals, TokenType.Modulus, TokenType.ModulusEquals, TokenType.ColonEquals, TokenType.Equals, TokenType.LeftBrace, 
		TokenType.Comment:
			
			addSpace = true; 
			
			
		case TokenType.Plus:
			
			shouldHaveSpaceBeforePlus = lastType == TokenType.Identifier || lastType == TokenType.IntegerValue || lastType == TokenType.StringValue; 
			
			if shouldHaveSpaceBeforePlus {
				addSpace = true; 
				
			}
			
			
		case TokenType.Multiply:
			
			if lastType == TokenType.Identifier {
				addSpace = true; 
				
			}
			
			
		case TokenType.Divide:
			addSpace = true; 
			
			
		default:
			
			
	}
	
	
	if thisType == TokenType.Minus {
		var isOperator bool  = 
		lastType == TokenType.Identifier || 
		lastType == TokenType.IntegerValue || 
		IsOperator(lastType); 
		if isOperator {
			addSpace = true; 
			
		}
		
	}
	
	if thisType == TokenType.Identifier {
		if lastType == TokenType.Identifier {
			addSpace = true; 
			
		}
		
	}
	
	if addSpace {
		convertData.AppendChar(' '); 
		
	}
	
}

func AddSpaceAfter(convertData *ConvertData, thisType IntTokenType, lastType IntTokenType, tokenIndex int) bool {
	
	var addSpace bool  = false; 
	switch thisType {
		
		case TokenType.If, TokenType.Else, TokenType.For, TokenType.Switch, 
		TokenType.Struct, TokenType.Bool, TokenType.Int, TokenType.Int16, 
		TokenType.Int32, TokenType.Int64, TokenType.Int8, TokenType.AndAnd, 
		TokenType.Or, TokenType.Return, TokenType.PlusEquals, TokenType.MinusEquals, 
		TokenType.DivideEquals, TokenType.Enum, TokenType.Enumstruct, TokenType.GreaterThanEquals, 
		TokenType.LessThanEquals, TokenType.Goto, TokenType.Equals, TokenType.Divide, 
		TokenType.Fn, TokenType.Package, TokenType.Import, TokenType.Comma, 
		TokenType.Const, TokenType.Semicolon, TokenType.OrOr, TokenType.ColonEquals, 
		TokenType.NotEquals, TokenType.ModulusEquals, TokenType.Modulus, TokenType.EqualsEquals, 
		TokenType.LessThan, TokenType.MultiplyEquals, TokenType.GreaterThan, TokenType.Defer, 
		TokenType.ErrReturn, TokenType.Case, TokenType.Break, TokenType.Continue:
			
			addSpace = true; 
			
			
			
		default:
			
			
	}
	
	
	if thisType == TokenType.Multiply {
		addSpace = true; 
		
	}
	
	if thisType == TokenType.Minus {
		var isOperator bool  = 
		lastType == TokenType.Identifier || 
		lastType == TokenType.IntegerValue; 
		if isOperator {
			addSpace = true; 
			
		}
		
	}
	
	if thisType == TokenType.Plus || thisType == TokenType.And {
		var isOperator bool  = 
		lastType == TokenType.Identifier || 
		lastType == TokenType.IntegerValue; 
		if isOperator {
			addSpace = true; 
			
		}
		
	}
	
	if addSpace {
		convertData.AppendChar(' '); 
		return true; 
	}
	return false; 
}

func AddSpaceAfterBlock(convertData *ConvertData, nodeType IntNodeType, nestCount int) {
	
	switch nodeType {
		
		
		case NodeType.Invalid, 
		NodeType.Channel_Declaration, 
		NodeType.Channel_Declaration_With_Value, 
		NodeType.Interface_Declaration, 
		NodeType.Single_Declaration_No_Value, 
		NodeType.Multiple_Declarations_No_Value, 
		NodeType.Multiple_Declarations_With_Value, 
		NodeType.Multiple_Declarations_Same_Type_No_Value, 
		NodeType.Multiple_Declarations_Same_Type_With_Value, 
		NodeType.Multiple_Declarations_One_Type_One_Set_Value, 
		NodeType.Constant_Global_Variable, 
		NodeType.Constant_Global_Variable_With_Type, 
		NodeType.Struct_Variable_Declaration, 
		NodeType.Else_Statement, 
		NodeType.For_Loop, 
		NodeType.For_Loop_With_Declaration, 
		NodeType.Err_Return, 
		NodeType.Err_Check, 
		NodeType.Single_Import, 
		NodeType.Single_Import_With_Alias, 
		NodeType.NestedStruct, 
		NodeType.Struct_Declaration, 
		NodeType.Enum_Declaration, 
		NodeType.Enum_Variable, 
		NodeType.Enum_Variable_With_Value, 
		NodeType.Enum_Struct_Declaration, 
		NodeType.Comment, 
		NodeType.Append, 
		NodeType.Other, 
		NodeType.Switch:
			
			convertData.NewLineWithTabs(); 
			
			
		default:
			
			
	}
	
	
}
