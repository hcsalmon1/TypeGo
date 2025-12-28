
package main

import (
"fmt"
."TypeGo/core"
."TypeGo/parse"
."TypeGo/formatting"
"TypeGo/converting"
"os"
"golang.org/x/sys/windows"
"strings"
"path/filepath"
)

const HELP = "help"
const VERSION = "version"
const CONVERT_DIRECTORY = "convertdir"
const CONVERT_FILE = "convertfile"
const CONVERT_FILE_ABS = "convertfileabs"

func enableVirtualTerminalProcessing() {
	var stdout windows.Handle  = windows.Handle(os.Stdout.Fd()); 
	var mode uint32
	windows.GetConsoleMode(stdout, &mode); 
	mode = mode | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING; 
	windows.SetConsoleMode(stdout, mode); 
	
}

func convertToGo(code string, success *bool, file_name string) string {
	
	fmt.Printf("\tConverting %s\t", file_name)
	
	var parse_result IntParseResult  = ParseResult.Ok
	
	var token_list []Token  = ParseToTokens( &parse_result, code)
	
	if parse_result != ParseResult.Ok {
		fmt.Println("Error:", parse_result.ToString()); 
		* success = false
		return ""
	}
	//PrintTokenList(token_list)
	
	var format_result IntFormatResult  = FormatResult.Ok
	
	var code_format CodeFormat  = FormatCode(token_list, &format_result, code)
	
	if format_result != FormatResult.Ok {
		fmt.Println("Error:", format_result.ToString()); 
		* success = false
		return ""
	}
	
	var convert_result IntConvertResult  = ConvertResult.Ok
	
	var generated_code string  = converting.ConvertToGo( &code_format, &convert_result, code)
	
	if convert_result != ConvertResult.Ok {
		fmt.Println("Error:", convert_result.ToString()); 
		* success = false
		return ""
	}
	
	* success = true
	return generated_code
}

func testingInput() {
	var code string  = `

enumstruct ConvertResult {
    Ok,
    Missing_Expected_Type,
    Unexpected_Type,
    Unexpected_End_Of_File,
    No_Token_In_Node,
    Null_Token,
    Invalid_Node_Type,
    Unsupported_Type,
    Internal_Error,

	fn bool IsError() {
		return self != ConvertResult.Ok
	}
}

enumstruct ConvertResult : IntConvertResult {
    Ok,
    Missing_Expected_Type,
    Unexpected_Type,
    Unexpected_End_Of_File,
    No_Token_In_Node,
    Null_Token,
    Invalid_Node_Type,
    Unsupported_Type,
    Internal_Error,

	fn bool IsError() {
		return self != ConvertResult.Ok
	}
}

fn main() {

}

`
	
	
	fmt.Println("Imported code:\n", code); 
	
	var success bool  = true
	
	var go_code string  = convertToGo(code, &success, "None")
	
	if success == false {
		return 
	}
	
	fmt.Println("Done")
	fmt.Println("Converted Go code: \n", go_code)
	
	var err error  = os.WriteFile("output.txt", []byte(go_code), 0644)
	
	if err != nil {
	fmt.Println("Error writing file:", err)
	return 
	}
	
	fmt.Println("Output written to output.txt")
	
}

func ConvertDirectory(currentDirectory string) {
	
	var entries []os.DirEntry
	var err error
	entries, err = os.ReadDir(currentDirectory)
	if err != nil {
		fmt.Println("Error: ", err)
		return 
	}
	
	for _, entry := range entries {
	
		var fullPath string  = filepath.Join(currentDirectory, entry.Name())
		
		
		if entry.IsDir() {
			// Recursively walk subdirectories
			ConvertDirectory(fullPath)
			
			
		} else  {
		// Process file
		ConvertAndWriteFile(fullPath, entry.Name())
		
		
		}
		
		
		
	}
	
	
}

func ConvertAndWriteFile(tgoFilePath string, file_name string) bool {
	
	var has_suffix bool  = strings.HasSuffix(tgoFilePath, ".tgo")
	
	
	if has_suffix == false {
		return false
	}
	
	var tgoCodeBytes []byte
	var readError error
	tgoCodeBytes, readError = os.ReadFile(tgoFilePath)
	if readError != nil {
		return false
	}
	
	var tgoCode string  = string(tgoCodeBytes)
	
	
	var goCode string
	
	var success bool
	if file_name == "" {
		goCode = convertToGo(tgoCode, &success, "unknown")
		
	} else  {
	goCode = convertToGo(tgoCode, &success, file_name)
	
	}
	
	if success == false {
		return false
	}
	
	if goCode == "" {
		return false
	}
	
	var goPath string  = strings.TrimSuffix(tgoFilePath, filepath.Ext(tgoFilePath))+".go"
	
	
	var writeError error  = os.WriteFile(goPath, []byte(goCode), 0644)
	
	if writeError != nil {
		return false
	}
	
	fmt.Printf("%sDone%s\n", CYAN_TEXT, RESET_TEXT)
	return true
}

func ConvertFile(filePath string) {
	
	var has_suffix bool  = strings.HasSuffix(filePath, ".tgo")
	
	
	if has_suffix == false {
		return 
	}
	
	var fileInfo os.FileInfo
	var err error
	fileInfo, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: File not found:", filePath)
			return 
		}
		
		// Other filesystem error
		fmt.Println("Error accessing file:", filePath)
		return 
	}
	
	if fileInfo.IsDir() {
		fmt.Println("Error: Path is a directory:", filePath)
		return 
	}
	
	ConvertAndWriteFile(filePath, filePath)
	
}

func ConsoleInput(args []string) {
	
	if len(args) == 0 {
		ShowHelp()
		return 
	}
	
	var command string  = strings.ToLower(args[0])
	
	
	var current_working_directory string
	var err error
	var file_path string
	
	switch command {
	
	case HELP:
	ShowHelp()
	
	case VERSION:
	fmt.Println("TypeGo version 0.4")
	
	case CONVERT_FILE:
	if len(args) < 2 {
	fmt.Println("Error: You must specify a filename (e.g. typego convertfile myfile.tgo)")
	break 
	
	}
	
	current_working_directory, err = os.Getwd()
	if err != nil {
	fmt.Println("Error getting current directory:", err)
	return 
	
	}
	
	file_path = filepath.Join(current_working_directory, args[1])
	ConvertFile(file_path)
	
	case CONVERT_FILE_ABS:
	ConvertFile(args[1])
	
	case CONVERT_DIRECTORY:
	
	current_working_directory, err = os.Getwd()
	if err != nil {
	fmt.Println("Error getting current directory:", err)
	return 
	
	}
	ConvertDirectory(current_working_directory)
	fmt.Println("Done")
	
	default:
	fmt.Printf("Unknown command: %s\n", command)
	ShowHelp()
	
	}
	
	
}

func ShowHelp() {
	fmt.Println("TypeGo"); 
	fmt.Println("Usage:"); 
	fmt.Println("  tgo help                 Show this help message"); 
	fmt.Println("  tgo version              Show the version of TypeGo"); 
	fmt.Println("  tgo convertfile file.tgo   Convert a single .tgo file in the current directory"); 
	fmt.Println("  tgo convertdir           Convert all .tgo files in the current directory (recursively)"); 
	fmt.Println("  tgo convertfileabs file.tgo - Convert a single .tgo file an absolute directory")
	
}

func main() {
	
	enableVirtualTerminalProcessing()
	//testingInput()
	var args []string  = os.Args[1:]
	
	ConsoleInput(args)
	
}
