using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class FormatBody
    {

        

        static void ProcessTokenInBody(FormatData formatData, CodeBlock block, Token token)
        {
            const string thisFunctionName = "ProcessTokenInBody";

            BlockData? blockData = null;

            switch (token.Type)
            {
                case TokenType.Chan:
                    blockData = FormatChan.ProcessChannel(formatData, token);
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
                case TokenType.Map:
                    blockData = FormatDeclarations.ProcessDeclaration(formatData, token);
                    break;
                case TokenType.If:
                    blockData = FormatIf.ProcessIf(formatData, token);
                    break;
                case TokenType.Else:
                    blockData = FormatIf.ProcessElse(formatData, token);
                    break;
                case TokenType.For:
                    blockData = FormatFor.ProcessFor(formatData, token);
                    break;
                case TokenType.Return:
                    blockData = FormatReturn.ProcessReturn(formatData);
                    break;
                case TokenType.Break:
                    blockData = new BlockData();
                    blockData.NodeType = NodeType.Other;
                    blockData.Tokens.Add(token);
                    break;
                case TokenType.Continue:
                    blockData = new BlockData();
                    blockData.NodeType = NodeType.Other;
                    blockData.Tokens.Add(token);
                    break;
                case TokenType.Defer:
                    blockData = FormatDefer.ProcessDefer(formatData, in token);
                    break;

                case TokenType.Goto:
                    blockData = FormatDefer.ProcessDefer(formatData, in token);
                    break;

                case TokenType.Identifier:
                    blockData = FormatIdentifier.ProcessIdentifier(formatData, in token);
                    break;
                case TokenType.LeftParenthesis:
                    break;
                case TokenType.Semicolon:
                    return;
                case TokenType.Tab:
                    break;
                case TokenType.MultiLineStart:
                    break;
                case TokenType.MultiLineEnd:
                    break;
                case TokenType.Comment:
                    blockData = FormatComment.ProcessComment(formatData);
                    break;
                case TokenType.EndComment:
                    return;
                case TokenType.NewLine:
                    blockData = new();
                    blockData.NodeType = NodeType.NewLine;
                    break;
                case TokenType.Var:
                    formatData.Result = FormatResult.UnexpectedType;
                    formatData.ErrorDetail = "'var' isn't ever used in TypeGo. Correct syntax: 'int number = 10', not 'var number int = 10'";
                    formatData.ErrorToken = token;
                    return;
                case TokenType.Go:
                    blockData = FormatGo.ProcessGo(formatData, in token);
                    break;

                case TokenType.Switch:
                    blockData = FormatSwitch.ProcessSwitch(formatData);
                    break;

                case TokenType.Interface:
                    blockData = FormatInterface.FormatInterfaceVar(formatData, token);
                    break;

                default:
                    break;
            }

            if (formatData.IsError()) {
                formatData.AddTrace(thisFunctionName);
                return;
            }
            if (blockData == null) {
                formatData.UnsupportedFeatureError(token, thisFunctionName);
                return;
            }
            block.BlockDataList.Add(blockData);

        }

        public static CodeBlock? FillBody(FormatData formatData)
        {
            const string thisFunctionName = "FillBody";

            CodeBlock block = new CodeBlock {
                BlockDataList = new List<BlockData>(),
            };

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {
                int previousIndex = formatData.TokenIndex;

                Token? token = formatData.GetToken();
                if (token == null) {
                    formatData.EndOfFileError(null, thisFunctionName);
                    return null;
                }

                if (token.Value.Type == TokenType.RightBrace) {
                    break;
                }

                ProcessTokenInBody(formatData, block, token.Value);
                if (formatData.IsError()){
                    formatData.AddTrace(thisFunctionName);
                    return null;
                }

                formatData.IncrementIfSame(previousIndex);

            }
            return block;
        }

        public static CodeBlock? FillStructBody(FormatData formatData)
        {
            const string thisFunctionName = "FillStructBody";

            CodeBlock block = new CodeBlock
            {
                BlockDataList = new List<BlockData>(),
                MethodList = new List<Function>()
            };

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {
                int previousIndex = formatData.TokenIndex;

                Token? token = formatData.GetToken();
                if (token == null) {
                    formatData.EndOfFileError(null, thisFunctionName);
                    return null;
                }

                if (token.Value.Type == TokenType.RightBrace) {
                    break;
                }
                if (token.Value.Type == TokenType.Fn) {
                    FormatFunction.ProcessFunction(formatData, block.MethodList, token.Value);
                    formatData.Increment();
                    continue;
                }

                ProcessTokenInBody(formatData, block, token.Value);
                if (formatData.IsError())
                {
                    formatData.AddTrace(thisFunctionName);
                    return null;
                }

                formatData.IncrementIfSame(previousIndex);
            }
            return block;
        }

        public static CodeBlock? FillInterfaceBody(FormatData formatData)
        {
            const string thisFunctionName = "FillInterfaceBody";

            CodeBlock block = new CodeBlock {
                BlockDataList = new List<BlockData>(),
                MethodList = new List<Function>()
            };

            while (formatData.TokenIndex < formatData.TokenList.Count)
            {
                int previousIndex = formatData.TokenIndex;

                Token? token = formatData.GetToken();
                if (token == null) {
                    formatData.EndOfFileError(null, thisFunctionName);
                    return null;
                }

                if (token.Value.Type == TokenType.RightBrace) {
                    break;
                }
                if (token.Value.Type == TokenType.NewLine) {
                    formatData.Increment();
                    continue;
                }
                FormatFunction.ProcessInterfaceFunction(formatData, block.MethodList, token.Value);
                formatData.IncrementIfSame(previousIndex);
            }
            return block;
        }
    }
}
