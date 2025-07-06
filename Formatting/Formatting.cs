using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class Formatting
    {
        public static CodeFormat? FormatCode(List<Token> tokenList, ref FormatResult result, string code)
        {
            const string thisFunctionName = "FormatCode";

            //Fmt.PrintColor("\tFormatting:\t\t\t", ConsoleColor.DarkGray);

            CodeBlock globalBlock = new();
            List<Function> functionList = new();

            FormatData formatData = new FormatData {
                TokenList = tokenList,
                ErrorToken = null,
                ErrorDetail = "",
                Result = FormatResult.Ok,
                ErrorTrace = new List<string>(),
            };

            if (tokenList.Count == 0) {
                result = FormatResult.NoTokens;
                return null;
            }

            while (formatData.TokenIndex < formatData.TokenList.Count) {

                Token? token = formatData.GetToken();
                if (token == null) {
                    formatData.EndOfFileError(token, thisFunctionName);
                    return null;
                }

                int previousIndex = formatData.TokenIndex;

                BlockData? blockData = null;

                switch (token.Value.Type)
                {
                    case TokenType.Chan:
                        blockData = FormatChan.ProcessChannel(formatData, token.Value);
                        if (blockData == null) {
                            return null;
                        }
                        globalBlock.BlockDataList.Add(blockData);
                        break;

                    case TokenType.Const:
                        FormatGlobalStuff.Variables.ProcessConstant(formatData, globalBlock, token.Value);
                        break;
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
                    case TokenType.Bool:
                    case TokenType.LeftSquareBracket:
                    case TokenType.Error:
                    case TokenType.Multiply:
                        blockData = FormatDeclarations.ProcessDeclaration(formatData, token.Value);
                        if (blockData == null) {
                            return null;
                        }
                        globalBlock.BlockDataList.Add(blockData);
                        break;

                    case TokenType.Type:
                        blockData = FormatReturn.ProcessReturn(formatData);
                        if (blockData == null) {
                            return null;
                        }
                        globalBlock.BlockDataList.Add(blockData);
                        break;

                    case TokenType.Package:
                        FormatGlobalStuff.Packages.ProcessPackage(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Import:
                        FormatGlobalStuff.Imports.ProcessImport(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Struct:
                        FormatGlobalStuff.Structs.ProcessStruct(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Enum:
                        FormatGlobalStuff.Enums.ProcessEnum(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Enumstruct:
                        FormatGlobalStuff.Enums.ProcessEnumstruct(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Interface:
                        FormatGlobalStuff.Interfaces.Process(formatData, globalBlock, token.Value);
                        break;

                    case TokenType.Fn:
                        FormatFunction.ProcessFunction(formatData, functionList, token.Value);
                        break;

                    case TokenType.Var:
                        formatData.Result = FormatResult.UnexpectedType;
                        formatData.ErrorDetail = "'var' isn't ever used in TypeGo. Correct syntax: 'int number = 10', not 'var number int = 10'";
                        formatData.ErrorToken = token;
                        return null;

                    default:
                        break;
                }

                if (formatData.IsError()) {
                    Fmt.PrintlnColor($"Error {formatData.Result}", ConsoleColor.Red);
                    Debugging.PrintFormatError(ref formatData, code);
                    result = formatData.Result;
                    return null;
                }

                formatData.IncrementIfSame(previousIndex);
            }

            //globalBlock.PrintData();


            CodeFormat codeFormat = new CodeFormat {
                Functions = functionList, 
                GlobalBlock = globalBlock,
            };

            //Fmt.PrintlnColor("Done", ConsoleColor.DarkCyan);

            return codeFormat;
        }
    }
}
