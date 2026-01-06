
package core

type IntTokenType int
var TokenType = struct {
	Fn IntTokenType
IntegerValue IntTokenType
DecimalValue IntTokenType
CharValue IntTokenType
StringValue IntTokenType
True IntTokenType
False IntTokenType
Switch IntTokenType
Channel_Setter IntTokenType
ErrReturn IntTokenType
ErrCheck IntTokenType
Case IntTokenType
Default IntTokenType
Int IntTokenType
Int8 IntTokenType
Int16 IntTokenType
Int32 IntTokenType
Int64 IntTokenType
Uint IntTokenType
Uint8 IntTokenType
Uint16 IntTokenType
Uint32 IntTokenType
Uint64 IntTokenType
Float32 IntTokenType
Float64 IntTokenType
String IntTokenType
Byte IntTokenType
Rune IntTokenType
Bool IntTokenType
Void IntTokenType
Struct IntTokenType
Enum IntTokenType
Enumstruct IntTokenType
Error IntTokenType
Map IntTokenType
Chan IntTokenType
Var IntTokenType
Type IntTokenType
Interface IntTokenType
Goto IntTokenType
If IntTokenType
Else IntTokenType
For IntTokenType
While IntTokenType
Return IntTokenType
Break IntTokenType
Continue IntTokenType
Defer IntTokenType
Go IntTokenType
Not IntTokenType
Minus IntTokenType
Plus IntTokenType
PlusPlus IntTokenType
Divide IntTokenType
Multiply IntTokenType
Equals IntTokenType
And IntTokenType
AndAnd IntTokenType
Or IntTokenType
OrOr IntTokenType
PlusEquals IntTokenType
MinusEquals IntTokenType
MultiplyEquals IntTokenType
DivideEquals IntTokenType
GreaterThan IntTokenType
LessThan IntTokenType
EqualsEquals IntTokenType
GreaterThanEquals IntTokenType
LessThanEquals IntTokenType
Modulus IntTokenType
ModulusEquals IntTokenType
NotEquals IntTokenType
Colon IntTokenType
ColonEquals IntTokenType
MinusMinus IntTokenType
Identifier IntTokenType
Print IntTokenType
Println IntTokenType
LeftParenthesis IntTokenType
RightParenthesis IntTokenType
LeftBrace IntTokenType
RightBrace IntTokenType
LeftSquareBracket IntTokenType
RightSquareBracket IntTokenType
Make IntTokenType
Append IntTokenType
Semicolon IntTokenType
Package IntTokenType
Import IntTokenType
Comma IntTokenType
FullStop IntTokenType
In IntTokenType
Const IntTokenType
Tab IntTokenType
MultiLineStart IntTokenType
MultiLineEnd IntTokenType
Comment IntTokenType
EndComment IntTokenType
NewLine IntTokenType
NA IntTokenType
}{
Fn: 0,
IntegerValue: 1,
DecimalValue: 2,
CharValue: 3,
StringValue: 4,
True: 5,
False: 6,
Switch: 7,
Channel_Setter: 8,
ErrReturn: 9,
ErrCheck: 10,
Case: 11,
Default: 12,
Int: 13,
Int8: 14,
Int16: 15,
Int32: 16,
Int64: 17,
Uint: 18,
Uint8: 19,
Uint16: 20,
Uint32: 21,
Uint64: 22,
Float32: 23,
Float64: 24,
String: 25,
Byte: 26,
Rune: 27,
Bool: 28,
Void: 29,
Struct: 30,
Enum: 31,
Enumstruct: 32,
Error: 33,
Map: 34,
Chan: 35,
Var: 36,
Type: 37,
Interface: 38,
Goto: 39,
If: 40,
Else: 41,
For: 42,
While: 43,
Return: 44,
Break: 45,
Continue: 46,
Defer: 47,
Go: 48,
Not: 49,
Minus: 50,
Plus: 51,
PlusPlus: 52,
Divide: 53,
Multiply: 54,
Equals: 55,
And: 56,
AndAnd: 57,
Or: 58,
OrOr: 59,
PlusEquals: 60,
MinusEquals: 61,
MultiplyEquals: 62,
DivideEquals: 63,
GreaterThan: 64,
LessThan: 65,
EqualsEquals: 66,
GreaterThanEquals: 67,
LessThanEquals: 68,
Modulus: 69,
ModulusEquals: 70,
NotEquals: 71,
Colon: 72,
ColonEquals: 73,
MinusMinus: 74,
Identifier: 75,
Print: 76,
Println: 77,
LeftParenthesis: 78,
RightParenthesis: 79,
LeftBrace: 80,
RightBrace: 81,
LeftSquareBracket: 82,
RightSquareBracket: 83,
Make: 84,
Append: 85,
Semicolon: 86,
Package: 87,
Import: 88,
Comma: 89,
FullStop: 90,
In: 91,
Const: 92,
Tab: 93,
MultiLineStart: 94,
MultiLineEnd: 95,
Comment: 96,
EndComment: 97,
NewLine: 98,
NA: 99,
}

func (self IntTokenType) ToString() string {
	switch self {
	case TokenType.Fn:
		return "Fn"
	case TokenType.IntegerValue:
		return "IntegerValue"
	case TokenType.DecimalValue:
		return "DecimalValue"
	case TokenType.CharValue:
		return "CharValue"
	case TokenType.StringValue:
		return "StringValue"
	case TokenType.True:
		return "True"
	case TokenType.False:
		return "False"
	case TokenType.Switch:
		return "Switch"
	case TokenType.Channel_Setter:
		return "Channel_Setter"
	case TokenType.ErrReturn:
		return "ErrReturn"
	case TokenType.ErrCheck:
		return "ErrCheck"
	case TokenType.Case:
		return "Case"
	case TokenType.Default:
		return "Default"
	case TokenType.Int:
		return "Int"
	case TokenType.Int8:
		return "Int8"
	case TokenType.Int16:
		return "Int16"
	case TokenType.Int32:
		return "Int32"
	case TokenType.Int64:
		return "Int64"
	case TokenType.Uint:
		return "Uint"
	case TokenType.Uint8:
		return "Uint8"
	case TokenType.Uint16:
		return "Uint16"
	case TokenType.Uint32:
		return "Uint32"
	case TokenType.Uint64:
		return "Uint64"
	case TokenType.Float32:
		return "Float32"
	case TokenType.Float64:
		return "Float64"
	case TokenType.String:
		return "String"
	case TokenType.Byte:
		return "Byte"
	case TokenType.Rune:
		return "Rune"
	case TokenType.Bool:
		return "Bool"
	case TokenType.Void:
		return "Void"
	case TokenType.Struct:
		return "Struct"
	case TokenType.Enum:
		return "Enum"
	case TokenType.Enumstruct:
		return "Enumstruct"
	case TokenType.Error:
		return "Error"
	case TokenType.Map:
		return "Map"
	case TokenType.Chan:
		return "Chan"
	case TokenType.Var:
		return "Var"
	case TokenType.Type:
		return "Type"
	case TokenType.Interface:
		return "Interface"
	case TokenType.Goto:
		return "Goto"
	case TokenType.If:
		return "If"
	case TokenType.Else:
		return "Else"
	case TokenType.For:
		return "For"
	case TokenType.While:
		return "While"
	case TokenType.Return:
		return "Return"
	case TokenType.Break:
		return "Break"
	case TokenType.Continue:
		return "Continue"
	case TokenType.Defer:
		return "Defer"
	case TokenType.Go:
		return "Go"
	case TokenType.Not:
		return "Not"
	case TokenType.Minus:
		return "Minus"
	case TokenType.Plus:
		return "Plus"
	case TokenType.PlusPlus:
		return "PlusPlus"
	case TokenType.Divide:
		return "Divide"
	case TokenType.Multiply:
		return "Multiply"
	case TokenType.Equals:
		return "Equals"
	case TokenType.And:
		return "And"
	case TokenType.AndAnd:
		return "AndAnd"
	case TokenType.Or:
		return "Or"
	case TokenType.OrOr:
		return "OrOr"
	case TokenType.PlusEquals:
		return "PlusEquals"
	case TokenType.MinusEquals:
		return "MinusEquals"
	case TokenType.MultiplyEquals:
		return "MultiplyEquals"
	case TokenType.DivideEquals:
		return "DivideEquals"
	case TokenType.GreaterThan:
		return "GreaterThan"
	case TokenType.LessThan:
		return "LessThan"
	case TokenType.EqualsEquals:
		return "EqualsEquals"
	case TokenType.GreaterThanEquals:
		return "GreaterThanEquals"
	case TokenType.LessThanEquals:
		return "LessThanEquals"
	case TokenType.Modulus:
		return "Modulus"
	case TokenType.ModulusEquals:
		return "ModulusEquals"
	case TokenType.NotEquals:
		return "NotEquals"
	case TokenType.Colon:
		return "Colon"
	case TokenType.ColonEquals:
		return "ColonEquals"
	case TokenType.MinusMinus:
		return "MinusMinus"
	case TokenType.Identifier:
		return "Identifier"
	case TokenType.Print:
		return "Print"
	case TokenType.Println:
		return "Println"
	case TokenType.LeftParenthesis:
		return "LeftParenthesis"
	case TokenType.RightParenthesis:
		return "RightParenthesis"
	case TokenType.LeftBrace:
		return "LeftBrace"
	case TokenType.RightBrace:
		return "RightBrace"
	case TokenType.LeftSquareBracket:
		return "LeftSquareBracket"
	case TokenType.RightSquareBracket:
		return "RightSquareBracket"
	case TokenType.Make:
		return "Make"
	case TokenType.Append:
		return "Append"
	case TokenType.Semicolon:
		return "Semicolon"
	case TokenType.Package:
		return "Package"
	case TokenType.Import:
		return "Import"
	case TokenType.Comma:
		return "Comma"
	case TokenType.FullStop:
		return "FullStop"
	case TokenType.In:
		return "In"
	case TokenType.Const:
		return "Const"
	case TokenType.Tab:
		return "Tab"
	case TokenType.MultiLineStart:
		return "MultiLineStart"
	case TokenType.MultiLineEnd:
		return "MultiLineEnd"
	case TokenType.Comment:
		return "Comment"
	case TokenType.EndComment:
		return "EndComment"
	case TokenType.NewLine:
		return "NewLine"
	case TokenType.NA:
		return "NA"
	default:
		return "Unknown"
}
}

type IntLastTokenType int
var LastTokenType = struct {
	Null IntLastTokenType
Identifier IntLastTokenType
CustomVarType IntLastTokenType
Vartype IntLastTokenType
LeftSqBracket IntLastTokenType
RightSqBracket IntLastTokenType
LeftParenth IntLastTokenType
RightParenth IntLastTokenType
Comma IntLastTokenType
Semicolon IntLastTokenType
Newline IntLastTokenType
IntegerValue IntLastTokenType
Map IntLastTokenType
Pointer IntLastTokenType
FullStop IntLastTokenType
Interface IntLastTokenType
LeftBrace IntLastTokenType
RightBrace IntLastTokenType
}{
Null: 0,
Identifier: 1,
CustomVarType: 2,
Vartype: 3,
LeftSqBracket: 4,
RightSqBracket: 5,
LeftParenth: 6,
RightParenth: 7,
Comma: 8,
Semicolon: 9,
Newline: 10,
IntegerValue: 11,
Map: 12,
Pointer: 13,
FullStop: 14,
Interface: 15,
LeftBrace: 16,
RightBrace: 17,
}

func (self IntLastTokenType) ToString() string {
	switch self {
	case LastTokenType.Null:
		return "Null"
	case LastTokenType.Identifier:
		return "Identifier"
	case LastTokenType.CustomVarType:
		return "CustomVarType"
	case LastTokenType.Vartype:
		return "Vartype"
	case LastTokenType.LeftSqBracket:
		return "LeftSqBracket"
	case LastTokenType.RightSqBracket:
		return "RightSqBracket"
	case LastTokenType.LeftParenth:
		return "LeftParenth"
	case LastTokenType.RightParenth:
		return "RightParenth"
	case LastTokenType.Comma:
		return "Comma"
	case LastTokenType.Semicolon:
		return "Semicolon"
	case LastTokenType.Newline:
		return "Newline"
	case LastTokenType.IntegerValue:
		return "IntegerValue"
	case LastTokenType.Map:
		return "Map"
	case LastTokenType.Pointer:
		return "Pointer"
	case LastTokenType.FullStop:
		return "FullStop"
	case LastTokenType.Interface:
		return "Interface"
	case LastTokenType.LeftBrace:
		return "LeftBrace"
	case LastTokenType.RightBrace:
		return "RightBrace"
	default:
		return "Unknown"
}
}

type IntNodeType int
var NodeType = struct {
	Invalid IntNodeType
Channel_Declaration IntNodeType
Channel_Declaration_With_Value IntNodeType
Interface_Declaration IntNodeType
Single_Declaration_With_Value IntNodeType
Single_Declaration_No_Value IntNodeType
Multiple_Declarations_No_Value IntNodeType
Multiple_Declarations_With_Value IntNodeType
Multiple_Declarations_Same_Type_No_Value IntNodeType
Multiple_Declarations_Same_Type_With_Value IntNodeType
Multiple_Declarations_One_Type_One_Set_Value IntNodeType
Constant_Global_Variable IntNodeType
Constant_Global_Variable_With_Type IntNodeType
Struct_Variable_Declaration IntNodeType
If_Statement_With_Declaration IntNodeType
If_Statement IntNodeType
Else_Statement IntNodeType
For_Loop IntNodeType
For_Loop_With_Declaration IntNodeType
Err_Return IntNodeType
Err_Check IntNodeType
Multi_Line_Import IntNodeType
Single_Import IntNodeType
Single_Import_With_Alias IntNodeType
NestedStruct IntNodeType
Struct_Declaration IntNodeType
Enum_Declaration IntNodeType
Enum_Variable IntNodeType
Enum_Variable_With_Value IntNodeType
Enum_Struct_Declaration IntNodeType
Enum_Struct_Declaration_With_Alias IntNodeType
NewLine IntNodeType
Comment IntNodeType
Append IntNodeType
Package IntNodeType
Other IntNodeType
Switch IntNodeType
Switch_Case_Body IntNodeType
Switch_Case IntNodeType
Return IntNodeType
Break IntNodeType
Continue IntNodeType
}{
Invalid: 0,
Channel_Declaration: 1,
Channel_Declaration_With_Value: 2,
Interface_Declaration: 3,
Single_Declaration_With_Value: 4,
Single_Declaration_No_Value: 5,
Multiple_Declarations_No_Value: 6,
Multiple_Declarations_With_Value: 7,
Multiple_Declarations_Same_Type_No_Value: 8,
Multiple_Declarations_Same_Type_With_Value: 9,
Multiple_Declarations_One_Type_One_Set_Value: 10,
Constant_Global_Variable: 11,
Constant_Global_Variable_With_Type: 12,
Struct_Variable_Declaration: 13,
If_Statement_With_Declaration: 14,
If_Statement: 15,
Else_Statement: 16,
For_Loop: 17,
For_Loop_With_Declaration: 18,
Err_Return: 19,
Err_Check: 20,
Multi_Line_Import: 21,
Single_Import: 22,
Single_Import_With_Alias: 23,
NestedStruct: 24,
Struct_Declaration: 25,
Enum_Declaration: 26,
Enum_Variable: 27,
Enum_Variable_With_Value: 28,
Enum_Struct_Declaration: 29,
Enum_Struct_Declaration_With_Alias: 30,
NewLine: 31,
Comment: 32,
Append: 33,
Package: 34,
Other: 35,
Switch: 36,
Switch_Case_Body: 37,
Switch_Case: 38,
Return: 39,
Break: 40,
Continue: 41,
}

func (self IntNodeType) ToString() string {
	switch self {
	case NodeType.Invalid:
		return "Invalid"
	case NodeType.Channel_Declaration:
		return "Channel_Declaration"
	case NodeType.Channel_Declaration_With_Value:
		return "Channel_Declaration_With_Value"
	case NodeType.Interface_Declaration:
		return "Interface_Declaration"
	case NodeType.Single_Declaration_With_Value:
		return "Single_Declaration_With_Value"
	case NodeType.Single_Declaration_No_Value:
		return "Single_Declaration_No_Value"
	case NodeType.Multiple_Declarations_No_Value:
		return "Multiple_Declarations_No_Value"
	case NodeType.Multiple_Declarations_With_Value:
		return "Multiple_Declarations_With_Value"
	case NodeType.Multiple_Declarations_Same_Type_No_Value:
		return "Multiple_Declarations_Same_Type_No_Value"
	case NodeType.Multiple_Declarations_Same_Type_With_Value:
		return "Multiple_Declarations_Same_Type_With_Value"
	case NodeType.Multiple_Declarations_One_Type_One_Set_Value:
		return "Multiple_Declarations_One_Type_One_Set_Value"
	case NodeType.Constant_Global_Variable:
		return "Constant_Global_Variable"
	case NodeType.Constant_Global_Variable_With_Type:
		return "Constant_Global_Variable_With_Type"
	case NodeType.Struct_Variable_Declaration:
		return "Struct_Variable_Declaration"
	case NodeType.If_Statement_With_Declaration:
		return "If_Statement_With_Declaration"
	case NodeType.If_Statement:
		return "If_Statement"
	case NodeType.Else_Statement:
		return "Else_Statement"
	case NodeType.For_Loop:
		return "For_Loop"
	case NodeType.For_Loop_With_Declaration:
		return "For_Loop_With_Declaration"
	case NodeType.Err_Return:
		return "Err_Return"
	case NodeType.Err_Check:
		return "Err_Check"
	case NodeType.Multi_Line_Import:
		return "Multi_Line_Import"
	case NodeType.Single_Import:
		return "Single_Import"
	case NodeType.Single_Import_With_Alias:
		return "Single_Import_With_Alias"
	case NodeType.NestedStruct:
		return "NestedStruct"
	case NodeType.Struct_Declaration:
		return "Struct_Declaration"
	case NodeType.Enum_Declaration:
		return "Enum_Declaration"
	case NodeType.Enum_Variable:
		return "Enum_Variable"
	case NodeType.Enum_Variable_With_Value:
		return "Enum_Variable_With_Value"
	case NodeType.Enum_Struct_Declaration:
		return "Enum_Struct_Declaration"
	case NodeType.Enum_Struct_Declaration_With_Alias:
		return "Enum_Struct_Declaration_With_Alias"
	case NodeType.NewLine:
		return "NewLine"
	case NodeType.Comment:
		return "Comment"
	case NodeType.Append:
		return "Append"
	case NodeType.Package:
		return "Package"
	case NodeType.Other:
		return "Other"
	case NodeType.Switch:
		return "Switch"
	case NodeType.Switch_Case_Body:
		return "Switch_Case_Body"
	case NodeType.Switch_Case:
		return "Switch_Case"
	case NodeType.Return:
		return "Return"
	case NodeType.Break:
		return "Break"
	case NodeType.Continue:
		return "Continue"
	default:
		return "Unknown"
}
}

type IntLoopAction int
var LoopAction = struct {
	Continue IntLoopAction
Break IntLoopAction
Return IntLoopAction
Error IntLoopAction
}{
Continue: 0,
Break: 1,
Return: 2,
Error: 3,
}

func (self IntLoopAction) ToString() string {
	switch self {
	case LoopAction.Continue:
		return "Continue"
	case LoopAction.Break:
		return "Break"
	case LoopAction.Return:
		return "Return"
	case LoopAction.Error:
		return "Error"
	default:
		return "Unknown"
}
}

type IntIdentLoopAction int
var IdentLoopAction = struct {
	Continue IntIdentLoopAction
Break IntIdentLoopAction
Declaration IntIdentLoopAction
Other IntIdentLoopAction
Append IntIdentLoopAction
Error IntIdentLoopAction
}{
Continue: 0,
Break: 1,
Declaration: 2,
Other: 3,
Append: 4,
Error: 5,
}

func (self IntIdentLoopAction) ToString() string {
	switch self {
	case IdentLoopAction.Continue:
		return "Continue"
	case IdentLoopAction.Break:
		return "Break"
	case IdentLoopAction.Declaration:
		return "Declaration"
	case IdentLoopAction.Other:
		return "Other"
	case IdentLoopAction.Append:
		return "Append"
	case IdentLoopAction.Error:
		return "Error"
	default:
		return "Unknown"
}
}

type IntParameterPhase int
var ParameterPhase = struct {
	TypeOrName IntParameterPhase
	End IntParameterPhase
	}{
	TypeOrName: 0,
	End: 1,
	}

func (self IntParameterPhase) ToString() string {
	switch self {
	case ParameterPhase.TypeOrName:
		return "TypeOrName"
	case ParameterPhase.End:
		return "End"
	default:
		return "Unknown"
	}
}


type IntMethodType int
var MethodType = struct {
	None IntMethodType
Struct IntMethodType
Enum IntMethodType
}{
None: 0,
Struct: 1,
Enum: 2,
}

func (self IntMethodType) ToString() string {
	switch self {
	case MethodType.None:
		return "None"
	case MethodType.Struct:
		return "Struct"
	case MethodType.Enum:
		return "Enum"
	default:
		return "Unknown"
}
}

