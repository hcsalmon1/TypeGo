

namespace TypeGo
{
    public static class Program
    {

        static string? ConvertToGo(string code, string fileName)
        {

            //Logging.CreateLog();

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

            //codeFormat.PrintNodeType();

            ConvertResult convertResult = ConvertResult.Ok;
            string? goCode = ConvertCode.ConvertToGo(codeFormat, ref convertResult, code);
            if (convertResult != ConvertResult.Ok) {
                return null;
            }

            //Logging.PrintLog();

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
            if (args.Length == 0) {
                ShowHelp();
                return;
            }

            string command = args[0].ToLower();

            switch (command) {

                case HELP:
                    ShowHelp();
                    break;

                case VERSION:
                    Console.WriteLine("TypeGo version 0.3");
                    break;

                case CONVERT_FILE:

                    if (args.Length < 2) {
                        Console.WriteLine("Error: You must specify a filename (e.g. typego convertfile myfile.tgo)");
                    } else {
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

            string CODE = @"


";

            string CODE2 = @"

package main

import (
    ""bufio""
    ""fmt""
    ""os""
    ""path/filepath""
    ""strings""
)

string extension_to_find = ""NA""
string phrase_to_find = ""NA""
string phrase_to_replace = ""NA""

fn printCommands() {
    fmt.Println(""\nCommands:\n extension name - sets the file extension to look for. E.g extension c\n help - shows this command list\n find example - it will look 'example' in your files\n replace EXAMPLE - will replace all found phrases with this\n run - will find and replace in all files in directory\n exit - ends the program\n"");
}

fn bool containsCommand(string command, string input) {
    int command_length = len(command)
    int input_length = len(input)

    if input_length <= command_length {
        return false;
    }

    for i := 0; i < command_length; i += 1 {
        if input[i] != command[i] {
            return false;
        }
    }
    return true;
}

fn findAndSetExtension(string input) {
    int START_INDEX = 10;
    extension_to_find = input[START_INDEX:];
}

fn findAndSetFind(string input) {
    int START_INDEX = 5;
    phrase_to_find = input[START_INDEX:];
}

fn findAndSetReplace(string input) {
    int START_INDEX = 8;
    phrase_to_replace = input[START_INDEX:];
}

fn bool processInput(string input) {

    if containsCommand(""extension "", input) == true {
        findAndSetExtension(input);
        fmt.Printf("" set extension to: '%s'\n"", extension_to_find);
        return false;
    }

    if containsCommand(""find "", input) == true {
        findAndSetFind(input);
        fmt.Printf("" set phrase to find to: '%s'\n"", phrase_to_find);
        return false;
    }

    if containsCommand(""replace "", input) == true {
        findAndSetReplace(input);
        fmt.Printf("" set phrase to replace to: '%s'\n"", phrase_to_replace);
        return false;
    }

    switch input {

        case ""run"":
            findAndReplaceFiles();
            return false;

        case ""help"":
            printCommands();
            return false;

        case ""exit"":
            return true;

        default:
            fmt.Println("" Unknown command\n"");
            return false;
    }
}

int filesProcessed = 0
int totalReplacements = 0

fn walkDirectory(string dirPath) {
    
    []os.DirEntry entries, error err = os.ReadDir(dirPath)
    if err != nil {
        fmt.Printf(""Error reading directory %s: %v\n"", dirPath, err)
        return
    }
    
    for _, entry := range entries {

        string fullPath = filepath.Join(dirPath, entry.Name())
        
        if entry.IsDir() {
            // Recursively walk subdirectories
            walkDirectory(fullPath)
        } else {
            // Process file
            processFileSimple(fullPath, entry)
        }
    }
}

fn processFileSimple(string path, os.DirEntry entry) {
    // Check if file has the correct extension

    string entry_lower_case = strings.ToLower(entry.Name())
    string extension_with_dot = "".""+strings.ToLower(extension_to_find)
    bool has_suffix = strings.HasSuffix(entry_lower_case, extension_with_dot)

    if has_suffix == false {
        return
    }
    
    fmt.Printf(""Processing file: %s\n"", entry.Name())
    
    // Read file content
    []byte content, error err = os.ReadFile(path)
    if err != nil {
        fmt.Printf(""Error reading file %s: %v\n"", path, err)
        return
    }
    
    // Convert to string and count occurrences
    string fileContent = string(content)
    int occurrences = strings.Count(fileContent, phrase_to_find)
    
    if occurrences > 0 {

        string newContent = strings.ReplaceAll(fileContent, phrase_to_find, phrase_to_replace)
        
        // Get file info for permissions
        os.FileInfo info, err = entry.Info()
        if err != nil {
            fmt.Printf(""Error getting file info for %s: %v\n"", path, err)
            return
        }
        
        // Write back to file
        err = os.WriteFile(path, []byte(newContent), info.Mode())
        if err != nil {
            fmt.Printf(""Error writing file %s: %v\n"", path, err)
            return
        }
        
        fmt.Printf(""  - Replaced %d occurrences in %s\n"", occurrences, entry.Name())
        totalReplacements += occurrences
        filesProcessed += 1
        
    } else {
        fmt.Printf(""  - No occurrences found in %s\n"", entry.Name())
    }
}

func findAndReplaceFiles() {
    // Check if all required values are set
    if extension_to_find == ""NA"" || extension_to_find == """" {
        fmt.Println(""Error: Extension not set. Use 'extension <ext>' command first."")
        return
    }
    if phrase_to_find == ""NA"" || phrase_to_find == """" {
        fmt.Println(""Error: Phrase to find not set. Use 'find <phrase>' command first."")
        return
    }
    if phrase_to_replace == ""NA"" || phrase_to_replace == """" {
        fmt.Println(""Error: Replacement phrase not set. Use 'replace <phrase>' command first."")
        return
    }
    
    fmt.Printf(""Looking for .%s files containing '%s' to replace with '%s'\n"",
        extension_to_find, phrase_to_find, phrase_to_replace)
    
    // Get current directory
    string currentDir, error err = os.Getwd()
    if err != nil {
        fmt.Println(""Error getting current directory:"", err)
        return
    }
    
    // Reset counters
    filesProcessed = 0
    totalReplacements = 0
    
    // Walk through all files in current directory and subdirectories
    walkDirectory(currentDir)
    
    fmt.Printf(""\nOperation completed!\n"")
    fmt.Printf(""Files processed: %d\n"", filesProcessed)
    fmt.Printf(""Total replacements made: %d\n"", totalReplacements)
}

fn consoleLoop() {

    *bufio.Scanner scanner = bufio.NewScanner(os.Stdin)

    for {

        fmt.Println(""Write command:"")
        if scanner.Scan() {

            string input = scanner.Text()
            bool should_exit = processInput(input)
            if should_exit == true {
                break
            }
        }
        
        error err = scanner.Err()

        if err != nil {
            fmt.Println(""Error reading input:"", err)
            break
        }
    }
}

fn main() {

    fmt.Println(""Code Renamer\nType 'help' for commands\n"");
    consoleLoop();
}



";

            Fmt.Println($"Inputted code:\n{CODE2}");
            try {

                string? goCode = ConvertToGo(CODE2, "none");
                if (goCode == null) {
                    return;
                }

                Fmt.PrintlnColor($"Done", ConsoleColor.DarkCyan);

                Fmt.Println($"\nConverted Go code: \n{goCode}");
            }
            catch (Exception e) {
                   Fmt.Println($"Internal error: {e.GetType()}");
            }
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