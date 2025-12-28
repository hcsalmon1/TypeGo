
package core

import (
"fmt"
"strings"
"errors"
)

type CodeFormat struct {
GlobalBlock CodeBlock
Functions []Function

}


func (self *CodeFormat) Print() {
	
	fmt.Println("Printing Code Format:")
	
	var indent int  = 1
	
	self.GlobalBlock.Print(0)
	fmt.Println()
	
	var count int  = len(self.Functions)
	
	
	for i := 0; i < count; i++ {
	
		var function Function  = self.Functions[i]
		
		function.Print(indent)
		
	}
	
	
}

func (self *CodeFormat) PrintFunctions() {
	fmt.Println("Printing Functions:"); 
	
	var count int  = len(self.Functions)
	
	for i := 0; i < count; i++ {
	
		var function Function  = self.Functions[i]
		
		function.PrintData(); 
		
	}
	
	
}

func (self *CodeFormat) PrintNodeData() {
	
	fmt.Println("Printing Node Data:"); 
	
	var count int  = len(self.Functions)
	
	for i := 0; i < count; i++ {
	
		var function Function  = self.Functions[i]
		
		function.PrintData()
		
	}
	
	
}

func (self *CodeFormat) PrintNodeType() {
	
	fmt.Println("Printing Functions:"); 
	
	var count int  = len(self.Functions)
	
	for i := 0; i < count; i++ {
	
		var function Function  = self.Functions[i]
		
		function.PrintNodeTypes()
		
	}
	
	
}

type CodeBlock struct {
BlockDataList []BlockData
MethodList []Function

}


func (self *CodeBlock) Print(indent int) {
	
	var block_count int  = len(self.BlockDataList)
	
	if block_count != 0 {
		for i := 0; i < block_count; i++ {
		
			var block_data BlockData  = self.BlockDataList[i]
			
			if block_data.NodeType == NodeType.NewLine {
				continue 
				
			}
			fmt.Println()
			fmt.Print(addTabs(indent))
			
			block_data.Print(indent + 1)
			
		}
		
		
	}
	
	var method_count int  = len(self.MethodList)
	
	if method_count != 0 {
		for i := 0; i < method_count; i++ {
		
			fmt.Println()
			var function Function  = self.MethodList[i]
			
			function.Print(indent)
			
		}
		
		
	}
	
	fmt.Println()
	
}

func (self *CodeBlock) PrintData() {
	fmt.Println("\nPrinting blockdata:"); 
	var count int  = len(self.BlockDataList)
	
	if count == 0 {
		fmt.Println("\tempty\n"); 
		
	}
	for i := 0; i < count; i++ {
	
		self.BlockDataList[i].Print(0); 
		fmt.Println(); 
		
	}
	
	fmt.Println(); 
	
}

type BlockData struct {
NodeType IntNodeType
Tokens []Token
StartingToken Token
Block *CodeBlock
Variables []Variable
VarName string

}


func (self *BlockData) Print(indent int) {
	
	fmt.Printf("%s ", self.NodeType.ToString())
	var token_count int  = len(self.Tokens)
	
	if token_count == 0 {
		return 
	}
	for i := 0; i < len(self.Variables); i++ {
	
		self.Variables[i].Print()
		
	}
	
	for i := 0; i < token_count; i++ {
	
		fmt.Printf("%s ", self.Tokens[i].Text)
		
	}
	
	if self.Block != nil {
		self.Block.Print(indent + 1)
		
	}
	
}

func (self *BlockData) PrintNodeType(newLine bool) {
	fmt.Print("\tNode type: "); 
	if self.NodeType == NodeType.Invalid {
		fmt.Print("Invalid "); 
		self.AddNewLine(newLine); 
		return; 
	}
	fmt.Print("{NodeType} "); 
	self.AddNewLine(newLine); 
	
}

func (self *BlockData) PrintStartingToken(newLine bool) {
	
	if (newLine) {
		fmt.Print('\t'); 
		
	}
	if self.StartingToken.Type == TokenType.NA {
		fmt.Print("Starting token: "); 
		fmt.Print("null "); 
		self.AddNewLine(newLine); 
		return; 
	}
	fmt.Print("Starting token: "); 
	fmt.Print("'%s' ", self.StartingToken.Text); 
	self.AddNewLine(newLine); 
	
}

func (self *BlockData) AddNewLine(newLine bool) {
	if (newLine == true) {
		fmt.Println(); 
		
	}
	
}

func (self *BlockData) PrintTokens(newLine bool) {
	if (newLine) {
		fmt.Print('\t'); 
		
	}
	fmt.Print("tokens: "); 
	var count int  = len(self.Tokens)
	
	if count == 0 {
		fmt.Print("empty "); 
		self.AddNewLine(newLine); 
		return; 
	}
	for i := 0; i < count; i++ {
	
		fmt.Print("'{Tokens[i].Text}' "); 
		
	}
	
	self.AddNewLine(newLine); 
	
}

func (self *BlockData) PrintWithCodeBlock() {
	var NO_NEW_LINE bool  = false; 
	
	fmt.Print("\tNode type: "); 
	fmt.Printf("%s ", self.NodeType.ToString()); 
	self.PrintStartingToken(NO_NEW_LINE); 
	self.PrintTokens(NO_NEW_LINE); 
	
	if (self.Block == nil) {
		fmt.Print("Block: "); 
		fmt.Print("null "); 
		
	} else  {
	fmt.Print("Block: "); 
	fmt.Print("not null "); 
	
	}
	
	
}

func (self *BlockData) PrintVariables(newLine bool) {
	if newLine {
		fmt.Print('\t'); 
		
	}
	fmt.Print("Vars: "); 
	var count int  = len(self.Variables)
	
	if count == 0 {
		fmt.Print("empty "); 
		self.AddNewLine(newLine); 
		return; 
	}
	for i := 0; i < count; i++ {
	
		self.Variables[i].Print(); 
		
	}
	
	self.AddNewLine(newLine); 
	
}

func (self *BlockData) Validate(formatData *FormatData) bool {
	
	var validator NodeValidator
	var err error  = nil
	
	
	switch self.NodeType {
	case NodeType.Invalid:
	formatData.Result = FormatResult.Invalid_Node
	return false
	case NodeType.Channel_Declaration:
	case NodeType.Channel_Declaration_With_Value:
	case NodeType.Interface_Declaration:
	case NodeType.Single_Declaration_With_Value:
	err = validator.ValidateSingleDeclarationWithValue(self)
	if err != nil {
	formatData.SetError(FormatResult.Invalid_Node, err.Error())
	return false
	
	}
	return true
	case NodeType.Single_Declaration_No_Value:
	case NodeType.Multiple_Declarations_No_Value:
	case NodeType.Multiple_Declarations_With_Value:
	case NodeType.Multiple_Declarations_Same_Type_No_Value:
	case NodeType.Multiple_Declarations_Same_Type_With_Value:
	case NodeType.Multiple_Declarations_One_Type_One_Set_Value:
	case NodeType.Constant_Global_Variable:
	case NodeType.Constant_Global_Variable_With_Type:
	case NodeType.Struct_Variable_Declaration:
	case NodeType.If_Statement_With_Declaration:
	case NodeType.If_Statement:
	case NodeType.Else_Statement:
	case NodeType.For_Loop:
	case NodeType.For_Loop_With_Declaration:
	case NodeType.Err_Return:
	case NodeType.Err_Check:
	case NodeType.Multi_Line_Import:
	case NodeType.Single_Import:
	case NodeType.Single_Import_With_Alias:
	case NodeType.NestedStruct:
	case NodeType.Struct_Declaration:
	case NodeType.Enum_Declaration:
	case NodeType.Enum_Variable:
	case NodeType.Enum_Variable_With_Value:
	case NodeType.Enum_Struct_Declaration:
	case NodeType.Enum_Struct_Declaration_With_Alias:
	case NodeType.NewLine:
	case NodeType.Comment:
	case NodeType.Append:
	case NodeType.Package:
	case NodeType.Other:
	case NodeType.Switch:
	case NodeType.Return:
	case NodeType.Break:
	default:
	return true
	
	}
	
	return true
}

type Variable struct {
TypeList []Token
NameToken []Token

}


func (self *Variable) SetToDefaults() {
	
	if (self.TypeList == nil) {
		self.TypeList = make([]Token, 0)
		
	} else  {
	self.TypeList = nil
	
	}
	
	
	if (self.NameToken == nil) {
		self.NameToken = make([]Token, 0); 
		
	} else  {
	self.NameToken = nil
	
	}
	
	
}

func (self *Variable) PrintTypeList() {
	
	if self.TypeList == nil {
		fmt.Print("Typelist null "); 
		return; 
	}
	var count int  = len(self.TypeList)
	
	if count == 0 {
		fmt.Print("Typelist empty "); 
		return; 
	}
	for i := 0; i < count; i++ {
	
		fmt.Print(" '{TypeList[i].Text}', "); 
		
	}
	
	
}

func (self *Variable) PrintVarName() {
	
	if self.NameToken == nil {
		fmt.Println(" null"); 
		return; 
	}
	var count int  = len(self.NameToken)
	
	if count == 0 {
		fmt.Println(" zero"); 
		return; 
	}
	fmt.Print("varname: "); 
	fmt.Print("{NameToken[0].Text} "); 
	
}

func (self *Variable) Print() {
	var type_list_count int  = len(self.TypeList)
	
	var name_list_count int  = len(self.NameToken)
	
	if type_list_count == 0 || name_list_count == 0 {
		fmt.Printf("Invalid, type count: %d name count: %d", type_list_count, name_list_count)
		return 
	}
	for i := 0; i < type_list_count; i++ {
	
		fmt.Printf("%s", self.TypeList[i].Text)
		
	}
	
	fmt.Printf(" %s\n", self.NameToken[0].Text)
	
}

func (self *Variable) Println() {
	self.PrintTypeList(); 
	self.PrintVarName(); 
	fmt.Println(); 
	
}

func (self *Variable) ConvertToString() string {
	var sb []byte  = make([]byte, 0)
	
	var count int  = len(self.TypeList)
	
	if count == 0 {
		var input string  = "INVALID_TYPE "
		
		var length int  = len(input)
		
		for i := 0; i < length; i++ {
		
			sb = append(sb, input[i])
			
			
		}
		
		
	}
	for i := 0; i < count; i++ {
	
		var token Token  = self.TypeList[i]
		
		AppendStringToSlice( &sb, token.Text)
		
	}
	
	sb = append(sb, ' ')
	
	if self.NameToken == nil {
		var input string  = "null"
		
		var length int  = len(input)
		
		for i := 0; i < length; i++ {
		
			sb = append(sb, input[i])
			
			
		}
		
		
	} else  {
	var input string  = self.NameToken[0].Text
	
	var length int  = len(input)
	
	for i := 0; i < length; i++ {
	
		sb = append(sb, input[i])
		
		
	}
	
	
	}
	
	
	return string(sb); 
}

func (self *Variable) MoveTypeToName() {
	
	if len(self.TypeList) == 0 {
		return; 
	}
	self.NameToken = append(self.NameToken, self.TypeList[0])
	
	self.TypeList = nil
	
}

type Function struct {
ReturnType string
Name string
Parameters []Variable
InnerBlock *CodeBlock
StartingToken Token

}


func (self *Function) Print(indent int) {
	fmt.Print(addTabs(indent))
	fmt.Print("fn ")
	if len(self.ReturnType) != 0 {
		fmt.Printf("%s ", self.ReturnType)
		
	}
	
	fmt.Printf("%s%s%s(", CREAM_TEXT, self.Name, RESET_TEXT)
	var count int  = len(self.Parameters)
	
	for i := 0; i < count; i++ {
	
		var parameter Variable  = self.Parameters[i]
		
		parameter.Print()
		fmt.Print(", ")
		
	}
	
	
	fmt.Printf("%c %c", ')', '{')
	self.InnerBlock.Print(indent + 1)
	fmt.Print(addTabs(indent))
	fmt.Printf("%c\n", '}')
	
}

func (self *Function) PrintData() {
	fmt.Printf("\t%s", self.ReturnType); 
	fmt.Printf(" %s(", self.Name); 
	var count int  = len(self.Parameters)
	
	for i := 0; i < count; i++ {
	
		var parameter Variable  = self.Parameters[i]
		
		parameter.Print(); 
		fmt.Print(", "); 
		
	}
	
	fmt.Println(')'); 
	self.InnerBlock.PrintData(); 
	
}

func (self *Function) PrintNodeTypes() {
	
	fmt.Println("Printing Node Types:"); 
	if (self.InnerBlock == nil) {
		fmt.Println("\tinner block is null!:"); 
		return; 
	}
	var block *CodeBlock  = self.InnerBlock; 
	
	var count int  = len(block.BlockDataList)
	
	if (count == 0) {
		fmt.Println("\tblock data list is 0!"); 
		return; 
	}
	
	for i := 0; i < count; i++ {
	
		var blockData BlockData  = block.BlockDataList[i]; 
		fmt.Printf("\t%s\n", blockData.NodeType.ToString()); 
		
	}
	
	
}

type Token struct {
Text string
Type IntTokenType
LineNumber int
CharNumber int

}


func (self *Token) IsType(token_type IntTokenType) bool {
	return self.Type == token_type; 
}

func (self *Token) PrintText() {
	fmt.Printf("\tToken: text %s\n", self.Text)
	
}

type ParseData struct {
TokenList []Token
LastToken Token
CharacterIndex int
Code string
LineCount int
CharCount int
ParseResult IntParseResult
CodeLength int

}


func (self *ParseData) IsError() bool {
	return self.ParseResult != ParseResult.Ok
}

type FormatData struct {
TokenList []Token
Result IntFormatResult
ErrorDetail string
ErrorFunction string
ErrorProcess string
FunctionLog []string
ProcessLog []string
LogIndent int
ErrorToken Token
TokenIndex int
TempToken Token

}


func (self *FormatData) IndexInBounds() bool {
	if self.TokenIndex >= len(self.TokenList) {
		return false
	}
	return true
}

func (self *FormatData) GetToken() Token {
	
	if self.TokenIndex >= len(self.TokenList) {
		return EmptyToken()
	}
	var token Token  = self.TokenList[self.TokenIndex]
	
	self.ErrorToken = token
	return token
}

func (self *FormatData) GetTokenByIndex(index int) Token {
	
	if (index >= len(self.TokenList)) {
		return EmptyToken(); 
	}
	return self.TokenList[index]; 
}

func (self *FormatData) Increment() {
	self.TokenIndex += 1; 
	
}

func (self *FormatData) IncrementTwice() {
	self.TokenIndex += 2; 
	
}

func (self *FormatData) GetNextToken() Token {
	
	self.Increment(); 
	return self.GetToken(); 
}

func (self *FormatData) IncrementIfSame(indexBefore int) {
	if (indexBefore == self.TokenIndex) {
		self.Increment(); 
		
	}
	
}

func (self *FormatData) EndOfFileError(errorToken Token) {
	if (self.IsError()) {
		return; 
	}
	self.Result = FormatResult.EndOfFile; 
	self.ErrorToken = errorToken; 
	
}

func (self *FormatData) IsError() bool {
	return self.Result != FormatResult.Ok; 
}

func (self *FormatData) AddFnTrace(functionName string) {
	self.ErrorFunction = functionName; 
	
}

func (self *FormatData) MissingExpectedTypeError(errorToken Token, detail string) {
	if (self.IsError()) {
		return; 
	}
	self.Result = FormatResult.MissingExpectedType; 
	self.ErrorDetail = detail; 
	self.ErrorToken = errorToken; 
	
}

func (self *FormatData) UnexpectedTypeError(errorToken Token, detail string) {
	
	if (self.IsError()) {
		return; 
	}
	self.Result = FormatResult.UnexpectedType; 
	self.ErrorDetail = detail; 
	self.ErrorToken = errorToken; 
	
}

func (self *FormatData) ExpectType(token_type IntTokenType, detail string) bool {
	var token Token  = self.GetToken(); 
	if (token.Type == TokenType.NA) {
		self.EndOfFileError(token); 
		return false; 
	}
	if (token.Type == token_type) {
		return true; 
	}
	self.MissingExpectedTypeError(token, detail); 
	return false; 
}

func (self *FormatData) ExpectNextType(token_type IntTokenType, detail string) bool {
	
	var token Token  = self.GetNextToken(); 
	if (token.Type == TokenType.NA) {
		self.EndOfFileError(token); 
		return false; 
	}
	if (token.Type == token_type) {
		return true; 
	}
	self.MissingExpectedTypeError(token, detail); 
	return false; 
}

func (self *FormatData) IncrementIfNewLine() {
	var token Token  = self.GetToken(); 
	if (token.Type == TokenType.NA) {
		return; 
	}
	if (token.Type == TokenType.NewLine) {
		self.Increment(); 
		
	}
	
}

func (self *FormatData) UnsupportedFeatureError(errorToken Token, detail string) {
	if (self.IsError()) {
		return; 
	}
	self.Result = FormatResult.UnsupportedFeature; 
	self.ErrorToken = errorToken; 
	self.ErrorDetail = detail; 
	
}

func (self *FormatData) IsValidToken(token Token) bool {
	return token.Type != TokenType.NA
}

func (self *FormatData) SetErrorFunction(function string) {
	self.ErrorFunction = function
	
}

func (self *FormatData) SetErrorProcess(process string) {
	self.ErrorProcess = process
	
}

func (self *FormatData) IncreaseLogIndent() {
	self.LogIndent += 1
	
}

func (self *FormatData) DecreaseLogIndent() {
	if self.LogIndent == 0 {
		return 
	}
	self.LogIndent -= 1
	
}

func (self *FormatData) AddToProcessLog(info string) {
	var info_with_spaces string  = addSpaces(info, self.LogIndent)
	
	self.ProcessLog = append(self.ProcessLog, info_with_spaces)
	
}

func (self *FormatData) PrintProcessLog() {
	fmt.Println("\nPrinting process log:")
	var count int  = len(self.ProcessLog)
	
	if count == 0 {
		fmt.Println("\tEmpty\n")
		return 
	}
	for i := 0; i < count; i++ {
	
		var log string  = self.ProcessLog[i]
		
		fmt.Printf("\t%s\n", log)
		
	}
	
	fmt.Println()
	
}

func (self *FormatData) SetError(result IntFormatResult, detail string) {
	self.Result = result; 
	self.ErrorDetail = detail; 
	
}

func (self *FormatData) AddToFunctionLog(function string) {
	self.FunctionLog = append(self.FunctionLog, function)
	
}

func (self *FormatData) PrintFunctionLog() {
	fmt.Println("Printing function log:")
	var count int  = len(self.FunctionLog)
	
	if count == 0 {
		fmt.Println("\tEmpty")
		
	}
	for i := 0; i < count; i++ {
	
		fmt.Print("\t")
		fmt.Println(self.FunctionLog[i])
		
	}
	
	fmt.Println()
	
}

type ConvertData struct {
GeneratedCode []byte
ConvertResult IntConvertResult
ErrorDetail string
CodeFormat CodeFormat
ErrorProcess string
ErrorFunction string
ProcessLog []string
LogIndent int
ErrorToken Token
LastNodeType IntNodeType
MethodVarNames []string
MethodType IntMethodType
StructName string
NestCount int

}


func (self *ConvertData) IsInfiniteWhileLoop(count *int, max int) bool {
	_ = self
	* count += 1; 
	return * count >= max
}

func (self *ConvertData) IsError() bool {
	if self.ConvertResult != ConvertResult.Ok {
		return true; 
	}
	return false; 
}

func (self *ConvertData) SetError(result IntConvertResult, detail string, errorToken Token) {
	if self.ConvertResult != ConvertResult.Ok {
		return; 
	}
	self.ConvertResult = result; 
	self.ErrorToken = errorToken; 
	self.ErrorDetail = detail; 
	
}

func (self *ConvertData) EndOfFileError(lastToken Token) {
	self.ConvertResult = ConvertResult.Unexpected_End_Of_File; 
	self.ErrorToken = lastToken; 
	
}

func (self *ConvertData) MissingTypeError(lastToken Token, detail string) {
	self.ConvertResult = ConvertResult.Missing_Expected_Type; 
	self.ErrorDetail = detail; 
	self.ErrorToken = lastToken; 
	
}

func (self *ConvertData) NewLine() {
	_ = self
	self.AppendChar('\n'); 
	
}

func (self *ConvertData) NewLineWithTabs() {
	self.AppendChar('\n'); 
	if self.NestCount == 0 {
		return; 
	}
	for i := 0; i < self.NestCount; i++ {
	
		self.AppendChar('\t'); 
		
	}
	
	
}

func (self *ConvertData) AddTabs() {
	for i := 0; i < self.NestCount; i++ {
	
		self.AppendChar('\t'); 
		
	}
	
	
}

func (self *ConvertData) IncrementNestCount() {
	self.NestCount += 1; 
	
}

func (self *ConvertData) DecrementNestCount() {
	self.NestCount -= 1; 
	if self.NestCount < 0 {
		self.NestCount = 0; 
		
	}
	
}

func (self *ConvertData) UnexpectedTypeError(token Token, detail string) {
	self.ConvertResult = ConvertResult.Unexpected_Type; 
	self.ErrorDetail = detail; 
	self.ErrorToken = token; 
	
}

func (self *ConvertData) AppendToken(token Token) {
	self.AppendString(token.Text)
	
}

func (self *ConvertData) AppendString(input string) {
	for i := 0; i < len(input); i++ {
	
		self.GeneratedCode = append(self.GeneratedCode, input[i])
		
	}
	
	
}

func (self *ConvertData) AppendChar(input byte) {
	self.GeneratedCode = append(self.GeneratedCode, input)
	
	
}

func (self *ConvertData) NoTokenError(errorToken Token, detail string) {
	self.ConvertResult = ConvertResult.No_Token_In_Node; 
	self.ErrorToken = errorToken; 
	self.ErrorDetail = detail; 
	
}

func (self *ConvertData) SetErrorFunction(function string) {
	self.ErrorFunction = function
	
}

type NodeValidator struct {

}


func (self *NodeValidator) ValidateSingleDeclarationWithValue(block_data *BlockData) error {
	//IntNodeType NodeType
	//[]Token Tokens
	//Token StartingToken
	//*CodeBlock Block
	//[]Variable Variables
	//string VarName
	
	//We expect: 
	//-1 Variable, with a valid type and a name
	//-Non empty token list
	
	if block_data.NodeType != NodeType.Single_Declaration_With_Value {
		return errors.New("NodeType Mismatch, single declaration with value")
	}
	
	var variables []Variable  = block_data.Variables
	
	if variables == nil {
		return errors.New("variables are nil, single declaration with value")
	}
	var variable_count int  = len(variables)
	
	if variable_count != 1 {
		return errors.New("variable count != 1, single declaration with value")
	}
	
	var variable Variable  = variables[0]
	
	var type_list []Token  = variable.TypeList
	
	var var_names []Token  = variable.NameToken
	
	
	if type_list == nil {
		return errors.New("Type list is nil, single declaration with value")
	}
	if len(type_list) == 0 {
		return errors.New("Type list is empty, single declaration with value")
	}
	if var_names == nil {
		return errors.New("Var names is nil, single declaration with value")
	}
	var var_name_count int  = len(var_names)
	
	if var_name_count == 0 {
		return errors.New("No var name, single declaration with value")
	}
	var tokens []Token  = block_data.Tokens
	
	if tokens == nil {
		return errors.New("tokens is nil, single declaration with value")
	}
	var token_count int  = len(tokens)
	
	if token_count == 0 {
		return errors.New("zero tokens, value expected, single declaration with value")
	}
	
	return nil
}

func (self *NodeValidator) ValidateSingleDeclarationNoValue(block_data *BlockData) error {
	//IntNodeType NodeType
	//[]Token Tokens - 0
	//Token StartingToken - 
	//*CodeBlock Block - nil
	//[]Variable Variables - 1
	//string VarName
	
	//We expect: 
	//-1 Variable, with a valid type and a name
	//-empty token list
	
	if block_data.NodeType != NodeType.Single_Declaration_No_Value {
		return errors.New("NodeType Mismatch, single declaration no value")
	}
	
	var variables []Variable  = block_data.Variables
	
	if variables == nil {
		return errors.New("variables are nil, single declaration no value")
	}
	var variable_count int  = len(variables)
	
	if variable_count != 1 {
		return errors.New("variable count != 1, single declaration no value")
	}
	
	var variable Variable  = variables[0]
	
	
	//TYPE
	var type_list []Token  = variable.TypeList
	
	
	if type_list == nil {
		return errors.New("Type list is nil, single declaration no value")
	}
	if len(type_list) == 0 {
		return errors.New("Type list is empty, single declaration no value")
	}
	
	//NAME
	var var_names []Token  = variable.NameToken
	
	if var_names == nil {
		return errors.New("Var names is nil, single declaration no value")
	}
	var var_name_count int  = len(var_names)
	
	if var_name_count == 0 {
		return errors.New("No var name, single declaration no value")
	}
	if var_name_count != 1 {
		return errors.New("var name count should be 1, single declaration no value")
	}
	
	return nil
}

func (self *NodeValidator) ValidateIfStatement(block_data *BlockData) error {
	//IntNodeType NodeType
	//[]Token Tokens - if condition
	//Token StartingToken - if
	//*CodeBlock Block
	//[]Variable Variables - nil
	//string VarName
	
	if block_data.NodeType != NodeType.If_Statement {
		return errors.New("NodeType Mismatch, if statement")
	}
	var tokens []Token  = block_data.Tokens
	
	
	if tokens != nil {
		return errors.New("Tokens is nil, no condition, if statement")
	}
	if len(tokens) == 0 {
		return errors.New("Tokens is zero, no condition, if statement")
	}
	if block_data.Block == nil {
		return errors.New("Block is null, if statement")
	}
	
	return nil
}

type BracketCounts struct {
OpenSqCount int
OpenBraceCount int
OpenBracketCount int

}


func (self *BracketCounts) AreAllZero() bool {
	return self.OpenSqCount == 0 && self.OpenBraceCount == 0 && self.OpenBracketCount == 0
}


func EmptyBlockData() BlockData {
	return BlockData {NodeType:NodeType.Invalid, }
}

func EmptyToken() Token {
	return Token {Text:"", Type:TokenType.NA, LineNumber:0, CharNumber:0, }
}

func addSpaces(s string, count int) string {
	return strings.Repeat(" ", count)+s
}

func addTabs(count int) string {
	return strings.Repeat("\t", count)
}
