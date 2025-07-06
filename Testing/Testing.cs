using System;
using System.Collections.Generic;
using System.Text;

namespace TypeGo
{
    public static class Testing
    {
        public static class Declarations
        {
            public static void TestDeclarations()
            {
                string[] declarations = {
                    "int number = 10;",
                    @"string message = ""hello""
                    ",
                    @"*int ptr = &number
                    ",
                    @"[]string messages = {""This"", ""is"", ""an"", ""example"",}",
                    @"*[]bool conditions = &condition_mer
                    ",
                    @"int number = 10; int value = 20
                    ",
                    @"int number, value
                    ",
                };

                for (int i = 0; i < declarations.Length; i++)
                {
                    IndividualTest(declarations[i]);
                    Fmt.Println();
                    Fmt.Println();
                }
            }

            static void IndividualTest(string code)
            {
                Fmt.Println($"Testing: '{code}'");
                ParseResult parseResult = ParseResult.Ok;
                List<Token> tokenList = Parse.ParseToTokens(ref parseResult, code);
                if (parseResult != ParseResult.Ok)
                {
                    Fmt.PrintlnRed($"\nError: {parseResult}");
                    return;
                }

                FormatData formatData = new FormatData
                {
                    TokenList = tokenList,
                    ErrorToken = null,
                    ErrorDetail = "",
                    Result = FormatResult.Ok,
                    ErrorTrace = new List<string>(),
                };

                BlockData? declarationData = FormatDeclarations.ProcessDeclaration(formatData, tokenList[0]);
                if (formatData.IsError())
                {
                    Fmt.PrintlnRed($"\nError: {formatData.Result}");
                    return;
                }

                if (declarationData == null)
                {
                    Fmt.PrintlnRed("\nError declaration data = null");
                    return;
                }

                declarationData.Print();
            }
        }
    }
}
