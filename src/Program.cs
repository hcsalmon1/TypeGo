
using System.IO;

namespace TypeGo
{
    public static class Program
    {

        static string? ConvertToGo(string code, string fileName)
        {
            Fmt.PrintColor("\tConverting: ", ConsoleColor.DarkGray);
            Fmt.Print($"{fileName} - ");

            ParseResult parseResult = ParseResult.Ok;
            List<Token> tokenList = Parse.ParseToTokens(ref parseResult, code);
            if (parseResult != ParseResult.Ok) {
                Fmt.Println($"Error: {parseResult}");
                return null;
            }

            //Debugging.PrintTokenList(tokenList);

            FormatResult formatResult = FormatResult.Ok;
            CodeFormat? codeFormat = Formatting.FormatCode(tokenList, ref formatResult, code);
            if (formatResult != FormatResult.Ok) {
                return null;
            }

            ConvertResult convertResult = ConvertResult.Ok;
            string? goCode = ConvertCode.ConvertToGo(codeFormat, ref convertResult, code);
            if (convertResult != ConvertResult.Ok) {
                return null;
            }

            return goCode;
        }

        static bool ConvertAndWriteFile(string tgoFilePath)
        {
            //Fmt.Println();

            string tgoCode = File.ReadAllText(tgoFilePath);
            string? fileName = Path.GetFileName(tgoFilePath);

            string? goCode;

            if (fileName == null) {
                goCode = ConvertToGo(tgoCode, "unknown");
            } else {
                goCode = ConvertToGo(tgoCode, fileName);
            }

            if (goCode == null) {
                //Fmt.Println($"Skipped: {tgoFilePath} (conversion returned null)");
                return false;
            }

            string goPath = Path.ChangeExtension(tgoFilePath, ".go");
            File.WriteAllText(goPath, goCode);

            Fmt.PrintlnColor("Done", ConsoleColor.DarkCyan);
            return true;
        }

        static void ConvertFile(string filePath)
        {
            if (!File.Exists(filePath))
            {
                Fmt.Println($"Error: File not found: {filePath}");
                return;
            }

            ConvertAndWriteFile(filePath);
        }

        static void ConvertDirectory(string currentDirectory)
        {
            string[] typeGoFiles = Directory.GetFiles(currentDirectory, "*.tgo", SearchOption.AllDirectories);

            if (typeGoFiles.Length == 0) {
                Fmt.Println("no tgo files found");
                return;
            }

            foreach (string filePath in typeGoFiles) {
                if (ConvertAndWriteFile(filePath) == false) {
                    return;
                }
            }

            Fmt.Println("Done");
        }

        const string HELP = "help";
        const string VERSION = "version";
        const string CONVERT_DIRECTORY = "convertdir";
        const string CONVERT_FILE = "convertfile";


        static void ConsoleInput(string[] args)
        {
            if (args.Length == 0)
            {
                ShowHelp();
                return;
            }

            string command = args[0].ToLower();

            switch (command)
            {
                case HELP:
                    ShowHelp();
                    break;

                case VERSION:
                    Console.WriteLine("TypeGo version 0.1");
                    break;

                case CONVERT_FILE:
                    if (args.Length < 2)
                    {
                        Console.WriteLine("Error: You must specify a filename (e.g. typego convertfile myfile.tgo)");
                    }
                    else
                    {
                        string filePath = Path.Combine(Directory.GetCurrentDirectory(), args[1]);
                        ConvertFile(filePath);
                    }
                    break;

                case CONVERT_DIRECTORY:
                    string currentDir = Directory.GetCurrentDirectory();
                    ConvertDirectory(currentDir);
                    break;

                default:
                    Console.WriteLine($"Unknown command: {command}");
                    ShowHelp();
                    break;
            }
        }

        static void TestingInput()
        {
            const string CODE = @"

package main

import ""fmt""

fn main() {

}
    ";

            Fmt.Println($"Inputted code:\n{CODE}");

            string? goCode = ConvertToGo(CODE, "none");
            if (goCode == null) {
                return;
            }

            Fmt.PrintlnColor($"Done", ConsoleColor.DarkCyan);

            Fmt.Println($"\nConverted Go code: \n{goCode}");
        }

        static void Main(string[] args)
        {
            ConsoleInput(args);

            //TestingInput();
        }

        static void ShowHelp()
        {
            Console.WriteLine("TypeGo");
            Console.WriteLine("Usage:");
            Console.WriteLine("  tgo help                 Show this help message");
            Console.WriteLine("  tgo version              Show the version of TypeGo");
            Console.WriteLine("  tgo convertfile file.tgo   Convert a single .tgo file in the current directory");
            Console.WriteLine("  tgo convertdir           Convert all .tgo files in the current directory (recursively)");
        }
    }
}
