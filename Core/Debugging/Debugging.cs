using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Text;

namespace TypeGo
{
    public enum IncrementError
    {
        None,
        IndexTooHigh,
        IndexTooLow,
    }

    public static class Debugging
    {

        static void PrintTypeParameter(ref readonly Variable parameter)
        {
            if (parameter.TypeList == null)
            {
                Fmt.Println("\tparameter typelist is null");
                return;
            }
            if (parameter.TypeList.Count == 0)
            {
                Fmt.Println("\tparameter typelist count is 0");
                return;
            }
            Fmt.Print("\tType: ");
            for (int j = 0; j < parameter.TypeList.Count; j++)
            {
                Token typeToken = parameter.TypeList[j];
                Fmt.Print(typeToken.Text);
            }
            Fmt.Println();
        }

        static void PrintParameterName(Token? nameToken)
        {
            if (nameToken == null)
            {
                Fmt.Println($"\tname token: null");
                return;
            }
            Fmt.Println($"\tname: {nameToken.Value.Text}");
        }

        public static void PrintParameters(List<Variable> parameters)
        {
            Fmt.PrintlnColor("\nPrinting: parameters:", ConsoleColor.DarkCyan);
            if (parameters == null) {
                Fmt.Println("\tList is null");
                return;
            }
            if (parameters.Count == 0) {
                Fmt.Println("\tList is empty");
                return;
            }
            for (int i = 0; i < parameters.Count; i++) {
                Variable parameter = parameters[i];
                Fmt.Println($"index: {i}");
                PrintTypeParameter(in parameter);

                Token? nameToken = parameters[i].NameToken;
                PrintParameterName(nameToken);
            }
        }
        public static bool InfiniteWhileCheck(ref int count, int CAP)
        {
            count += 1;
            if (count >= CAP) {
                return true;
            }
            return false;
        }

        public static void PrintTokenList(List<Token>? tokenList)
        {
            if (tokenList == null) {
                Fmt.Println("  Error: token list is null");
                return;
            }
            if (tokenList.Count == 0) {
                Fmt.Println("  Error: token list is zero");
                return;
            }
            Fmt.PrintlnColor("\nFound tokens:", ConsoleColor.Cyan);
            for (int i = 0; i < tokenList.Count; i++) {

                Fmt.PrintColor("\ttoken:", ConsoleColor.Gray);
                Fmt.PrintColor($" '{tokenList[i].Text}', ", ConsoleColor.DarkCyan);
                Fmt.PrintColor("type: ", ConsoleColor.Gray);
                Fmt.PrintColor($"{tokenList[i].Type}", ConsoleColor.Yellow);
                Fmt.PrintColor($", lineCount: ", ConsoleColor.Gray);
                Fmt.PrintlnColor($"{tokenList[i].LineNumber}", ConsoleColor.White);
            }
        }

        public static void PrintError(ParseResult error)
        {
            Fmt.Println($"ERROR: {error}");
        }

        public static IncrementError IncrementIndex(ref int index, int count)
        {
            if (index + 1 >= count) {
                return IncrementError.IndexTooHigh;
            }
            index += 1;
            return IncrementError.None;
        }

        static int GetLineNumber(Token? error_token)
        {
            if (error_token == null) {
                return -1;
            }
            return error_token.Value.LineNumber;
        }
        static int GetCharNumber(Token? error_token)
        {
            if (error_token == null)
            {
                return -1;
            }
            return error_token.Value.CharNumber;
        }

        static string GetCodeLine(int line_number, string[] codeLines)
        {
            if (line_number >= codeLines.Length || line_number < 0) {
                return "Invalid";
            } else {
                return codeLines[line_number];
            }
        }


        private static void PrintErrorTrace(List<string> errorTrace)
        {
            if (errorTrace == null)
            {
                Fmt.Println("\nWarning: error trace is null\n");
                return;
            }
            if (errorTrace.Count == 0)
            {
                Fmt.Println("\nWarning: error trace count is 0\n");
                return;
            }

            Fmt.PrintlnColor($"\n\tError trace:", ConsoleColor.DarkGray);

            for (int i = 0; i < errorTrace.Count; i += 1)
            {
                if (i != 0)
                {
                    Fmt.Print(" <- ");
                }
                else
                {
                    Fmt.Print("\t");
                }
                string function_name = errorTrace[i];
                Fmt.Print(function_name);
            }
            Fmt.Println(' ');
            Fmt.Println(' ');
        }

        private static void PrintErrorToken(Token? error_token)
        {
            Fmt.Print("\tToken: ");
            if (error_token.HasValue == true)
            {
                Fmt.PrintlnColor($"'{error_token.Value.Text}'", ConsoleColor.DarkGreen);
            }
            else
            {
                Fmt.Println("Error token not set");
            }
        }

        private static void PrintCodeLines(int line_number, string[] codeLines, string code_line)
        {
            if (line_number - 1 >= 0 && line_number < codeLines.Length)
            {
                Fmt.Println($"\tLine {line_number}: '{codeLines[line_number - 1]}'");
            }
            Fmt.Println($"\tLine {line_number + 1}: '{code_line}'");
        }

        public static void PrintASTNodes(List<ASTNode> ast_nodes) {

            Fmt.PrintlnColor("\nPrinting Ast Nodes:", ConsoleColor.Cyan);

            if (ast_nodes.Count == 0) {
                Fmt.PrintlnColor("\tNo nodes\n", ConsoleColor.Red);
                return;
            }
            for (int i = 0; i < ast_nodes.Count; i += 1) {
                ASTNode node = ast_nodes[i];
                PrintASTNode(node, 3, "base");
            }
        }

        static void PrintASTNode(ASTNode? node, int indent, string ast_type_text) {
            if (node == null) {
                return;
            }

            StringBuilder padding = new();
        
            for (int i = 0; i < indent; i++)
            {
                AddSpacing(indent, padding, i);
            }

            //node type
            Fmt.Print(padding.ToString());
            Fmt.PrintColor($"{node.NodeType} ", ConsoleColor.Yellow);

            //token
            if (node.Token != null) {
                Fmt.PrintColor($"'{node.Token.Value.Text}'", ConsoleColor.DarkCyan);
                Fmt.Println($" - {ast_type_text}");
            }
            else
            {
                Fmt.Println($"NA - {ast_type_text}");
            }

            // Recursively print children with increased indent
            if (node.Left != null) {
                PrintASTNode(node.Left, indent + 1, "left");
            }
            if (node.Middle != null) {
                PrintASTNode(node.Middle, indent + 1, "middle");
            }
            if (node.Right != null) {
                PrintASTNode(node.Right, indent + 1, "right");
            }

            if (node.Children != null) {
                int children_count = node.Children.Count;
                for (int i = 0; i < children_count; i += 1) {
                    PrintASTNode(node.Children[i], indent + 1, "child");
                }
            }
        }

        private static void AddSpacing(int indent, StringBuilder padding, int i)
        {
            if (i + 1 == indent)
            {
                padding.Append('|');
                padding.Append('-');
                return;
            }
            if (i == 2)
            {
                padding.Append('|');
                padding.Append(' ');
                return;
            }
            if (i == 3)
            {
                padding.Append('|');
                padding.Append(' ');
                return;
            }
            
            padding.Append(' ');
            padding.Append(' ');
            
        }

   
        public static void PrintGlobalDefinitions(ref readonly GlobalDefinitions globalDefinitions)
        {
            Fmt.PrintlnColor("\nPrinting Global Definitions", ConsoleColor.Cyan);
            PrintGlobalVariables(in globalDefinitions);
            PrintFunctionDefinitions(in globalDefinitions);
        }

        private static void PrintGlobalVariables(ref readonly GlobalDefinitions globalDefinitions)
        {
            Fmt.Println("\tGlobal variables:");
            if (globalDefinitions.Globals.Count == 0)
            {
                Fmt.PrintlnColor("\t\tNo global variables defined", ConsoleColor.DarkYellow);
                return;
            }
            for (int i = 0; i < globalDefinitions.Globals.Count; i++)
            {
                ASTNode? definitionNode = globalDefinitions.Globals[i];
                if (definitionNode == null)
                {
                    Fmt.PrintlnColor("\t\tNode is null", ConsoleColor.Red);
                    continue;
                }
                Fmt.PrintColor("\t\tType: ", ConsoleColor.DarkGray);
                Debug.Assert(definitionNode.Left != null);
                ASTNode? tempLeftNode = definitionNode.Left;
                while (tempLeftNode != null)
                {
                    Debug.Assert(tempLeftNode.Token != null);
                    Fmt.PrintColor($"{tempLeftNode.Token.Value.Text} ", ConsoleColor.DarkCyan);
                    tempLeftNode = tempLeftNode.Left;
                }
                Debug.Assert(definitionNode.Right != null);
                Debug.Assert(definitionNode.Right.Token != null);
                Fmt.PrintlnColor(definitionNode.Right.Token.Value.Text, ConsoleColor.Cyan);
            }
        }

        private static void PrintFunctionDefinitions(ref readonly GlobalDefinitions globalDefinitions)
        {
            Fmt.Println("\tFunctions:");
            if (globalDefinitions.Functions.Count == 0)
            {
                Fmt.PrintlnColor("\t\tNo functions defined", ConsoleColor.DarkYellow);
                return;
            }
            for (int functionIndex = 0; functionIndex < globalDefinitions.Functions.Count; functionIndex++)
            {
                ASTNode? functionNode = globalDefinitions.Functions[functionIndex];
                if (functionNode == null)
                {
                    Fmt.PrintlnColor("\t\tNode is null", ConsoleColor.Red);
                    continue;
                }
                Fmt.PrintColor("\t\tType: ", ConsoleColor.DarkGray);

                Debug.Assert(functionNode.Left != null);

                //Print function type
                ASTNode? tempLeftNode = functionNode.Left;
                while (tempLeftNode != null)
                {
                    Debug.Assert(tempLeftNode.Token != null);
                    Fmt.PrintColor($"{tempLeftNode.Token.Value.Text} ", ConsoleColor.DarkCyan);
                    tempLeftNode = tempLeftNode.Left;
                }

                //Print function name
                Debug.Assert(functionNode.Token != null);
                Fmt.PrintlnColor(functionNode.Token.Value.Text, ConsoleColor.Yellow);

                //Print parameters:
                Fmt.PrintColor("\t\tParameters: ", ConsoleColor.DarkGray);

                Debug.Assert(functionNode.Middle != null);
                Debug.Assert(functionNode.Middle.Children != null);

                if (functionNode.Middle.Children.Count == 0)
                {
                    Fmt.PrintlnColor(" None\n", ConsoleColor.White);
                }

                //Print function type
                for (int parameterIndex = 0; parameterIndex < functionNode.Middle.Children.Count; parameterIndex++)
                {
                    ASTNode? tempChildNode = functionNode.Middle.Children[parameterIndex].Left;
                    while (tempChildNode != null) {
                        //Debug.Assert(tempChildNode.Token != null);
                        if (tempChildNode.Token == null)
                        {
                            if (tempChildNode.NodeType == ASTNodeType.Pointer) {

                            }
                            Fmt.PrintColor($"NULL token ", ConsoleColor.Red);
                            tempChildNode = tempChildNode.Left;
                            continue;
                        }
                        Fmt.PrintColor($"{tempChildNode.Token.Value.Text} ", ConsoleColor.DarkCyan);
                        tempChildNode = tempChildNode.Left;
                    }
                    Debug.Assert(functionNode.Middle.Children[parameterIndex] != null);
                    Debug.Assert(functionNode.Middle.Children[parameterIndex].Token != null);
                    Fmt.PrintColor($"{functionNode.Middle.Children[parameterIndex].Token.Value.Text}, ", ConsoleColor.Cyan);
                }
                Fmt.Println('\n');
            }
        }

        public static void PrintConvertError(ref ConvertData convertData, string code)
        {
            Token? error_token = convertData.ErrorToken;

            int line_number = GetLineNumber(error_token);
            int char_number = GetCharNumber(error_token);

            string[] codeLines = code.Split(new[] { "\r\n", "\n" }, StringSplitOptions.None); // Split code into lines

            string code_line = GetCodeLine(line_number, codeLines);

            Fmt.PrintlnColor($"\tError on line {line_number + 1}, {char_number}: {convertData.ErrorDetail}", ConsoleColor.Yellow);

            PrintCodeLines(line_number, codeLines, code_line);

            PrintErrorToken(error_token);

            //PrintErrorTrace(convertData.ErrorTrace);
        }

        public static void PrintFormatError(ref FormatData formatData, string code)
        {
            Token? error_token = formatData.ErrorToken;

            int line_number = GetLineNumber(error_token);
            int char_number = GetCharNumber(error_token);

            string[] codeLines = code.Split(new[] { "\r\n", "\n" }, StringSplitOptions.None); // Split code into lines

            string code_line = GetCodeLine(line_number, codeLines);

            Fmt.PrintlnColor($"\tError on line {line_number}, {char_number}: {formatData.ErrorDetail}", ConsoleColor.Yellow);

            PrintCodeLines(line_number, codeLines, code_line);

            PrintErrorToken(error_token);

            //PrintErrorTrace(formatData.ErrorTrace);
        }
    
        
    
    }
}
