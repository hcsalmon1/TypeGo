


using System.Text;

namespace TypeGo
{
    public static class ConvertBlock
    {

        public static void ProcessBlockData(ConvertData convertData, BlockData blockData, int nestCount)
        {
            NodeType nodeType = blockData.NodeType;

            bool wasIfOrElseIf =
                convertData.lastNodeType == NodeType.If_Statement ||
                convertData.lastNodeType == NodeType.Else_Statement;

            if (wasIfOrElseIf)
            {
                if (nodeType != NodeType.Else_Statement) {
                    //convertData.NewLineWithTabs(nestCount);
                }
                convertData.lastNodeType = NodeType.Invalid;
            }

            switch (nodeType)
            {
                case NodeType.Invalid:
                    convertData.ConvertResult = ConvertResult.Invalid_Node_Type;
                    convertData.ErrorDetail = "invalid node type in block data";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;

                case NodeType.Single_Declaration_With_Value:
                    ProcessSingleDeclarationWithValue(convertData, blockData, nestCount);
                    break;

                case NodeType.Single_Declaration_No_Value:
                    ProcessSingleDeclarationNoValue(convertData, blockData, nestCount);
                    break;

                case NodeType.Multiple_Declarations_No_Value:
                    ProcessMultipleDeclarationNoValue(convertData, blockData, nestCount);
                    break;

                case NodeType.Multiple_Declarations_With_Value:
                case NodeType.Multiple_Declarations_Same_Type_No_Value:
                    convertData.ConvertResult = ConvertResult.Unsupported_Type;
                    convertData.ErrorDetail = $"unsupported node type: {nodeType}";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;

                case NodeType.Multiple_Declarations_Same_Type_With_Value:
                case NodeType.Multiple_Declarations_One_Type_One_Set_Value:
                    ProcessMultipleDeclarationWithSetValue(convertData, blockData, nestCount);
                    break;

                case NodeType.If_Statement:
                    ConvertIf.PrintIfStatement(convertData, blockData, nestCount);
                    break;
                case NodeType.Else_Statement:
                    ConvertIf.PrintElseStatement(convertData, blockData, nestCount);
                    break;
                case NodeType.For_Loop:
                    ConvertFor.PrintFor(convertData, blockData, nestCount);
                    break;

                case NodeType.Multi_Line_Import:
                case NodeType.Single_Import:
                case NodeType.Single_Import_With_Alias:
                case NodeType.Package:
                    PrintTokens(convertData, blockData, nestCount, newLine: true);
                    break;
                case NodeType.Other:
                    PrintTokens(convertData, blockData, nestCount, newLine: false);
                    break;
                case NodeType.Struct_Declaration:
                    PrintStruct.ProcessStruct(convertData, blockData, nestCount);
                    break;
                case NodeType.Enum_Declaration:
                    PrintEnum.ProcessEnum(convertData, blockData, nestCount);
                    break;

                case NodeType.Enum_Struct_Declaration:
                    PrintEnum.ProcessEnumstruct(convertData, blockData, nestCount);
                    break;

                case NodeType.Append:
                    PrintAppend.Print(convertData, blockData, nestCount);
                    break;

                case NodeType.Constant_Global_Variable:
                    convertData.AppendString("const ");
                    PrintTokens(convertData, blockData, nestCount, newLine: true);
                    break;

                case NodeType.Constant_Global_Variable_With_Type:
                    PrintConstant(convertData, blockData, nestCount);
                    break;

                case NodeType.Comment:
                    PrintTokens(convertData, blockData, nestCount, newLine: true);
                    break;

                case NodeType.Channel_Declaration:
                    ProcessChannelDeclarationNoValue(convertData, blockData, nestCount);
                    break;

                case NodeType.Channel_Declaration_With_Value:
                    ProcessChannelDeclarationWithValue(convertData, blockData, nestCount);
                    break;

                case NodeType.Interface_Declaration:
                    ProcessInterfaceDeclaration(convertData, blockData, nestCount);
                    break;

                case NodeType.NewLine:
                    convertData.NewLineWithTabs(nestCount);
                    break;

                case NodeType.Err_Return:
                    ConvertErrReturn.Process(convertData, blockData, nestCount);
                    break;

                default:
                    convertData.ConvertResult = ConvertResult.Unexpected_Type;
                    convertData.ErrorDetail = $"unexpected node type '{nodeType}' in block data";
                    convertData.ErrorToken = blockData.StartingToken;
                    break;
            }
        }

        private static void PrintConstant(ConvertData convertData, BlockData blockData, int nestCount)
        {
            convertData.AppendString("const ");
            if (blockData.Variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "no variable tokens in constant declaration";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            Variable variable = blockData.Variables[0];
            if (variable.NameToken == null) {
                convertData.ConvertResult = ConvertResult.Null_Token;
                convertData.ErrorDetail = "null name token in constant declaration";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            convertData.AppendToken(variable.NameToken[0]);
            convertData.AppendChar(' ');
            if (variable.TypeList.Count == 0) {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "no variable type in constant declaration";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            convertData.AppendString(variable.TypeList[0].Text);
            PrintTokens(convertData, blockData, nestCount, newLine: true);
        }

        private static void ProcessSingleDeclarationNoValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            StringBuilder sb = new();

            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "no variables in single declaration";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int varIndex = 0; varIndex < variables.Count; varIndex++)
            {

                List<Token> varTypeTokenList = variables[varIndex].TypeList;

                if (varTypeTokenList == null) {
                    convertData.UnexpectedTypeError(blockData.StartingToken, "Type is invalid in single declaration", "ProcessSingleDeclarationNoValue");
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null) {
                    convertData.UnexpectedTypeError(blockData.StartingToken, "Type is invalid in single declaration", "ProcessSingleDeclarationNoValue");
                    return;
                }

                List<Token> varNameList = variables[varIndex].NameToken;

                for (int nameTokenIndex = 0; nameTokenIndex < varNameList.Count; nameTokenIndex++)
                {
                    Token? nameToken = variables[varIndex].NameToken[nameTokenIndex];
                    if (nameToken == null) {
                        convertData.UnexpectedTypeError(blockData.StartingToken, "Variable name is invalid in single declaration", "ProcessSingleDeclarationNoValue");
                        return;
                    }

                    if (nameTokenIndex != 0) {
                        sb.Append(',');
                        sb.Append(' ');
                    }
                    sb.Append(nameToken.Value.Text);
                }

                convertData.AppendString($"var {sb.ToString()} {varTypeAsText}");
                convertData.NewLineWithTabs(nestCount);
            }
        }

        private static void ProcessChannelDeclarationNoValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "no variables in single declaration";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++)
            {

                List<Token> varTypeTokenList = variables[i].TypeList;

                if (varTypeTokenList == null)
                {
                    convertData.ConvertResult = ConvertResult.Unexpected_Type;
                    convertData.ErrorDetail = "Type is invalid in single declaration";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null)
                {
                    convertData.ConvertResult = ConvertResult.Unexpected_Type;
                    convertData.ErrorDetail = "Type is invalid in single declaration";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                Token? nameToken = variables[i].NameToken[0];
                if (nameToken == null)
                {
                    convertData.ConvertResult = ConvertResult.Unexpected_Type;
                    convertData.ErrorDetail = "Variable name is invalid in single declaration";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                convertData.AppendString($"var {nameToken.Value.Text} chan {varTypeAsText}");
                convertData.NewLineWithTabs(nestCount);
            }
        }

        private static void ProcessChannelDeclarationWithValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Token> tokens = blockData.Tokens;
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No variables in single declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (tokens.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "No tokens in single declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++)
            {

                List<Token> varTypeTokenList = variables[i].TypeList;

                if (varTypeTokenList == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "invalid var type in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "invalid var type text in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                Token? nameToken = variables[i].NameToken[0];
                if (nameToken == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "No variable name in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                convertData.AppendString($"var {nameToken.Value.Text} chan {varTypeAsText} ");
                PrintTokensDeclaration(convertData, blockData, nestCount, newLine: false, varTypeAsText);
            }
        }

        private static void ProcessSingleDeclarationWithValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Token> tokens = blockData.Tokens;
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No variables in single declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (tokens.Count == 0) {
                convertData.ConvertResult = ConvertResult.No_Token_In_Node;
                convertData.ErrorDetail = "No tokens in single declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++) {

                List<Token> varTypeTokenList = variables[i].TypeList;
                 
                if (varTypeTokenList == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "invalid var type in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "invalid var type text in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                Token? nameToken = variables[i].NameToken[0];
                if (nameToken == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "No variable name in single declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                convertData.AppendString($"var {nameToken.Value.Text} {varTypeAsText} ");
                PrintTokensDeclaration(convertData, blockData, nestCount, newLine: false, varTypeAsText);
            }
        }

        private static void ProcessMultipleDeclarationNoValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Token> tokens = blockData.Tokens;
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No variables in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (tokens.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No Tokens in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++)
            {

                List<Token> varTypeTokenList = variables[i].TypeList;

                if (varTypeTokenList == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = " var type is invalid in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "var type text is invalid in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                List<Token> varNameTokenList = variables[i].NameToken;

                convertData.AppendChar('v');
                convertData.AppendChar('a');
                convertData.AppendChar('r');
                convertData.AppendChar(' ');

                for (int varNameIndex = 0; varNameIndex < varNameTokenList.Count; varNameIndex++)
                {
                    Token? nameToken = varNameTokenList[varNameIndex];
                    if (nameToken == null)
                    {
                        convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                        convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                        convertData.ErrorToken = blockData.StartingToken;
                        return;
                    }

                    if (varNameIndex != 0) {
                        convertData.AppendChar(',');
                        convertData.AppendChar(' ');
                    }
                    convertData.AppendString($"{nameToken.Value.Text}");
                }

                convertData.AppendString($" {varTypeAsText}\n\t");
            }


            convertData.AppendChar(' ');
            PrintTokensDeclaration(convertData, blockData, nestCount, newLine: false, "");
        }


        private static void ProcessMultipleDeclarationWithValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Token> tokens = blockData.Tokens;
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No variables in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (tokens.Count == 0) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No Tokens in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++)
            {

                List<Token> varTypeTokenList = variables[i].TypeList;

                if (varTypeTokenList == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = " var type is invalid in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);
                if (varTypeAsText == null) {
                    //convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    //convertData.ErrorDetail = "var type text is invalid in multiple declaration with value";
                    //convertData.ErrorToken = blockData.StartingToken;
                    continue;
                }

                Token? nameToken = variables[i].NameToken[0];
                if (nameToken == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                convertData.AppendString($"var {nameToken.Value.Text} {varTypeAsText}\n\t");
            }

            for (int i = 0; i < variables.Count; i++)
            {
                if (variables[i].NameToken.Count == 0)
                {
                    continue;
                }
                Token? nameToken = variables[i].NameToken[0];
                if (nameToken == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                if (i != 0) {
                    convertData.AppendChar(',');
                    convertData.AppendChar(' ');
                }
                convertData.AppendString($"{nameToken.Value.Text}");
            }
            convertData.AppendChar(' ');
            PrintTokensDeclaration(convertData, blockData, nestCount, newLine: false, "");
        }

        private static void ProcessMultipleDeclarationWithSetValue(ConvertData convertData, BlockData blockData, int nestCount)
        {
            List<Token> tokens = blockData.Tokens;
            List<Variable> variables = blockData.Variables;
            if (variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No variables in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (tokens.Count == 0) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "No Tokens in multiple declaration with value";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            for (int i = 0; i < variables.Count; i++)
            {

                List<Token> varTypeTokenList = variables[i].TypeList;

                if (varTypeTokenList == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = " var type is invalid in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);

                if (variables[i].NameToken.Count == 0) {
                    continue;
                }

                Token? nameToken = variables[i].NameToken[0]; 
                if (nameToken == null) {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                if (varTypeAsText == null) {
                    continue;
                }

                convertData.AppendString($"var {nameToken.Value.Text} {varTypeAsText}\n\t");
            }

            for (int i = 0; i < variables.Count; i++)
            {
                Token? nameToken = null;

                if (variables[i].NameToken.Count == 0) {

                    List<Token> varTypeTokenList = variables[i].TypeList;
                    string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(varTypeTokenList);

                    if (varTypeAsText == null) {
                        convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                        convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                        convertData.ErrorToken = blockData.StartingToken;
                        return;
                    }

                } else {

                    nameToken = variables[i].NameToken[0];
                }


                if (nameToken == null)
                {
                    convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                    convertData.ErrorDetail = "name Token is null in multiple declaration with value";
                    convertData.ErrorToken = blockData.StartingToken;
                    return;
                }

                if (i != 0)
                {
                    convertData.AppendChar(',');
                    convertData.AppendChar(' ');
                }
                convertData.AppendString($"{nameToken.Value.Text}");
            }
            convertData.AppendChar(' ');
            PrintTokensDeclaration(convertData, blockData, nestCount, newLine: false, "");
        }


        static void WriteVarTypeText(ConvertData convertData, string varTypeAsText)
        {
            char firstChar = varTypeAsText[0];

            if (firstChar == '*') {

                char[] charArray = varTypeAsText.ToCharArray();
                charArray[0] = '&';
                convertData.AppendString(new string(charArray));
                return;
            }
            convertData.AppendString(varTypeAsText);
        }

        private static void PrintTokensDeclaration(ConvertData convertData, BlockData blockData, int nestCount, bool newLine, string varTypeAsText)
        {
            if (blockData.Tokens.Count == 0) {
                convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData", "PrintTokens");
                return;
            }
            TokenType lastType = TokenType.NA;
            bool inMake = false;
            bool addedSpace = false;

            for (int i = 0; i < blockData.Tokens.Count; i++) {

                Token token = blockData.Tokens[i];

                if (token.Type == TokenType.NewLine) {
                    convertData.NewLineWithTabs(nestCount);
                    lastType = token.Type;
                    continue;
                }
                if (i == 2 || i == 1) {
                    if (token.Type == TokenType.LeftBrace) {
                        WriteVarTypeText(convertData, varTypeAsText);
                    }
                    if (lastType == TokenType.Make) {
                        if (token.Type == TokenType.LeftParenthesis) {
                            convertData.AppendChar('(');
                            convertData.AppendString(varTypeAsText);
                            inMake = true;
                            continue;
                        }
                    }
                }
                if (inMake == true) {

                    if (token.Type != TokenType.RightParenthesis) {
                        convertData.AppendChar(',');
                        convertData.AppendChar(' ');
                    }
                    inMake = false;
                }

                AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace);
                convertData.AppendString(token.Text);
                addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i);
                lastType = token.Type;
            }

            if (lastType != TokenType.NewLine) {
                convertData.NewLineWithTabs(nestCount);
            }
            if (newLine == true) {
                convertData.NewLineWithTabs(nestCount);
            }
        }

        static void ProcessInterfaceDeclaration(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Tokens.Count == 0)
            {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "No tokens in interface declaration";
                return;
            }

            //first token is the name of the interface
            Token? interfaceNameToken = blockData.Tokens[0];

            if (interfaceNameToken == null)
            {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "interface name token is null in interface declaration";
                return;
            }

            //block has method list and that's all, 0 count is blank interface
            convertData.AppendString($"type {interfaceNameToken.Value.Text} interface {{\n\t");

            if (blockData.Block == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Block is null in interface declaration";
                return;
            }

            List<Function> methodList = blockData.Block.MethodList;

            if (methodList.Count == 0) {
                EndInterface(convertData);
                return;
            }

            for (int i = 0; i < methodList.Count; i++)
            {
                Function function = methodList[i];

                //public string? ReturnType = "";
                //public string Name = "";
                //public List<Variable> Parameters = new List<Variable>();
                //public CodeBlock? InnerBlock = null;

                convertData.GeneratedCode.Append($"{function.Name}(");

                PrintParameters(convertData, function);
                convertData.GeneratedCode.Append(')');
                convertData.GeneratedCode.Append(' ');
                if (function.ReturnType != null) {
                    convertData.GeneratedCode.Append(function.ReturnType);
                    convertData.GeneratedCode.Append(' ');
                }
                convertData.GeneratedCode.Append('\n');
                convertData.GeneratedCode.Append('\t');
            }

            EndInterface(convertData);

            static void EndInterface(ConvertData convertData)
            {
                convertData.AppendChar('\r');
                convertData.AppendChar('}');
                convertData.AppendChar('\n');
            }

            static void PrintParameters(ConvertData convertData, Function function)
            {
                List<Variable> parameters = function.Parameters;
                if (parameters.Count == 0)
                {
                    return;
                }

                for (int parameterIndex = 0; parameterIndex < parameters.Count; parameterIndex++)
                {

                    Variable parameter = parameters[parameterIndex];
                    if (parameter.NameToken == null)
                    {
                        convertData.ConvertResult = ConvertResult.Internal_Error;
                        convertData.ErrorDetail = "name token is null in PrintParameters";
                        return;
                    }
                    string? typeAsText = TokenUtils.JoinTextInListOfTokens(parameter.TypeList);
                    if (typeAsText == null)
                    {
                        convertData.ConvertResult = ConvertResult.Internal_Error;
                        convertData.ErrorDetail = "var type text is null in PrintParameters";
                        return;
                    }
                    if (parameterIndex != 0)
                    {
                        convertData.GeneratedCode.Append(',');
                        convertData.GeneratedCode.Append(' ');
                    }
                    convertData.GeneratedCode.Append($"{parameter.NameToken[0].Text} {typeAsText}");
                }
            }
        }

        static void AddMethodPrefix(ConvertData convertData, string varName)
        {
            if (convertData.IsMethod == false) {
                return;
            }

            if (convertData.MethodVarNames.Count == 0) {
                return;
            }
            if (IsMethodVar(convertData, varName) == false) {
                return;
            }

            char firstLetter = convertData.StructName[0];
            firstLetter = char.ToLower(firstLetter);
            convertData.AppendChar(firstLetter);
            convertData.AppendChar('.');
            return;
        }

        static bool IsMethodVar(ConvertData convertData, string varName)
        {
            for (int i = 0; i < convertData.MethodVarNames.Count; i++)
            {
                string name = convertData.MethodVarNames[i];
                if (name == varName) {
                    return true;
                }
            }
            return false;
        }

        public static void PrintTokens(ConvertData convertData, BlockData blockData, int nestCount, bool newLine)
        {
            if (blockData.Tokens.Count == 0) {
                convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData", "PrintTokens");
                return;
            }
            TokenType lastType = TokenType.NA;
            bool addedSpace = false;

            for (int i = 0; i < blockData.Tokens.Count; i++)
            {

                Token token = blockData.Tokens[i];

                if (token.Type == TokenType.NewLine)
                {
                    convertData.NewLineWithTabs(nestCount);
                    lastType = token.Type;
                    continue;
                }
                AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace);
                HandleToken(convertData, token);
                addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i);
                lastType = token.Type;
            }

            if (lastType != TokenType.NewLine) {
                convertData.NewLineWithTabs(nestCount);
            }
            if (newLine == true) {
                convertData.NewLineWithTabs(nestCount);
            }

            static void HandleToken(ConvertData convertData, Token token)
            {
                if (token.Type == TokenType.Semicolon) {

                    int codeLength = convertData.GeneratedCode.Length;
                    char lastChar = convertData.GeneratedCode[codeLength - 1];
                    if (lastChar == ' ') {
                        convertData.GeneratedCode[codeLength - 1] = ';';
                        return;
                    }
                }
                convertData.AppendString(token.Text);
            }
        }

        public static void PrintTokensNoNL(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Tokens.Count == 0)
            {
                convertData.NoTokenError(blockData.StartingToken, "no tokens in blockData", "PrintTokens");
                return;
            }
            TokenType lastType = TokenType.NA;
            bool addedSpace = false;

            for (int i = 0; i < blockData.Tokens.Count; i++)
            {

                Token token = blockData.Tokens[i];

                if (token.Type == TokenType.NewLine)
                {
                    convertData.NewLineWithTabs(nestCount);
                    lastType = token.Type;
                    continue;
                }
                AddSpaceBefore(convertData, token.Type, lastType, i, addedSpace);
                convertData.AppendString(token.Text);
                addedSpace = AddSpaceAfter(convertData, token.Type, lastType, i);
                lastType = token.Type;
            }
        }

        static void AddSpaceBefore(ConvertData convertData, TokenType thisType, TokenType lastType, int tokenIndex, bool addedSpace)
        {
            int codeLength = convertData.GeneratedCode.Length;

            if (codeLength > 0)
            {
                char lastCharAdded = convertData.GeneratedCode[codeLength - 1];

                if (lastCharAdded == ' ')
                {
                    return;
                }
            }

            if (addedSpace == true) {
                return;
            }

            bool addSpace = false;


            switch (thisType)
            {
                case TokenType.NotEquals:
                case TokenType.And:
                case TokenType.AndAnd:
                case TokenType.Or:
                case TokenType.OrOr:
                case TokenType.PlusEquals:
                case TokenType.MinusEquals:
                case TokenType.MultiplyEquals:
                case TokenType.DivideEquals:
                case TokenType.GreaterThan:
                case TokenType.LessThan:
                case TokenType.EqualsEquals:
                case TokenType.GreaterThanEquals:
                case TokenType.LessThanEquals:
                case TokenType.Modulus:
                case TokenType.ModulusEquals:
                case TokenType.ColonEquals:
                case TokenType.Equals:
                case TokenType.LeftBrace:
                case TokenType.Comment:
                    addSpace = true;
                    break;

                case TokenType.Plus:

                    bool shouldHaveSpaceBeforePlus =
                        lastType == TokenType.Identifier ||
                        lastType == TokenType.IntegerValue ||
                        lastType == TokenType.StringValue;

                    if (shouldHaveSpaceBeforePlus) {
                        addSpace = true;
                    }
                    break;

                case TokenType.Multiply:
                    if (lastType == TokenType.Identifier) {
                        addSpace = true;
                    }
                    break;


                //case TokenType.Minus:
                case TokenType.Divide:
                    addSpace = true;
                    break;

                default:
                    break;
            }

            if (thisType == TokenType.Minus)
            {

                bool isOperator =
                    lastType == TokenType.Identifier ||
                    lastType == TokenType.IntegerValue ||
                    TokenUtils.IsOperator(lastType);

                if (isOperator)
                {
                    addSpace = true;
                }
            }

            if (thisType == TokenType.Identifier)
            {
                if (lastType == TokenType.Identifier)
                {
                    addSpace = true;
                }
            }

            if (addSpace)
            {
                convertData.AppendChar(' ');
            }
        }

        static bool AddSpaceAfter(ConvertData convertData, TokenType thisType, TokenType lastType, int tokenIndex)
        {

            bool addSpace = false;

            switch (thisType)
            {
                case TokenType.If:
                case TokenType.Else:
                case TokenType.For:
                case TokenType.Switch:
                case TokenType.Struct:
                case TokenType.Bool:
                case TokenType.Int:
                case TokenType.Int16:
                case TokenType.Int32:
                case TokenType.Int64:
                case TokenType.Int8:
                case TokenType.AndAnd:
                case TokenType.Or:
                case TokenType.Return:
                case TokenType.PlusEquals:
                case TokenType.MinusEquals:
                case TokenType.DivideEquals:
                case TokenType.Enum:
                case TokenType.Enumstruct:
                case TokenType.GreaterThanEquals:
                case TokenType.LessThanEquals:
                case TokenType.Goto:
                case TokenType.Equals:
                case TokenType.Divide:
                case TokenType.Fn:
                case TokenType.Package:
                case TokenType.Import:
                case TokenType.Comma:
                case TokenType.Const:
                case TokenType.Semicolon:
                case TokenType.OrOr:
                case TokenType.ColonEquals:
                case TokenType.NotEquals:
                case TokenType.ModulusEquals:
                case TokenType.Modulus:
                case TokenType.EqualsEquals:
                case TokenType.LessThan:
                case TokenType.MultiplyEquals:
                case TokenType.GreaterThan:
                case TokenType.Defer:
                case TokenType.ErrReturn:
                case TokenType.Case:
                case TokenType.Break:
                case TokenType.Continue:
                    addSpace = true;
                    break;

                default:
                    break;
            }

            if (thisType == TokenType.Multiply) {

                addSpace = true;
            }

            if (thisType == TokenType.Minus)
            {
                bool isOperator =
                    lastType == TokenType.Identifier ||
                    lastType == TokenType.IntegerValue;

                if (isOperator)
                {
                    addSpace = true;
                }
            }

            if (thisType == TokenType.Plus || thisType == TokenType.And)
            {
                bool isOperator =
                    lastType == TokenType.Identifier ||
                    lastType == TokenType.IntegerValue;

                if (isOperator)
                {
                    addSpace = true;
                }
            }

            if (addSpace) {
                convertData.AppendChar(' ');
                return true;
            }
            return false;
        }

        public static void ProcessBlock(ConvertData convertData, CodeBlock codeBlock, int nestCount)
        {
            List<BlockData> blockDataList = codeBlock.BlockDataList;
            if (blockDataList.Count == 0) {
                return;
            }
            convertData.NewLineWithTabs(nestCount);

            for (int i = 0; i < blockDataList.Count; i++)
            {
                BlockData blockData = blockDataList[i];
                if (blockData.NodeType == NodeType.NewLine && i == 0) {
                    continue;
                }
                ProcessBlockData(convertData, blockData, nestCount);
            }
        }
    }
}
