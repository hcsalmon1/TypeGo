using System;
using System.Collections.Generic;

namespace TypeGo
{
    public static class PrintAppend
    {
        public static void Print(ConvertData convertData, BlockData blockData, int nestCount)
        {
            if (blockData.Variables.Count == 0) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "variables count in PrintAppend";
                return;
            }

            Variable variable = blockData.Variables[0];

            if (variable.NameToken == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "NameToken is null in PrintAppend";
                return;
            }

            string variableName = variable.NameToken[0].Text;

            convertData.AppendString($"{variableName} = append({variableName}, ");

            if (blockData.Tokens.Count == 0) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "tokens count is 0 in PrintAppend";
                return;
            }

            for (int i = 0; i < blockData.Tokens.Count; i++)
            {
                convertData.AppendToken(blockData.Tokens[i]);
            }

            convertData.AppendChar(')');
            convertData.NewLineWithTabs(nestCount);
        }
    }
}
