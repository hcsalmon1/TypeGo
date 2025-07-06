using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class PrintStruct
    {
        public static void ProcessStruct(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Tokens == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Tokens is null in ProcessStruct";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            if (blockData.Tokens.Count != 1) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "Token count is 1 in ProcessStruct";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }
            string structNameText = blockData.Tokens[0].Text;

            convertData.AppendString($"type {structNameText} struct {{");
            convertData.NewLineWithTabs(nestCount + 1);

            CodeBlock? structVariableBlock = blockData.Block;
            if (structVariableBlock == null) {
                convertData.AppendChar('\n');
                EndStruct(convertData);
                return;
            }

            if (structVariableBlock.BlockDataList == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "BlockDataList is null in ProcessStruct";
                convertData.ErrorToken = blockData.StartingToken;
                return;
            }

            List<BlockData> blockDataList = structVariableBlock.BlockDataList;

            if (blockDataList.Count == 0) {
                convertData.AppendChar('\n');
                EndStruct(convertData);
                return;
            }

            List<string> varNames = new();

            for (int blockIndex = 0; blockIndex < blockDataList.Count; blockIndex++) {

                BlockData varBlock = blockDataList[blockIndex];

                if (varBlock.NodeType != NodeType.Single_Declaration_No_Value) {
                    PrintOther(convertData, varBlock, nestCount);
                    continue;
                }

                for (int i = 0; i < varBlock.Variables.Count; i++)
                {
                    Variable variable = varBlock.Variables[i];
                    if (variable.NameToken == null) {
                        continue;
                    }
                    varNames.Add(variable.NameToken.Value.Text);
                    string? varTypeAsText = TokenUtils.JoinTextInListOfTokens(variable.TypeList);
                    if (varTypeAsText == null) {
                        continue;
                    }
                    convertData.AppendString($"{variable.NameToken.Value.Text} {varTypeAsText}");
                    convertData.NewLineWithTabs(nestCount + 1);
                }
            }

            convertData.AppendChar('\r');
            EndStruct(convertData);
            PrintMethods(convertData, structVariableBlock, nestCount, varNames, structNameText);
        }

        static void PrintMethods(ConvertData convertData, CodeBlock structBlock, int nestCount, List<string> varName, string structName)
        {
            List<Function> functions = structBlock.MethodList;
            if (functions == null) {
                return;
            }
            if (functions.Count == 0) {
                return;
            }

            convertData.MethodVarNames = varName;
            convertData.IsMethod = true;
            convertData.StructName = structName;

            for (int i = 0; i < functions.Count; i++)
            {
                ConvertFunctions.ProcessFunction(convertData, functions[i]);
            }
            convertData.IsMethod = false;
            convertData.MethodVarNames.Clear();
        }

        static void PrintOther(ConvertData convertData, BlockData varBlock, int nestCount)
        {
            if (varBlock.NodeType != NodeType.Other) {
                return;
            }

            if (varBlock.Tokens.Count == 0) {
                return;
            }

            for (int i = 0; i < varBlock.Tokens.Count; i++)
            {
                Token token = varBlock.Tokens[i];

                if (token.Type == TokenType.NewLine) {

                    continue;
                }
                convertData.AppendString(varBlock.Tokens[i].Text);
            }
            convertData.NewLineWithTabs(nestCount + 1);
        }

        private static void EndStruct(ConvertData convertData)
        {
            convertData.AppendChar('}');
            convertData.AppendChar('\n');
            convertData.AppendChar('\n');
        }
    }
}
