using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public enum ConvertResult
    {
        Ok,
        Missing_Expected_Type,
        Unexpected_Type,
        Unexpected_End_Of_File,
        No_Token_In_Node,
        Null_Token,
        Invalid_Node_Type,
        Unsupported_Type,
        Internal_Error,
    }

    public enum ParseResult
    {
        Ok,
        String_Length_Zero,
        Starting_Index_Out_Of_Range_Of_String,
        Index_Minus,
        Invalid_Token,
        Current_Index_Out_Of_Range,
        Buffer_Index_Over_Max,
        Unexpected_Value,
        Unterminated_Char,
        Unterminated_String,
    }

    public enum FormatResult
    {
        Ok,
        NoTokens,
        MissingExpectedType,
        UnexpectedType,
        EndOfFile,
        UnsupportedFeature,
        Infinite_While_Loop,
        Internal_Error,
    }

}
