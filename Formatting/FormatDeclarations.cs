using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class DeclarationUtils
    {
        static void ConvertToNodetypeWithValue(BlockData blockData)
        {
            switch (blockData.NodeType)
            {
                case NodeType.Single_Declaration_No_Value:
                    blockData.NodeType = NodeType.Single_Declaration_With_Value;
                    break;
                case NodeType.Multiple_Declarations_No_Value:
                    blockData.NodeType = NodeType.Multiple_Declarations_Same_Type_With_Value;
                    break;
                case NodeType.Multiple_Declarations_Same_Type_No_Value:
                    blockData.NodeType = NodeType.Multiple_Declarations_Same_Type_With_Value;
                    break;
                default:
                    break;
            }
        }

        public static WhileLoopAction ProcessAfterNull(FormatData formatData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (TokenUtils.IsVarType(token)) {
                lastType = LastTokenType.Vartype;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was null, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Multiply) {
                lastType = LastTokenType.Pointer;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was null, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.LeftSquareBracket) {
                lastType = LastTokenType.LeftSqBracket;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was null, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Identifier) {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was null, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Map) {
                lastType = LastTokenType.Map;
                tempVariable.TypeList.Add(token);

                return WhileLoopAction.Continue;
            }

            formatData.UnexpectedTypeError(token, $"unexpected type: {token.Type} in declaration", "ProcessAfterNull");
            return WhileLoopAction.Error;

        }
        
        public static WhileLoopAction ProcessAfterVarType(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable, string thisFunctionName)
        {
            if (TokenUtils.IsVarType(token)) {
                formatData.UnexpectedTypeError(token, "unexpected vartype in declaration", thisFunctionName);
                return WhileLoopAction.Return;
            }

            if (token.Type == TokenType.Identifier) {

                tempVariable.NameToken = token;
                blockData.Variables.Add(new Variable {
                    NameToken = tempVariable.NameToken,
                    TypeList = TokenUtils.CopyTokenList(tempVariable.TypeList),
                });
                lastType = LastTokenType.Identifier;
                tempVariable.SetToDefaults();
                formatData.Increment();
                return WhileLoopAction.Continue;
            }

            if (token.Type == TokenType.RightSquareBracket) {
                lastType = LastTokenType.RightSqBracket;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }

            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after variable", "ProcessAfterVarType");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterIdentifier(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (token.Type == TokenType.Semicolon) {
                return WhileLoopAction.Break;
            }

            if (token.Type == TokenType.Equals) {
                ConvertToNodetypeWithValue(blockData);
                return WhileLoopAction.Break;
            }

            if (token.Type == TokenType.Comma) {
                blockData.NodeType = NodeType.Multiple_Declarations_No_Value;
                lastType = LastTokenType.Comma;
                return WhileLoopAction.Continue;
            }

            if (token.Type == TokenType.NewLine) {
                return WhileLoopAction.Break;
            }

            if (token.Type == TokenType.Identifier) {
                tempVariable.NameToken = token;
                blockData.Variables.Add(new Variable
                {
                    NameToken = tempVariable.NameToken,
                    TypeList = TokenUtils.CopyTokenList(tempVariable.TypeList),
                });
                lastType = LastTokenType.Identifier;
                tempVariable.SetToDefaults();
                formatData.Increment();
                return WhileLoopAction.Continue;
            }

            if (token.Type == TokenType.FullStop) {
                lastType = LastTokenType.FullStop;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was null, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }

            if (token.Type == TokenType.RightSquareBracket) {
                lastType = LastTokenType.RightSqBracket;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }

            formatData.UnexpectedTypeError(token, "unexpected typed after identifier", "ProcessAfterIdentifier");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterPointer(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (TokenUtils.IsVarType(token)) {
                lastType = LastTokenType.Vartype;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was pointer, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.LeftSquareBracket)
            {
                lastType = LastTokenType.LeftSqBracket;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was pointer, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Identifier) {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was pointer, add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }

            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after '*'", "ProcessAfterPointer");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterLeftSqBracket(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (token.Type == TokenType.RightSquareBracket) {
                lastType = LastTokenType.RightSqBracket;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.IntegerValue) {
                lastType = LastTokenType.IntegerValue;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (TokenUtils.IsVarType(token)) {
                lastType = LastTokenType.Vartype;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Identifier) {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }

            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after '['", "ProcessAfterLeftSqBracket");
            return WhileLoopAction.Error;
        }
        
        public static WhileLoopAction ProcessAfterRightSqBracket(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (TokenUtils.IsVarType(token)) {
                lastType = LastTokenType.Vartype;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.LeftSquareBracket) {
                lastType = LastTokenType.LeftSqBracket;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Identifier) {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Multiply) {
                lastType = LastTokenType.Pointer;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Map) {
                lastType = LastTokenType.Map;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after ']'", "ProcessAfterRightSqBracket");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterComma(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (TokenUtils.IsVarType(token)) {
                lastType = LastTokenType.Vartype;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was ',', add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            if (token.Type == TokenType.Identifier) {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was ',', add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after ','", "ProcessAfterComma");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterFullStop(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (token.Type == TokenType.Identifier)
            {
                lastType = LastTokenType.Identifier;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was ',', add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after '.'", "ProcessAfterFullStop");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterIntegerValue(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (token.Type == TokenType.RightSquareBracket)
            {
                lastType = LastTokenType.RightSqBracket;
                tempVariable.TypeList.Add(token);
                //Fmt.PrintlnColor($"\tlast type was ',', add '{token.Text}' to typeList", ConsoleColor.Cyan);
                return WhileLoopAction.Continue;
            }
            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after integer value", "ProcessAfterIntegerValue");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessAfterMap(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable)
        {
            if (token.Type == TokenType.LeftSquareBracket)
            {
                lastType = LastTokenType.LeftSqBracket;
                tempVariable.TypeList.Add(token);
                return WhileLoopAction.Continue;
            }
            formatData.UnexpectedTypeError(token, $"unsupported type: {token.Type} after 'map'", "ProcessToken");
            return WhileLoopAction.Error;
        }

        public static WhileLoopAction ProcessToken(FormatData formatData, BlockData blockData, Token token, ref LastTokenType lastType, ref Variable tempVariable, string thisFunctionName)
        {
            switch (lastType)
            {
                case LastTokenType.Null:
                    return DeclarationUtils.ProcessAfterNull(formatData, token, ref lastType, ref tempVariable);
                case LastTokenType.Identifier:
                    return DeclarationUtils.ProcessAfterIdentifier(formatData, blockData, token, ref lastType, ref tempVariable);
                case LastTokenType.Vartype:
                    return DeclarationUtils.ProcessAfterVarType(formatData, blockData, token, ref lastType, ref tempVariable, thisFunctionName);
                case LastTokenType.Pointer:
                    return DeclarationUtils.ProcessAfterPointer(formatData, blockData, token, ref lastType, ref tempVariable);

                case LastTokenType.LeftSqBracket:
                    return DeclarationUtils.ProcessAfterLeftSqBracket(formatData, blockData, token, ref lastType, ref tempVariable);
                case LastTokenType.RightSqBracket:
                    return DeclarationUtils.ProcessAfterRightSqBracket(formatData, blockData, token, ref lastType, ref tempVariable);
                case LastTokenType.Comma:
                    return DeclarationUtils.ProcessAfterComma(formatData, blockData, token, ref lastType, ref tempVariable);
                case LastTokenType.IntegerValue:
                    return DeclarationUtils.ProcessAfterIntegerValue(formatData, blockData, token, ref lastType, ref tempVariable);
                case LastTokenType.FullStop:
                    return DeclarationUtils.ProcessAfterFullStop(formatData, blockData, token, ref lastType, ref tempVariable);

                case LastTokenType.Map:
                    return DeclarationUtils.ProcessAfterMap(formatData, blockData, token, ref lastType, ref tempVariable);

                default:
                    formatData.UnexpectedTypeError(token, $"unsupported last type: {lastType}, this type: {token.Type}", "ProcessToken");
                    return WhileLoopAction.Error;
            }
        }

        public static void WriteTokens(FormatData formatData, BlockData blockData)
        {
            switch (blockData.NodeType)
            {
                case NodeType.Invalid:
                    break;
                case NodeType.Single_Declaration_With_Value:

                    FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
                    break;
                case NodeType.Single_Declaration_No_Value:
                    FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
                    break;

                case NodeType.Multiple_Declarations_No_Value:

                    formatData.UnsupportedFeatureError(blockData.StartingToken, "Multiple declarations with no value, not supported yet", "WriteTokens");
                    return;

                case NodeType.Multiple_Declarations_With_Value:
                    formatData.UnsupportedFeatureError(blockData.StartingToken, "Multiple declarations with value, not supported yet", "WriteTokens");
                    return;

                case NodeType.Multiple_Declarations_Same_Type_No_Value:
                    formatData.UnsupportedFeatureError(blockData.StartingToken, "Multiple same type declarations with no value, not supported yet", "WriteTokens");
                    return;

                case NodeType.Multiple_Declarations_Same_Type_With_Value:
                    FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
                    break;

                default:
                    break;
            }
        }

    }



    public static class FormatDeclarations
    {

        public static BlockData? DeclarationLoop(FormatData formatData, Token firstToken) {

            const string thisFunctionName = "DeclarationVarFirst";

            //Fmt.PrintlnColor("\t__DeclarationVarFirst__", ConsoleColor.Yellow);

            int whileCount = 0;

            BlockData blockData = new BlockData {
                Block = null,
                NodeType = NodeType.Single_Declaration_No_Value,
                StartingToken = firstToken,
                Tokens = new List<Token>(),
                Variables = new List<Variable>(),
            };
            Variable tempVariable = new Variable {
                NameToken = null,
                TypeList = new List<Token>(),
            };
            LastTokenType lastType = LastTokenType.Null;

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {
                WhileLoopAction result = DeclarationInnerLoop(formatData, thisFunctionName, ref whileCount, blockData, ref tempVariable, ref lastType);
                if (result == WhileLoopAction.Break) {
                    break;
                }
                if (result == WhileLoopAction.Return) {
                    return null;
                }
            }

            DeclarationUtils.WriteTokens(formatData, blockData);
            return blockData;
        }

        private static WhileLoopAction DeclarationInnerLoop(FormatData formatData, string thisFunctionName, ref int whileCount, BlockData blockData, ref Variable tempVariable, ref LastTokenType lastType)
        {
            if (Debugging.InfiniteWhileCheck(ref whileCount, 10000) == true) {
                formatData.Result = FormatResult.Internal_Error;
                formatData.ErrorDetail = "Infinite white loop in DeclarationInnerLoop";
                return WhileLoopAction.Error;
            }

            int previousIndex = formatData.TokenIndex;

            Token? tempToken = formatData.GetToken();
            if (tempToken == null) {
                formatData.EndOfFileError(tempToken, thisFunctionName);
                return WhileLoopAction.Return;
            }
            WhileLoopAction loopResult = DeclarationUtils.ProcessToken(formatData, blockData, tempToken.Value, ref lastType, ref tempVariable, thisFunctionName);
            if (loopResult == WhileLoopAction.Break) {
                return WhileLoopAction.Break;
            }
            if (loopResult == WhileLoopAction.Return) {
                return WhileLoopAction.Return;
            }

            formatData.IncrementIfSame(previousIndex);
            return WhileLoopAction.Continue;
        }

        public static BlockData? ProcessDeclaration(FormatData formatData, Token firstToken)
        {
            //Fmt.PrintlnColor("\n__ProcessDeclaration__", ConsoleColor.Yellow);

            bool isDeclaration =
                TokenUtils.IsVarType(firstToken.Type) ||
                firstToken.Type == TokenType.LeftSquareBracket;

            if (isDeclaration) {
                return DeclarationLoop(formatData, firstToken);
            }

            if (firstToken.Type == TokenType.Multiply) {

                if (FormatUtils.IsPointerDeclaration(formatData)) {
                    return DeclarationLoop(formatData, firstToken);
                }
                return FillNonDeclaration(formatData, firstToken);
            }

            return null;
        }

        static BlockData? FillNonDeclaration(FormatData formatData, Token firstToken)
        {
            BlockData blockData = new BlockData
            {
                Block = null,
                NodeType = NodeType.Other,
                StartingToken = firstToken,
                Tokens = new List<Token>(),
                Variables = new List<Variable>(),
            };
            FormatUtils.LoopTokensUntilLineEnd(formatData, blockData, stopAtSemicolon: true);
            return blockData;
        }


    }
}
