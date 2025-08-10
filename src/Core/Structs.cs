using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public class CodeFormat
    {
        public CodeBlock GlobalBlock = new CodeBlock();
        public List<Function> Functions = new List<Function>();

        public void PrintFunctions()
        {
            Fmt.PrintlnColor("Printing Functions:", ConsoleColor.Cyan);

            foreach (Function function in Functions) {

                function.PrintData();
            }
        }

        public void PrintNodeType()
        {
            Fmt.PrintlnColor("Printing Functions:", ConsoleColor.Cyan);

            foreach (Function function in Functions) {

                function.PrintNodeTypes();
            }
        }
    }

    public readonly struct ConstInt
    {
        public readonly int Value { get; }

        public ConstInt(int value) {
            Value = value;
        }
    }
    public class Function
    {
        public string? ReturnType = "";
        public string Name = "";
        public List<Variable> Parameters = new List<Variable>();
        public CodeBlock? InnerBlock = null;
        public Token? startingToken;

        public void PrintData()
        {
            Fmt.PrintColor($"\t{ReturnType} ", ConsoleColor.DarkCyan);
            Fmt.PrintColor($"{Name}(", ConsoleColor.Yellow);
            foreach (Variable parameter in Parameters)
            {
                parameter.Print();
                Fmt.Print(", ");
            }
            Fmt.Println(')');
            if (InnerBlock != null) {
                InnerBlock.PrintData();
            }
        }

        public void PrintNodeTypes()
        {
            Fmt.Println("Printing Node Types:");
            if (InnerBlock == null) {
                Fmt.Println("\tinner block is null!:");
                return;
            }
            CodeBlock block = InnerBlock;

            if (block.BlockDataList.Count == 0) {
                Fmt.Println("\tblock data list is 0!");
                return;
            }

            for (int i = 0; i < block.BlockDataList.Count; i++) {
                BlockData blockData = block.BlockDataList[i];
                Fmt.Println($"\t{blockData.NodeType}");
            }
        }
    }

    public class CodeBlock
    {
        public List<BlockData> BlockDataList = new List<BlockData>();
        public List<Function> MethodList = new List<Function>();
        public void PrintData()
        {
            Fmt.PrintlnColor("\nPrinting blockdata:", ConsoleColor.Cyan);
            if (BlockDataList == null) {
                Fmt.PrintlnColor("\tblock data is null\n", ConsoleColor.Red);
                return;
            }
            if (BlockDataList.Count == 0) {
                Fmt.PrintlnColor("\tempty\n", ConsoleColor.Red);
            }
            for (int i = 0; i < BlockDataList.Count; i++) {
                BlockDataList[i].Print();
                Fmt.Println();
            }
            Fmt.Println();
        }
    }

    public class BlockData
    {
        public NodeType NodeType = NodeType.Invalid;
        public List<Token> Tokens = new List<Token>();
        public Token? StartingToken = null;
        public CodeBlock? Block = null;
        public List<Variable> Variables = new List<Variable>();

        public void Print()
        {
            const bool NEW_LINE = true;

            PrintNodeType(NEW_LINE);
            PrintStartingToken(NEW_LINE);
            PrintTokens(NEW_LINE);
            PrintVariables(NEW_LINE);

            if (Block == null) {
                Fmt.PrintColor("\tBlock: ", ConsoleColor.DarkGray);
                Fmt.PrintColor("null ", ConsoleColor.Yellow);
            } else {
                Fmt.PrintColor("\tBlock: ", ConsoleColor.DarkGray);
                Fmt.PrintColor("not null ", ConsoleColor.Yellow);
            }

            StringBuilder sb = new();
            sb.Append("\n\tVariable: ");
            foreach (Variable variable in Variables)
            {
                string varTypeAsText = variable.ConvertToString();
                sb.Append(varTypeAsText);
            }
            string? tokensAsText = TokenUtils.JoinTextInListOfTokensWithSpaces(Tokens);
            if (tokensAsText == null) {
                sb.Append(" null");
            } else {
                sb.Append(tokensAsText);
            }
            Fmt.Println(sb.ToString());
        }
        public void PrintNodeType(bool newLine)
        {
            Fmt.PrintColor("\tNode type: ", ConsoleColor.DarkGray);
            if (NodeType == NodeType.Invalid) {
                Fmt.PrintColor($"Invalid ", ConsoleColor.Red);
                AddNewLine(newLine);
                return;
            }
            Fmt.PrintColor($"{NodeType} ", ConsoleColor.Yellow);
            AddNewLine(newLine);
        }
        public void PrintStartingToken(bool newLine)
        {
            if (newLine) {
                Fmt.Print('\t');
            }
            if (StartingToken == null)
            {
                Fmt.PrintColor("Starting token: ", ConsoleColor.DarkGray);
                Fmt.PrintColor("null ", ConsoleColor.Red);
                AddNewLine(newLine);
                return;
            }
            Fmt.PrintColor("Starting token: ", ConsoleColor.DarkGray);
            Fmt.PrintColor($"'{StartingToken.Value.Text}' ", ConsoleColor.DarkCyan);
            AddNewLine(newLine);
        }
        public void AddNewLine(bool newLine) {
            if (newLine == true) {
                Fmt.Println();
            }
        }
        public void PrintTokens(bool newLine)
        {
            if (newLine) {
                Fmt.Print('\t');
            }
            Fmt.PrintColor("tokens: ", ConsoleColor.DarkGray);
            if (Tokens.Count == 0) {
                Fmt.PrintColor("empty ", ConsoleColor.Red);
                AddNewLine(newLine);
                return;
            }
            for (int i = 0; i < Tokens.Count; i++) {

                Fmt.PrintColor($"'{Tokens[i].Text}' ", ConsoleColor.DarkCyan);
            }
            AddNewLine(newLine);
        }
        public void PrintWithCodeBlock()
        {
            const bool NO_NEW_LINE = false;

            Fmt.PrintColor("\tNode type: ", ConsoleColor.DarkGray);
            Fmt.PrintColor($"{NodeType} ", ConsoleColor.Yellow);
            PrintStartingToken(NO_NEW_LINE);
            PrintTokens(NO_NEW_LINE);

            if (Block == null) {
                Fmt.PrintColor("Block: ", ConsoleColor.DarkGray);
                Fmt.PrintColor("null ", ConsoleColor.Yellow);
            } else {
                Fmt.PrintColor("Block: ", ConsoleColor.DarkGray);
                Fmt.PrintColor("not null ", ConsoleColor.Yellow);
            }
        }
        public void PrintVariables(bool newLine)
        {
            if (newLine) {
                Fmt.Print('\t');
            }
            Fmt.PrintColor("Vars: ", ConsoleColor.DarkGray);
            if (Variables.Count == 0) {
                Fmt.PrintColor("empty ", ConsoleColor.Magenta);
                AddNewLine(newLine);
                return;
            }
            for (int i = 0; i < Variables.Count; i++) {

                Variables[i].Print();
            }
            AddNewLine(newLine);
        }
    }

    public class FormatData
    {
        public List<Token> TokenList;
        public FormatResult Result = FormatResult.Ok;
        public string ErrorDetail = "";
        public List<string> ErrorTrace = new List<string>();
        public Token? ErrorToken = null;
        public int TokenIndex;
        public Token? TempToken = null;

        public Token? GetToken() {

            if (TokenIndex >= TokenList.Count) {
                return null;
            }
            return TokenList[TokenIndex];
        }
        public Token? GetTokenByIndex(int index) {

            if (index >= TokenList.Count) {
                return null;
            }
            return TokenList[index];
        }
        public void Increment() {
            TokenIndex += 1;
        }
        public Token? GetNextToken() {

            Increment();
            return GetToken();
        }
        public void IncrementIfSame(int indexBefore)
        {
            if (indexBefore == TokenIndex) {
                Increment();
            }
        }
        public void EndOfFileError(Token? errorToken, string thisFunctionName) {
            if (IsError()) {
                return;
            }
            Result = FormatResult.EndOfFile;
            ErrorTrace.Add(thisFunctionName);
            ErrorToken = errorToken;
        }
        public void UnsupportedFeatureError(Token? errorToken, string thisFunctionName) {
            if (IsError()) {
                return;
            }
            Result = FormatResult.UnsupportedFeature;
            ErrorTrace.Add(thisFunctionName);
            ErrorToken = errorToken;
        }
        public bool IsError() {
            return Result != FormatResult.Ok;
        }

        public void AddTrace(string functionName)
        {
            ErrorTrace.Add(functionName);
        }
        public void MissingExpectedTypeError(Token? errorToken, string detail, string thisFunctionName)
        {
            if (IsError()) {
                return;
            }
            Result = FormatResult.MissingExpectedType;
            ErrorDetail = detail;
            ErrorTrace.Add(thisFunctionName);
            ErrorToken = errorToken;
        }
        public void UnexpectedTypeError(Token? errorToken, string detail, string thisFunctionName)
        {
            if (IsError()) {
                return;
            }
            Result = FormatResult.UnexpectedType;
            ErrorDetail = detail;
            ErrorTrace.Add(thisFunctionName);
            ErrorToken = errorToken;
        }
        public bool ExpectType(TokenType type, string detail, string thisFunctionName)
        {
            Token? token = GetToken();
            if (token == null) {
                EndOfFileError(null, thisFunctionName);
                return false;
            }
            if (token.Value.Type == type) {
                return true;
            }
            MissingExpectedTypeError(token, detail, thisFunctionName);
            return false;
        }
        public bool ExpectNextType(TokenType type, string detail, string thisFunctionName)
        {
            Token? token = GetNextToken();
            if (token == null) {
                EndOfFileError(null, thisFunctionName);
                return false;
            }
            if (token.Value.Type == type) {
                return true;
            }
            MissingExpectedTypeError(token, detail, thisFunctionName);
            return false;
        }
        public void IncrementIfNewLine()
        {
            Token? token = GetToken();
            if (token == null) {
                return;
            }
            if (token.Value.Type == TokenType.NewLine) {
                Increment();
            }
        }
        public void UnsupportedFeatureError(Token? errorToken, string detail, string thisFunctionName)
        {
            if (IsError()) {
                return;
            }
            Result = FormatResult.UnsupportedFeature;
            ErrorToken = errorToken;
            ErrorDetail = detail;
            ErrorTrace.Add(thisFunctionName);
        }
    }

    public class ASTNode
    {
        public ASTNodeType NodeType;

        // These fields are optional depending on the kind
        public Token? Token;       // Token, for literals, var types etc

        public ASTNode? Left;     // For expressions (e.g., a + b), return types
        public ASTNode? Middle;   // Rare for things with 3, for loop, parameters
        public ASTNode? Right;      // right side of arithmetic a + b;, right = b
        public List<ASTNode?>? Children;    // For blocks, function bodies, etc.

        public bool IsArray;
        public bool IsGlobal;
        public bool IsConst;
        public int? Size;

        public void MakeAllNull()
        {
            Left = null;
            Middle = null;
            Right = null;
            Children = null;
            Token = null;
        }
    }

    public struct GlobalDefinitions
    {
        public List<ASTNode?> Functions;
        public List<ASTNode?> Globals;

        public bool FunctionNameExists(Token? functionNameToken)
        {
            if (functionNameToken == null) {
                return false;
            }
            if (Functions == null) {
                Fmt.PrintlnRed("\n\nERROR: functions list is null\n\n");
                return false;
            }
            int functionCount = Functions.Count;
            if (functionCount == 0) {
                return false;
            }
            for (int i = 0; i < functionCount; i++)
            {
                ASTNode? functionNode = Functions[i];
                if (functionNode == null) {
                    Fmt.PrintlnRed($"\n\nERROR: functionNode is null, index: {i}\n\n");
                    continue;
                }
                Token? savedFunctionNameToken = functionNode.Token;
                if (savedFunctionNameToken == null) {
                    Fmt.PrintlnRed($"\n\nERROR: functionNode token is null, index: {i}\n\n");
                    continue;
                }
                if (savedFunctionNameToken.Value.Text == functionNameToken.Value.Text)
                {
                    return true;
                }
            }

            return false;
        }
        public bool GlobalNameExists(Token? varNameToken)
        {
            if (varNameToken == null) {
                return false;
            }
            int globalCount = Globals.Count;
            if (globalCount == 0) {
                return false;
            }
            for (int i = 0; i < globalCount; i++)
            {
                ASTNode? declarationNode = Globals[i];
                if (declarationNode == null)
                {
                    Fmt.PrintlnRed($"\n\nERROR: declarationNode is null, index: {i}\n\n");
                    continue;
                }
                Token? globalVarToken = declarationNode.Token;
                if (globalVarToken == null)
                {
                    Fmt.PrintlnRed($"\n\nERROR: declarationNode token is null, index: {i}\n\n");
                    continue;
                }
                if (globalVarToken.Value.Text == varNameToken.Value.Text)
                {
                    return true;
                }
            }

            return false;
        }
    }

    public struct LocalDefinitions
    {
        public List<ASTNode?> LocalVariables;

        public void PrintDefinitions()
        {
            Fmt.PrintlnColor("Printing definitions:", ConsoleColor.Cyan);
            int definitionCount = LocalVariables.Count;
            if (definitionCount == 0) {
                Fmt.Println($"\tNone\n");
            }
            for (int i = 0; i < definitionCount; i++)
            {
                Debug.Assert(LocalVariables[i] != null);
                Debug.Assert(LocalVariables[i].Token != null);
                Fmt.Println($"Var: {LocalVariables[i].Token.Value.Text}");
            }
            Fmt.Println(' ');
        }

        enum DefinitionLoopResult
        {
            Continue,
            Break,
            ReturnTrue,
        }

        DefinitionLoopResult CheckExistsInner(int i, Token? token)
        {
            ASTNode? declarationNode = LocalVariables[i];
            if (declarationNode == null) {
                Fmt.PrintlnRed($"\n\nERROR: declaration node null!!!, index {i}\n\n");
                return DefinitionLoopResult.Continue;
            }

            Token? nameToken = declarationNode.Token;
            if (nameToken == null) {
                Fmt.PrintlnRed($"\n\nERROR: name token is null!!!, index {i}\n\n");
                return DefinitionLoopResult.Continue;
            }

            if (nameToken.Value.Text == token.Value.Text) {
                return DefinitionLoopResult.ReturnTrue;
            }
            return DefinitionLoopResult.Continue;
        }

        public bool CheckExistsLocally(Token? token)
        {
            int definitionCount = LocalVariables.Count;

            if (token == null) {
                return false;
            }

            if (LocalVariables == null) {
                Fmt.PrintlnRed("\n\nERROR: localvariables list is null!!!\n\n");
                return false;
            }

            if (definitionCount == 0) {
                return false;
            }
            for (int i = 0; i < definitionCount; i++) {

                if (CheckExistsInner(i, token) == DefinitionLoopResult.ReturnTrue) {
                    return true;
                }
            }
            return false;
        }
    }

    public readonly struct Token
    {
        public readonly string Text;
        public readonly TokenType Type;
        public readonly int LineNumber;
        public readonly int CharNumber;

        public Token(string _value, TokenType _type, int _lineNumber, int _charNumber)
        {
            Text = _value;
            Type = _type;
            LineNumber = _lineNumber;
            CharNumber = _charNumber;
        }

        public void PrintValues()
        {
            Fmt.Println($"Token values:\n text: '{Text}'\n type: {Type}\n line no: {LineNumber}\n");
        }
        public bool IsType(TokenType type)
        {
            return Type == type;
        }
    }

    public class ConvertData
    {
        public StringBuilder GeneratedCode = new();
        public ConvertResult ConvertResult;
        public string ErrorDetail = "";
        public CodeFormat CodeFormat;
        public List<string> ErrorTrace = new List<string>();
        public Token? ErrorToken;
        public NodeType lastNodeType = NodeType.Invalid;
        public List<string> MethodVarNames = new List<string>();
        public bool IsMethod = false;
        public string StructName = "";

        public bool IsError()
        {
            if (ConvertResult != ConvertResult.Ok) {
                return true;
            }
            return false;
        }
        public void AddTrace(string functionName)
        {
            ErrorTrace.Add(functionName);
        }
        public void SetError(ConvertResult result, string detail, Token? errorToken, string functionName)
        {
            if (ConvertResult != ConvertResult.Ok) {
                return;
            }
            ConvertResult = result;
            ErrorToken = errorToken;
            ErrorDetail = detail;
            ErrorTrace.Add(functionName);
        }
        public void EndOfFileError(Token? lastToken, string functionName)
        {
            ConvertResult = ConvertResult.Unexpected_End_Of_File;
            ErrorTrace.Add(functionName);
            ErrorToken = lastToken;
        }
        public void MissingTypeError(Token? lastToken, string detail, string functionName)
        {
            ConvertResult = ConvertResult.Missing_Expected_Type;
            ErrorDetail = detail;
            ErrorTrace.Add(functionName);
            ErrorToken = lastToken;
        }

        const int FIRST = 0, SECOND = 1, THIRD = 2, FOURTH = 3, FIFTH = 4;

        public void NewLine(int nestCount) {
            AppendChar('\n');
        }
        public void NewLineWithTabs(int nestCount)
        {
            AppendChar('\n');
            if (nestCount == 0) {
                return;
            }
            for (int i = 0; i < nestCount; i++)
            {
                AppendChar('\t');
            }
        }
        public void AddTabs(int nestCount)
        {
            for (int i = 0; i < nestCount; i++) {
                AppendChar('\t');
            }
        }

        public void UnexpectedTypeError(Token? token, string detail, string functionName)
        {
            ConvertResult = ConvertResult.Unexpected_Type;
            ErrorDetail = detail;
            ErrorTrace.Add(functionName);
            ErrorToken = token;
        }
        public void AppendToken(Token token)
        {
            GeneratedCode.Append(token.Text);
        }
        public void AppendString(string input)
        {
            GeneratedCode.Append(input);
        }
        public void AppendChar(char input)
        {
            GeneratedCode.Append(input);
        }

        public void NoTokenError(Token? errorToken, string detail, string functionName)
        {
            ConvertResult = ConvertResult.No_Token_In_Node;
            ErrorToken = errorToken;
            ErrorDetail = detail;
            ErrorTrace.Add(functionName);
        }
    }

    public class Variable
    {
        public List<Token> TypeList = new List<Token>(); //types can be many tokens *[]int = 3 tokens
        public List<Token> NameToken = new List<Token>();

        public void SetToDefaults()
        {
            if (TypeList == null) {
                TypeList = new List<Token>();
            } else {
                TypeList.Clear();
            }

            if (NameToken == null) {
                NameToken = new List<Token>();
            } else {
                NameToken.Clear();
            }
        }
        public void PrintTypeList()
        {
            if (TypeList == null)
            {
                Fmt.PrintColor("Typelist null ", ConsoleColor.Red);
                return;
            }
            if (TypeList.Count == 0)
            {
                Fmt.PrintColor("Typelist empty ", ConsoleColor.Red);
                return;
            }
            for (int i = 0; i < TypeList.Count; i++)
            {
                Fmt.PrintColor($" '{TypeList[i].Text}', ", ConsoleColor.DarkCyan);
            }
        }
        public void PrintVarName()
        {
            if (NameToken == null)
            {
                Fmt.PrintlnColor(" null", ConsoleColor.Red);
                return;
            }
            if (NameToken.Count == 0)
            {
                Fmt.PrintlnColor(" zero", ConsoleColor.Red);
                return;
            }
            Fmt.PrintColor("varname: ", ConsoleColor.DarkGray);
            Fmt.PrintColor($"{NameToken[0].Text} ", ConsoleColor.Cyan);
        }
        public void Print()
        {
            PrintTypeList();
            PrintVarName();

        }
        public void Println()
        {
            PrintTypeList();
            PrintVarName();
            Fmt.Println();
        }
        public string ConvertToString()
        {
            StringBuilder sb = new StringBuilder();
            if (TypeList.Count == 0) {
                sb.Append("INVALID_TYPE ");
            }
            foreach (Token token in TypeList) {
                sb.Append(token.Text);
            }
            sb.Append(' ');
            if (NameToken == null) {
                sb.Append("NULL");
            } else {
                sb.Append(NameToken[0].Text);
            }

            return sb.ToString();
        }
    
        public void MoveTypeToName()
        {
            if (TypeList == null) {
                TypeList = new List<Token>();
            }
            if (NameToken == null) {
                NameToken = new List<Token>();
            }
            if (TypeList.Count == 0) {
                return;
            }
            NameToken.Add(TypeList[0]);
            TypeList.Clear();
        }
    }

}
