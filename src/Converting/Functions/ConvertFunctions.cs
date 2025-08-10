

namespace TypeGo
{
    public static class ConvertFunctions
    {
        public static void ProcessFunction(ConvertData convertData, Function function)
        {
            PrintFunctionNameAndParameters(convertData, function);

            if (function.ReturnType != null) {
                convertData.GeneratedCode.Append(function.ReturnType);
                convertData.GeneratedCode.Append(' ');
            }
            convertData.GeneratedCode.Append('{');

            if (function.InnerBlock == null) {
                convertData.ConvertResult = ConvertResult.Internal_Error;
                convertData.ErrorDetail = "inner block is null in ProcessFunction";
                return;
            }

            ConvertBlock.ProcessBlock(convertData, function.InnerBlock, nestCount: 1);
            //convertData.GeneratedCode.Append('\n');
            convertData.GeneratedCode.Append('\r');
            convertData.GeneratedCode.Append('}');
            convertData.GeneratedCode.Append('\n');
            //convertData.GeneratedCode.Append('\n');
        }

        private static void PrintFunctionNameAndParameters(ConvertData convertData, Function function)
        {
            if (convertData.IsMethod) {
                char firstLetter = char.ToLower(convertData.StructName[0]);
                convertData.GeneratedCode.Append($"func ({firstLetter} *{convertData.StructName}) {function.Name}(");
            } else {
                convertData.GeneratedCode.Append($"func {function.Name}(");
            }

            PrintParameters(convertData, function);
            convertData.GeneratedCode.Append(')');
            convertData.GeneratedCode.Append(' ');
        }

        private static void PrintParameters(ConvertData convertData, Function function)
        {
            List<Variable> parameters = function.Parameters;
            if (parameters.Count == 0) {
                return;
            }

            for (int parameterIndex = 0; parameterIndex < parameters.Count; parameterIndex++) {

                Variable parameter = parameters[parameterIndex]; 
                if (parameter.NameToken == null) {
                    convertData.ConvertResult = ConvertResult.Internal_Error;
                    convertData.ErrorDetail = "name token is null in PrintParameters";
                    convertData.ErrorToken = function.startingToken;
                    return;
                }
                string? typeAsText = TokenUtils.JoinTextInListOfTokens(parameter.TypeList);
                if (typeAsText == null) {
                    convertData.ConvertResult = ConvertResult.Internal_Error;
                    convertData.ErrorDetail = "var type text is null in PrintParameters";
                    convertData.ErrorToken = function.startingToken;
                    return;
                }
                if (parameterIndex != 0) {
                    convertData.GeneratedCode.Append(',');
                    convertData.GeneratedCode.Append(' ');
                }
                if (parameter.NameToken.Count == 0) {
                    convertData.ConvertResult = ConvertResult.Internal_Error;
                    convertData.ErrorDetail = "parameter name invalid in PrintParameters";
                    convertData.ErrorToken = function.startingToken;
                    return;
                }
                convertData.GeneratedCode.Append($"{parameter.NameToken[0].Text} {typeAsText}");
            }
        }
    }
}
