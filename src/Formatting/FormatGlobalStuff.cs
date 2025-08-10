using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Text;

namespace TypeGo
{
    public static class FormatGlobalStuff
    {

        public static class Packages
        {
            public static void ProcessPackage(FormatData formatData, CodeBlock globalBlock, Token packageToken)
            {
                const string thisFunctionName = "ProcessPackage";

                BlockData blockData = new();
                Debug.Assert(blockData.Tokens != null);
                blockData.NodeType = NodeType.Package;
                blockData.StartingToken = packageToken;

                blockData.Tokens.Add(packageToken);

                Token? packageNameToken = formatData.GetNextToken();
                if (packageNameToken == null)
                {
                    formatData.EndOfFileError(packageToken, thisFunctionName);
                    return;
                }
                if (packageNameToken.Value.Type != TokenType.Identifier)
                {
                    formatData.MissingExpectedTypeError(packageNameToken, "Missing identifier in package", thisFunctionName);
                    return;
                }

                Token? newLineToken = formatData.GetNextToken();
                if (newLineToken == null)
                {
                    formatData.EndOfFileError(packageToken, thisFunctionName);
                    return;
                }
                blockData.Tokens.Add(packageNameToken.Value);
                blockData.Tokens.Add(newLineToken.Value);
                formatData.Increment();

                globalBlock.BlockDataList.Add(blockData);
            }
        }

        public static class Imports
        {

            static WhileLoopAction MultiImportInnerLoop(FormatData formatData, BlockData blockData, Token token, ref TokenType lastTokenType, string thisFunctionName)
            {
                bool isValid;

                switch (token.Type)
                {
                    case TokenType.StringValue:

                        isValid =
                            lastTokenType == TokenType.Identifier ||
                            lastTokenType == TokenType.FullStop ||
                            lastTokenType == TokenType.NewLine;

                        if (isValid) {
                            blockData.Tokens.Add(token);
                            break;
                        }
                        formatData.UnexpectedTypeError(token, $"invalid import lasttype {lastTokenType}, token: {token.Text}", thisFunctionName);
                        return WhileLoopAction.Return;

                    case TokenType.Identifier:
                        isValid =
                            lastTokenType == TokenType.NewLine;

                        if (isValid) {
                            blockData.Tokens.Add(token);
                            break;
                        }
                        formatData.UnexpectedTypeError(token, $"invalid import lasttype {lastTokenType}, token: {token.Text}", thisFunctionName);
                        return WhileLoopAction.Return;

                    case TokenType.NewLine:
                        isValid =
                            lastTokenType == TokenType.StringValue ||
                            lastTokenType == TokenType.LeftParenthesis ||
                            lastTokenType == TokenType.NewLine;

                        if (isValid) {
                            blockData.Tokens.Add(token);
                            break;
                        }
                        formatData.UnexpectedTypeError(token, $"invalid import lasttype {lastTokenType}, token: {token.Text}", thisFunctionName);
                        return WhileLoopAction.Return;

                    case TokenType.FullStop:
                        blockData.Tokens.Add(token);
                        return WhileLoopAction.Continue;

                    case TokenType.RightParenthesis:
                        blockData.Tokens.Add(token);
                        formatData.Increment();
                        return WhileLoopAction.Break;

                    default:
                        formatData.UnexpectedTypeError(token, $"invalid import lasttype {lastTokenType}, token: {token.Text}", thisFunctionName);
                        return WhileLoopAction.Return;
                }
                return WhileLoopAction.Continue;
            }

            static void ProcessMultiImport(FormatData formatData, CodeBlock globalBlock, BlockData blockData, Token leftParenthToken)
            {
                const string thisFunctionName = "ProcessMultiImport";
                blockData.Tokens.Add(leftParenthToken);
                formatData.Increment();
                TokenType lastTokenType = TokenType.LeftParenthesis;
                blockData.NodeType = NodeType.Multi_Line_Import;

                while (formatData.TokenIndex < formatData.TokenList.Count) {

                    int previousIndex = formatData.TokenIndex;
                    Token? tempToken = formatData.GetToken();
                    if (tempToken == null) {
                        formatData.EndOfFileError(leftParenthToken, thisFunctionName);
                        return;
                    }
                    Token token = tempToken.Value;

                    WhileLoopAction result = MultiImportInnerLoop(formatData, blockData, token, ref lastTokenType, thisFunctionName);
                    if (result == WhileLoopAction.Return) {
                        return;
                    } 
                    if (result == WhileLoopAction.Break) {
                        break;
                    }

                    lastTokenType = token.Type;
                    formatData.IncrementIfSame(previousIndex);
                }

                globalBlock.BlockDataList.Add(blockData);
            }

            public static void ProcessImport(FormatData formatData, CodeBlock globalBlock, Token packageToken)
            {
                const string thisFunctionName = "ProcessImport";

                BlockData blockData = new();
                Debug.Assert(blockData.Tokens != null);
                blockData.StartingToken = packageToken;
                blockData.Tokens.Add(packageToken);

                Token? tempToken = formatData.GetNextToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(packageToken, thisFunctionName);
                    return;
                }
                Token nextToken = tempToken.Value; 

                if (nextToken.IsType(TokenType.LeftParenthesis)) {
                    ProcessMultiImport(formatData, globalBlock, blockData, nextToken);
                    return;
                }
                ProcessSingleImport(formatData, globalBlock, blockData, ref nextToken);
            }

            private static void ProcessSingleImport(FormatData formatData, CodeBlock globalBlock, BlockData blockData, ref Token nextToken)
            {
                const string thisFunctionName = "ProcessSingleImport";

                Token? tempToken = null;

                blockData.NodeType = NodeType.Single_Import;
                //Is Alias
                if (nextToken.Type == TokenType.Identifier)
                {

                    blockData.NodeType = NodeType.Single_Import_With_Alias;
                    blockData.Tokens.Add(nextToken);

                    tempToken = formatData.GetNextToken();
                    if (tempToken == null)
                    {
                        formatData.EndOfFileError(nextToken, thisFunctionName);
                        return;
                    }
                    nextToken = tempToken.Value;
                }
                //Import name
                if (nextToken.Type == TokenType.StringValue)
                {

                    blockData.Tokens.Add(nextToken);
                    tempToken = formatData.GetNextToken();
                    if (tempToken == null)
                    {
                        formatData.EndOfFileError(nextToken, thisFunctionName);
                        return;
                    }
                    nextToken = tempToken.Value;
                }
                //new line
                if (nextToken.Type == TokenType.NewLine)
                {

                    blockData.Tokens.Add(nextToken);
                    globalBlock.BlockDataList.Add(blockData);
                    return;
                }

                formatData.UnexpectedTypeError(nextToken, "unexpected type in import", thisFunctionName);
            }
        }

        public static class Structs
        {
            public static void ProcessStruct(FormatData formatData, CodeBlock globalBlock, Token structToken)
            {
                const string thisFunctionName = "ProcessStruct";

                BlockData blockData = new();
                blockData.StartingToken = structToken;
                blockData.NodeType = NodeType.Struct_Declaration;

                //Add struct name
                Token? tempToken = formatData.GetNextToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }
                Token structNameToken = tempToken.Value;
                blockData.Tokens.Add(structNameToken);

                if (formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in struct declaration", thisFunctionName) == false ) { 
                    return;
                }
                formatData.Increment();
                formatData.IncrementIfNewLine();

                CodeBlock? structBlock = FormatBody.FillStructBody(formatData);
                if (structBlock == null) {
                    formatData.AddTrace(thisFunctionName);
                    return;
                }
                if (formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in struct declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                blockData.Block = structBlock;
                globalBlock.BlockDataList.Add(blockData);
            }
        }
    
        public static class Enums
        {
            public static void ProcessEnum(FormatData formatData, CodeBlock globalBlock, Token structToken)
            {
                const string thisFunctionName = "ProcessEnum";

                BlockData blockData = new();
                blockData.StartingToken = structToken;
                blockData.NodeType = NodeType.Enum_Declaration;

                //Add struct name
                Token? tempToken = formatData.GetNextToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }
                Token enumNameToken = tempToken.Value;
                blockData.Tokens.Add(enumNameToken);

                if (formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in enum declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                formatData.IncrementIfNewLine();

                CodeBlock? enumBlock = FillEnumBody(formatData);
                if (enumBlock == null) {
                    formatData.AddTrace(thisFunctionName);
                    return;
                }
                if (formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in enum declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                blockData.Block = enumBlock;
                globalBlock.BlockDataList.Add(blockData);
            }

            public static void ProcessEnumstruct(FormatData formatData, CodeBlock globalBlock, Token structToken)
            {
                const string thisFunctionName = "ProcessEnumstruct";

                BlockData blockData = new();
                blockData.StartingToken = structToken;
                blockData.NodeType = NodeType.Enum_Struct_Declaration;

                //Add struct name
                Token? tempToken = formatData.GetNextToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }

                Token enumNameToken = tempToken.Value;
                blockData.Tokens.Add(enumNameToken);

                if (formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in enum declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                formatData.IncrementIfNewLine();

                CodeBlock? enumBlock = FillEnumBody(formatData);
                if (enumBlock == null) {
                    formatData.AddTrace(thisFunctionName);
                    return;
                }
                if (formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in enum declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                blockData.Block = enumBlock;
                globalBlock.BlockDataList.Add(blockData);
            }

            static CodeBlock? FillEnumBody(FormatData formatData)
            {
                CodeBlock block = new CodeBlock {
                    BlockDataList = new List<BlockData>(),
                };

                BlockData tempBlockData = new();
                TokenType lastNodeType = TokenType.NA;

                int whileCount = 0;

                while (formatData.TokenIndex < formatData.TokenList.Count) {

                    whileCount += 1;
                    if (whileCount > 10000) {
                        formatData.Result = FormatResult.Infinite_While_Loop;
                        return null;
                    }

                    Token? token = formatData.GetToken();
                    if (token == null) {
                        formatData.EndOfFileError(token, "FillEnumBody");
                        return null;
                    }

                    if (token.Value.Type == TokenType.NewLine) {
                        block.BlockDataList.Add(tempBlockData);
                        tempBlockData = new BlockData();
                        formatData.Increment();
                        continue;
                    }
                    if (token.Value.Type == TokenType.RightBrace) {
                        break;
                    }

                    tempBlockData.Tokens.Add(token.Value);

                    lastNodeType = token.Value.Type;
                    formatData.Increment();
                }

                return block;
            }

        }

        public static class Variables
        {
            public static void ProcessConstant(FormatData formatData, CodeBlock globalBlock, Token constToken)
            {
                const string THIS_FUNCTION = "ProcessConstant";

                BlockData blockData = new BlockData {
                    Block = null,
                    NodeType = NodeType.Constant_Global_Variable,
                    StartingToken = constToken,
                    Tokens = new List<Token>(),
                    Variables = new List<Variable>(),
                };

                int startingIndex = formatData.TokenIndex;
                formatData.Increment();
                Token? tempToken = formatData.GetToken();
                if (tempToken == null) {
                    formatData.EndOfFileError(tempToken, THIS_FUNCTION);
                    return;
                }
                if (tempToken.Value.Type == TokenType.Identifier) {
                    blockData.Tokens.Add(tempToken.Value);
                    formatData.Increment();
                    LoopValue(formatData, globalBlock, blockData);
                    return;
                }
                if (IsVarType(tempToken.Value.Type) == true) {

                    blockData.NodeType = NodeType.Constant_Global_Variable_With_Type;
                    Variable variable = new();
                    variable.TypeList.Add(tempToken.Value);
                    formatData.Increment();
                    tempToken = formatData.GetToken();
                    if (tempToken == null) {
                        formatData.EndOfFileError(tempToken, THIS_FUNCTION);
                        return;
                    }
                    if (tempToken.Value.Type != TokenType.Identifier) {
                        formatData.UnexpectedTypeError(tempToken, "missing expended identifier in constant variable", THIS_FUNCTION);
                    }

                    variable.NameToken.Add(tempToken.Value);
                    blockData.Variables.Add(variable);
                    formatData.Increment();

                    LoopValue(formatData, globalBlock, blockData);

                }
            }

            static void LoopValue(FormatData formatData, CodeBlock globalBlock, BlockData blockData)
            {
                while (formatData.TokenIndex < formatData.TokenList.Count)
                {
                    Token? token = formatData.GetToken();
                    if (token == null) {
                        formatData.EndOfFileError(token, "LoopValue");
                        return;
                    }
                    if (token.Value.Type == TokenType.NewLine) {
                        break;
                    }
                    blockData.Tokens.Add(token.Value);
                    formatData.Increment();
                }

                globalBlock.BlockDataList.Add(blockData);
            }

            static bool IsVarType(TokenType type)
            {
                switch (type)
                {
                    case TokenType.Int:
                    case TokenType.Int8:
                    case TokenType.Int16:
                    case TokenType.Int32:
                    case TokenType.Int64:
                    case TokenType.Uint:
                    case TokenType.Uint8:
                    case TokenType.Uint16:
                    case TokenType.Uint32:
                    case TokenType.Uint64:
                    case TokenType.Float32:
                    case TokenType.Float64:
                    case TokenType.String:
                    case TokenType.Byte:
                    case TokenType.Rune:
                        return true;
                   
                    default:
                        return false;
                }
            }
        }

        public static class Interfaces
        {
            public static void Process(FormatData formatData, CodeBlock globalBlock, Token interfaceToken)
            {
                const string thisFunctionName = "ProcessInterface";

                BlockData blockData = new();
                blockData.StartingToken = interfaceToken;
                blockData.NodeType = NodeType.Interface_Declaration;

                //Add struct name
                Token? tempToken = formatData.GetNextToken();
                if (tempToken == null)
                {
                    formatData.EndOfFileError(tempToken, thisFunctionName);
                    return;
                }
                Token structNameToken = tempToken.Value;
                blockData.Tokens.Add(structNameToken);

                if (formatData.ExpectNextType(TokenType.LeftBrace, "missing expected '{' in interface declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                formatData.IncrementIfNewLine();

                CodeBlock? structBlock = FormatBody.FillInterfaceBody(formatData);
                if (structBlock == null) {
                    formatData.AddTrace(thisFunctionName);
                    return;
                }
                if (formatData.ExpectType(TokenType.RightBrace, "missing expected '}' in interface declaration", thisFunctionName) == false) {
                    return;
                }
                formatData.Increment();
                blockData.Block = structBlock;
                globalBlock.BlockDataList.Add(blockData);
            }
        }

    }
}
