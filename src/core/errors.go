
package core

import "fmt"


type IntConvertResult int
var ConvertResult = struct {
	Ok IntConvertResult
Missing_Expected_Type IntConvertResult
Unexpected_Type IntConvertResult
Unexpected_End_Of_File IntConvertResult
No_Token_In_Node IntConvertResult
Null_Token IntConvertResult
Invalid_Node_Type IntConvertResult
Unsupported_Type IntConvertResult
Internal_Error IntConvertResult
}{
Ok: 0,
Missing_Expected_Type: 1,
Unexpected_Type: 2,
Unexpected_End_Of_File: 3,
No_Token_In_Node: 4,
Null_Token: 5,
Invalid_Node_Type: 6,
Unsupported_Type: 7,
Internal_Error: 8,
}

func (self IntConvertResult) ToString() string {
	switch self {
	case ConvertResult.Ok:
		return "Ok"
	case ConvertResult.Missing_Expected_Type:
		return "Missing_Expected_Type"
	case ConvertResult.Unexpected_Type:
		return "Unexpected_Type"
	case ConvertResult.Unexpected_End_Of_File:
		return "Unexpected_End_Of_File"
	case ConvertResult.No_Token_In_Node:
		return "No_Token_In_Node"
	case ConvertResult.Null_Token:
		return "Null_Token"
	case ConvertResult.Invalid_Node_Type:
		return "Invalid_Node_Type"
	case ConvertResult.Unsupported_Type:
		return "Unsupported_Type"
	case ConvertResult.Internal_Error:
		return "Internal_Error"
	default:
		return "Unknown"
}
}


func (self IntConvertResult) IsError() bool {
	return self != ConvertResult.Ok
}

func (self IntConvertResult) Println() {
	fmt.Println("Errors:", self.ToString())
	
}
type IntParseResult int
var ParseResult = struct {
	Ok IntParseResult
String_Length_Zero IntParseResult
Starting_Index_Out_Of_Range_Of_String IntParseResult
Index_Minus IntParseResult
Invalid_Token IntParseResult
Current_Index_Out_Of_Range IntParseResult
Buffer_Index_Over_Max IntParseResult
Unexpected_Value IntParseResult
Unterminated_Char IntParseResult
Unterminated_String IntParseResult
Unterminated_Comment IntParseResult
Infinite_While_Loop IntParseResult
}{
Ok: 0,
String_Length_Zero: 1,
Starting_Index_Out_Of_Range_Of_String: 2,
Index_Minus: 3,
Invalid_Token: 4,
Current_Index_Out_Of_Range: 5,
Buffer_Index_Over_Max: 6,
Unexpected_Value: 7,
Unterminated_Char: 8,
Unterminated_String: 9,
Unterminated_Comment: 10,
Infinite_While_Loop: 11,
}

func (self IntParseResult) ToString() string {
	switch self {
	case ParseResult.Ok:
		return "Ok"
	case ParseResult.String_Length_Zero:
		return "String_Length_Zero"
	case ParseResult.Starting_Index_Out_Of_Range_Of_String:
		return "Starting_Index_Out_Of_Range_Of_String"
	case ParseResult.Index_Minus:
		return "Index_Minus"
	case ParseResult.Invalid_Token:
		return "Invalid_Token"
	case ParseResult.Current_Index_Out_Of_Range:
		return "Current_Index_Out_Of_Range"
	case ParseResult.Buffer_Index_Over_Max:
		return "Buffer_Index_Over_Max"
	case ParseResult.Unexpected_Value:
		return "Unexpected_Value"
	case ParseResult.Unterminated_Char:
		return "Unterminated_Char"
	case ParseResult.Unterminated_String:
		return "Unterminated_String"
	case ParseResult.Unterminated_Comment:
		return "Unterminated_Comment"
	case ParseResult.Infinite_While_Loop:
		return "Infinite_While_Loop"
	default:
		return "Unknown"
}
}


func (self IntParseResult) IsError() bool {
	return self != ParseResult.Ok
}

func (self IntParseResult) Println() {
	fmt.Println("Errors:", self.ToString())
	
}
type IntFormatResult int
var FormatResult = struct {
	Ok IntFormatResult
NoTokens IntFormatResult
MissingExpectedType IntFormatResult
UnexpectedType IntFormatResult
EndOfFile IntFormatResult
UnsupportedFeature IntFormatResult
Infinite_While_Loop IntFormatResult
Internal_Error IntFormatResult
Invalid_Node IntFormatResult
}{
Ok: 0,
NoTokens: 1,
MissingExpectedType: 2,
UnexpectedType: 3,
EndOfFile: 4,
UnsupportedFeature: 5,
Infinite_While_Loop: 6,
Internal_Error: 7,
Invalid_Node: 8,
}

func (self IntFormatResult) ToString() string {
	switch self {
	case FormatResult.Ok:
		return "Ok"
	case FormatResult.NoTokens:
		return "NoTokens"
	case FormatResult.MissingExpectedType:
		return "MissingExpectedType"
	case FormatResult.UnexpectedType:
		return "UnexpectedType"
	case FormatResult.EndOfFile:
		return "EndOfFile"
	case FormatResult.UnsupportedFeature:
		return "UnsupportedFeature"
	case FormatResult.Infinite_While_Loop:
		return "Infinite_While_Loop"
	case FormatResult.Internal_Error:
		return "Internal_Error"
	case FormatResult.Invalid_Node:
		return "Invalid_Node"
	default:
		return "Unknown"
}
}


func (self IntFormatResult) IsError() bool {
	return self != FormatResult.Ok
}

func (self IntFormatResult) Println() {
	fmt.Println("Errors:", self.ToString())
	
}
