
package main

import (
"fmt"
."TypeGo/core"
."TypeGo/parse"
."TypeGo/formatting"
"TypeGo/language_server"
"TypeGo/converting"
"os"
 //"golang.org/x/sys/windows"
"strings"
"path/filepath"
)

const HELP = "help"
const VERSION = "version"
const CONVERT_DIRECTORY = "convertdir"
const CONVERT_FILE = "convertfile"
const CONVERT_FILE_ABS = "convertfileabs"
const LSP = "start_lsp"

func convertToGo(code string, success *bool, file_name string) string {
	
	fmt.Printf("\tConverting %s\t", file_name)
	
	var parse_result IntParseResult  = ParseResult.Ok
	var token_list []Token  = ParseToTokens( &parse_result, code)
	if parse_result.IsError() {
		parse_result.Println()
		* success = false
		return ""
	}
	//PrintTokenList(token_list)
	
	var format_result IntFormatResult  = FormatResult.Ok
	var code_format CodeFormat  = FormatCode(token_list, &format_result, code)
	if format_result.IsError() {
		format_result.Println()
		* success = false
		return ""
	}
	//code_format.Print()
	
	var convert_result IntConvertResult  = ConvertResult.Ok
	var generated_code string  = converting.ConvertToGo( &code_format, &convert_result, code)
	if convert_result.IsError() {
		convert_result.Println()
		* success = false
		return ""
	}
	
	* success = true
	return generated_code
}

func testingInput() {
	var code string  = `


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

func ConvertDirectory(currentDirectory string) bool {
	
	var entries []os.DirEntry
	var err error
	entries, err = os.ReadDir(currentDirectory)
	if err != nil {
	fmt.Println("Error: ", err)
	return false
	}
	
	
	for _, entry := range entries {
	
		var fullPath string  = filepath.Join(currentDirectory, entry.Name())
		
		if entry.IsDir() {
			// Recursively walk subdirectories
			var success bool  = ConvertDirectory(fullPath)
			if success == false {
				return false
			}
			
			
		} else  {
			// Process file
			_ = ConvertAndWriteFile(fullPath, entry.Name())
			
		}
		
		
			}
	
	return true
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

func ConvertFile(filePath string) bool {
	
	var has_suffix bool  = strings.HasSuffix(filePath, ".tgo")
	
	if has_suffix == false {
		return false
	}
	
	var fileInfo os.FileInfo
	var err error
	fileInfo, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: File not found:", filePath)
			return false
		}
		
		// Other filesystem error
		fmt.Println("Error accessing file:", filePath)
		return false
	}
	
	if fileInfo.IsDir() {
		fmt.Println("Error: Path is a directory:", filePath)
		return false
	}
	
	return ConvertAndWriteFile(filePath, filePath)
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
	var ok bool  = true
	
	switch command {
		
		
		case HELP:
			ShowHelp()
			
			
		case VERSION:
			fmt.Println("TypeGo version 0.4")
			
			
		case CONVERT_FILE:
			if len(args) < 2 {
				fmt.Println("Error: You must specify a filename (e.g. typego convertfile myfile.tgo)")
				os.Exit(1)
				
			}
			
			current_working_directory, err = os.Getwd()
			if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
			
			}
			
			
			file_path = filepath.Join(current_working_directory, args[1])
			ok = ConvertFile(file_path)
			if ok == false {
				os.Exit(1)
				
			}
			
			
		case CONVERT_FILE_ABS:
			ok = ConvertFile(args[1])
			if ok == false {
				os.Exit(1)
				
			}
			
			
		case CONVERT_DIRECTORY:
			
			current_working_directory, err = os.Getwd()
			if err != nil {
			fmt.Println("Error getting current directory:", err)
			return 
			}
			
			ok = ConvertDirectory(current_working_directory)
			if ok == false {
				fmt.Println("Error, exit 1")
				os.Exit(1)
				
			}
			fmt.Println("Done")
			
			
		case LSP:
			language_server.Run()
			
			
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
	fmt.Println("  tgo convertfileabs file.tgo - Convert a single .tgo file with an absolute directory")
	fmt.Println("  tgo lsp - starts the language server")
	
}

func main() {
	
	//enableVirtualTerminalProcessing()
	//testingInput()
	var args []string  = os.Args[1:]
	ConsoleInput(args)
	
}
