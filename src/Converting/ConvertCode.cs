using System;
using System.Collections.Generic;
using System.Text;


namespace TypeGo
{
    public static class ConvertCode
    {

        static bool ShouldReturn(ref ConvertData convertData, ref int whileCount, int WHILE_CAP, string thisFunctionName)
        {
            if (Debugging.InfiniteWhileCheck(ref whileCount, WHILE_CAP)) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "InfiniteWhileLoop in Should Return";
                return true;
            }
            if (convertData.IsError()) {
                convertData.AddTrace(thisFunctionName);
                return true;
            }
            return false;
        }

        static void PrintCodeFormat(ConvertData convertData)
        {
            const string thisFunctionName = "LoopTokens";

            if (convertData.CodeFormat.GlobalBlock == null) {
                convertData.ConvertResult = ConvertResult.Missing_Expected_Type;
                convertData.ErrorDetail = "Internal error: CodeFormat is null";
                return;
            }
            CodeBlock globalBlock = convertData.CodeFormat.GlobalBlock;

            ConvertBlock.ProcessBlock(convertData, globalBlock, nestCount: 0);

            List<Function> functions = convertData.CodeFormat.Functions;

            for (int functionIndex = 0; functionIndex < functions.Count; functionIndex++) {

                ConvertFunctions.ProcessFunction(convertData, functions[functionIndex]);

                if (convertData.IsError()) {
                    convertData.AddTrace(thisFunctionName);
                    return;
                }
            }
        }

        public static string? ConvertToGo(CodeFormat codeFormat, ref ConvertResult convertResult, string code)
        {
            //Fmt.PrintColor("\tConverting:\t\t\t", ConsoleColor.DarkGray);
            ConvertData convertData = new ConvertData {
                CodeFormat = codeFormat,
                ErrorDetail = "",
                ConvertResult = ConvertResult.Ok,
                GeneratedCode = new StringBuilder(),
                ErrorTrace = new List<string>(),
            };

            PrintCodeFormat(convertData);
            if (convertData.IsError()){
                Fmt.PrintColor($"Error {convertData.ConvertResult}\n", ConsoleColor.Red);
                Debugging.PrintConvertError(ref convertData, code);
                return null;
            }

            //Fmt.PrintColor("Done\n", ConsoleColor.DarkCyan);

            return convertData.GeneratedCode.ToString();
        }
    }
}
