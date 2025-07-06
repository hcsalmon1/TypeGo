using System;
using System.Collections.Generic;


namespace TypeGo
{
    public enum SyntaxType
    {
        Function, //fn
        FunctionName, //anything not a type
        VariableName, //anything not a type
        VariableType, //i32, bool, string
        Operator, //+, -, /, *
        Seperator, //;, (, )
        Equals, //=
        Command, //for, if, return
        Number,
        Invalid,
    }


    public struct PossibleValues
    {
        public bool Numbers;
        public bool TypeDeclaration;
        public bool FunctionDeclaration;
        public bool Command;
        public bool Number;
        public bool Seperator;
        public bool Operator;

        public void Reset()
        {
            Numbers = true;
            TypeDeclaration = true;
            FunctionDeclaration = true;
            Command = true;
            Number = true;
            Seperator = true;
            Operator = true;
        }
        public void PrintValues()
        {
            Fmt.Println($"PossibleValues:\n numbers: {Numbers}\n TypeDeclaration: {TypeDeclaration}\n Function declaration: {FunctionDeclaration}\n Command: {Command}\n Number: {Number}\n Seperator: {Seperator}\n Operator: {Operator}");
        }
    }

    public static class SYNTAX
    {
        public static readonly char[] OPERATORS = {'+','-','/','*', '|', '&', '%', '>', '<', '=', ':', '!' };
        public static readonly char[] SEPERATORS = {';','(',')','{','}','[',']',',','.','\n','\r','\t','\\'};
        public static readonly string[] VARIABLE_KEYWORDS = {"i32", "bool", "string" };
        public const string FUNCTION_KEYWORD = "fn";
        public const int MAX_TOKEN_SIZE = 100;
    }
}
