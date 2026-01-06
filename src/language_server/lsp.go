
package language_server

import (
"bufio"
"encoding/json"
"fmt"
"io"
"os"
"net/url"
"os/exec"
"path/filepath"
"strconv"
"strings"
"time"
"regexp"
)

var logFile *os.File
type IntMapErrorTag int
var MapErrorTag = struct {
	None IntMapErrorTag
	KeyNotFound IntMapErrorTag
	IncorrectType IntMapErrorTag
	DataIsNull IntMapErrorTag
	}{
	None: 0,
	KeyNotFound: 1,
	IncorrectType: 2,
	DataIsNull: 3,
	}

func (self IntMapErrorTag) ToString() string {
	switch self {
	case MapErrorTag.None:
		return "None"
	case MapErrorTag.KeyNotFound:
		return "KeyNotFound"
	case MapErrorTag.IncorrectType:
		return "IncorrectType"
	case MapErrorTag.DataIsNull:
		return "DataIsNull"
	default:
		return "Unknown"
	}
}


func (self IntMapErrorTag) IsError() bool {
	return self != MapErrorTag.None
}

func (self IntMapErrorTag) Print() {
	fmt.Print("Error:", self.ToString())
	
}

func (self IntMapErrorTag) Println() {
	fmt.Println("Error:", self.ToString())
	
}

type MapError struct {
	ErrorTag IntMapErrorTag
	Detail string
	
}


func (self *MapError) IsError() bool {
	return self.ErrorTag.IsError()
}

func (self *MapError) Print() {
	fmt.Fprint(os.Stderr, "Error:", self.ErrorTag.ToString(), self.Detail)
	
}

func (self *MapError) Println() {
	fmt.Fprintln(os.Stderr, "Error:", self.ErrorTag.ToString(), self.Detail)
	
}

func (self *MapError) Clear() {
	self.ErrorTag = MapErrorTag.None
	self.Detail = ""
	
}

func (self *MapError) Set(tag IntMapErrorTag, detail string) {
	self.ErrorTag = tag
	self.Detail = detail
	
}

func (self *MapError) SetTag(tag IntMapErrorTag) {
	self.ErrorTag = tag
	
}

func (self *MapError) SetDetail(detail string) {
	self.Detail = detail
	
}

func (self *MapError) ToString() string {
	return self.ErrorTag.ToString()+" " +self.Detail
}

type Map struct {
	Data map[string]any
	
}


func (self *Map) Set(key string, value any) {
	self.Data[key] = value
	
}

func (self *Map) ContainsEntry(key string) bool {
	; var ok bool
	_, ok  = self.Data[key]
	
	return ok
}

func (self *Map) Get(key string, map_error *MapError) any {
	
	if self.IsValid() == false {
		map_error.SetTag(MapErrorTag.DataIsNull)
		return Map {Data:nil}
	}
	
	var val any; var ok bool
	val, ok  = self.Data[key]
	
	
	if ok == false {
		map_error.SetTag(MapErrorTag.KeyNotFound)
		return 0
	}
	return val
}

func (self *Map) GetString(key string, map_error *MapError) string {
	
	if self.IsValid() == false {
		map_error.SetTag(MapErrorTag.DataIsNull)
		return ""
	}
	
	var value any; var ok bool
	value, ok  = self.Data[key]
	
	if ok == true {
		var text string; 
		text, ok  = value.(string)
		
		if ok == true {
			return text
		}
		
		map_error.SetTag(MapErrorTag.IncorrectType)
		return ""
	}
	map_error.SetTag(MapErrorTag.KeyNotFound)
	return ""
}

func (self *Map) GetInt(key string, map_error *MapError) int {
	
	if self.IsValid() == false {
		map_error.SetTag(MapErrorTag.DataIsNull)
		return 0
	}
	
	if self.ContainsEntry(key) == false {
		map_error.SetTag(MapErrorTag.KeyNotFound)
		return 0
	}
	var variable any  = self.Get(key, map_error)
	if map_error.IsError() {
		return 0
	}
	
	switch var_type:=variable.(type) {
		
		
		case float64:
			return int (var_type)
			
		case int:
			return var_type
			
	}
	
	
	map_error.SetTag(MapErrorTag.IncorrectType)
	return 0
}

func (self *Map) GetMap(key string, map_error *MapError) Map {
	
	if self.IsValid() == false {
		map_error.SetTag(MapErrorTag.DataIsNull)
		return Map {Data:nil}
	}
	
	var value any  = self.Get(key, map_error)
	if map_error.IsError() {
		return Map {Data:nil}
	}
	
	if raw, ok := value.(map[string]any); ok {
		return Map {Data:raw}
	}
	map_error.SetTag(MapErrorTag.IncorrectType)
	return Map {Data:nil}
}

func (self *Map) IsValid() bool {
	return self.Data != nil
}

func (self *Map) Delete(key string) {
	delete(self.Data, key)
	
}

func (self *Map) Len() int {
	return len(self.Data)
}

func (self *Map) Clear() {
	self.Data = make(map[string]any)
	
}

func (self *Map) ToJSON() string {
	var b []byte; var err error
	b, err  = json.MarshalIndent(self.Data, "", "  ")
	
	if err != nil {
		return "<json marshal error>"
	}
	
	return string(b)
}


func initLogger() {
	
	var path string  = "C:\\Users\\Tom\\Desktop\\typego-lsp.log"
	
	; var err error
	logFile, err  = os.OpenFile(path, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	
	if err != nil {
		return 
	}
	
	log("=== LSP started ===")
	
}

func log(message string) {
	if logFile != nil {
		var now string  = time.Now().Format("15:04:05.000")
		logFile.WriteString("[" +now + "] " +message + "\n")
		
	}
	
}

func NewMapError() MapError {
	return MapError {ErrorTag:MapErrorTag.None, Detail:"", }
}

func NewMap() Map {
	return Map {Data:make(map[string]any), }
}

func NewMapFromData(input map[string]any) Map {
	return Map {Data:input, }
}

func readMessage(reader *bufio.Reader) (string,error) {
	
	var contentLength int
	
	for  {
	
		var line string; var err error
		line, err  = reader.ReadString('\n')
		
		if err != nil {
		return "", err
		}
		
		
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		
		
		if strings.HasPrefix(line, "Content-Length:") {
			var value string  = strings.TrimSpace(strings.TrimPrefix(line, "Content-Length:"))
			contentLength, _ = strconv.Atoi(value)
			continue
		}
			}
	
	
	var buffer []byte  = make([]byte, contentLength)
	; var err error
	_, err  = io.ReadFull(reader, buffer)
	
	if err != nil {
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		// Valid end-of-stream
		return string(buffer), nil
	}
	return "", err
	}
	
	return string(buffer), nil
}

func parseJSON(raw string) Map {
	raw = strings.TrimRight(raw, "\x00\r\n\t ")
	var obj map[string]any
	var err error  = json.Unmarshal([]byte(raw), &obj)
	if err != nil {
	return NewMap()
	}
	
	return NewMapFromData(obj)
}

func sendResponse(id int, result any) {
	var body []byte; 
	body, _  = json.Marshal(map[string]any {
		"jsonrpc":"2.0", 
		"id":id, 
		"result":result, 
	})
	
	
	os.Stdout.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(body), body), )
	
}

func handleInitialize(message_data_map Map, map_error *MapError) {
	var id int  = message_data_map.GetInt("id", map_error)
	if map_error.IsError() {
		map_error.SetDetail("Error getting id, handleInitialize")
		return 
	}
	
	sendResponse(id, map[string]any {"capabilities":map[string]any {"textDocumentSync":map[string]any {"openClose":true, "change":1, "save":map[string]any {"includeText":true, }, }, }, })
	
}

func extractLineCol(error_message string) (int,int) {
	
	log("Error message: " +error_message)
	
	var index int  = strings.Index(error_message, "Error on line ")
	if index == -1 {
		return 0, 1
	}
	
	var remaining_error_message string  = error_message[index + len("Error on line "):]
	var error_parts []string  = strings.SplitN(remaining_error_message, ":", 2)
	if len(error_parts) < 1 {
		return 0, 1
	}
	
	var numbers_as_text []string  = strings.Split(error_parts[0], ",")
	if len(numbers_as_text) != 2 {
		return 0, 1
	}
	
	var line_number int; var err error
	line_number, err  = strconv.Atoi(strings.TrimSpace(numbers_as_text[0]))
	
	if err != nil {
		return 0, 1
	}
	
	var column_number int; 
	column_number, err  = strconv.Atoi(strings.TrimSpace(numbers_as_text[1]))
	
	if err != nil {
		return 0, 1
	}
	
	return line_number, column_number
}

func sendDiagnostics(uri string, message string) {
	
	// Clear diagnostics
	if message == "" {
		var payload []byte; 
		payload, _  = json.Marshal(map[string]any {
			"jsonrpc":"2.0", 
			"method":"textDocument/publishDiagnostics", 
			"params":map[string]any {
				"uri":uri, 
				"diagnostics":[]any {}, 
			}, 
		})
		
		
		os.Stdout.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(payload), payload), )
		return 
	}
	
	var line_number int; var column_number int
	line_number, column_number  = extractLineCol(message)
	
	
	var clean string  = stripAnsi(string(message))
	var cleaned_message string  = extractUserError(clean)
	
	var diag map[string]any  = map[string]any{
		"uri":uri, 
		"diagnostics":[]any {
			map[string]any {
				"range":map[string]any {
					"start":map[string]int {
						"line":line_number, 
						"character":column_number, 
					}, 
					"end":map[string]int {
						"line":line_number, 
						"character":column_number + 1, 
					}, 
				}, 
				"severity":1, 
				"source":"typego", 
				"message":cleaned_message, 
			}, 
		}, 
	}
	
	var payload []byte; 
	payload, _  = json.Marshal(map[string]any {
		"jsonrpc":"2.0", 
		"method":"textDocument/publishDiagnostics", 
		"params":diag, 
	})
	
	
	os.Stdout.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(payload), payload), )
	
}

func stripAnsi(input string) string {
	// matches ESC[ ... m
	var regex *regexp.Regexp; var err error
	regex, err  = regexp.Compile("\x1b\\[[0-9;]*m")
	
	if err != nil {
		return ""
	}
	
	return regex.ReplaceAllString(input, "")
}

func extractUserError(output string) string {
	var lines []string  = strings.Split(output, "\n")
	
	for i := 0; i < len(lines); i++ {
	
		var line string  = lines[i]
		
		if strings.Contains(line, "Error on line") {
			var index int  = strings.Index(line, ":")
			if index != -1 {
				return strings.TrimSpace(line[index + 1:])
			}
			
		}
			}
	
	
	return "Compilation failed"
}

func handleDidSave(json_data_map Map, map_error *MapError) {
	
	var parameters Map  = json_data_map.GetMap("params", map_error)
	if map_error.IsError() {
		map_error.Detail = "Error getting parameters from map"
		return 
	}
	var doc Map  = parameters.GetMap("textDocument", map_error)
	if map_error.IsError() {
		map_error.Detail = "Error getting textDocument from map"
		return 
	}
	var uri string  = doc.GetString("uri", map_error)
	if map_error.IsError() {
		map_error.Detail = "Error getting uri from map"
		return 
	}
	
	var path string  = strings.TrimPrefix(uri, "file://")
	
	var decoded_path string; var err error
	decoded_path, err  = url.PathUnescape(path)
	
	if err != nil {
	log("failed to decode uri: " +err.Error())
	return 
	}
	
	
	// Windows: strip leading slash from /C:/...
	if len(decoded_path) > 2 && decoded_path[0] == '/' && decoded_path[2] == ':' {
		decoded_path = decoded_path[1:]
		
	}
	
	var final_path string  = filepath.FromSlash(decoded_path)
	
	var cmd *exec.Cmd  = exec.Command("tgo", "convertfileabs", final_path)
	var output []byte; 
	output, err  = cmd.CombinedOutput()
	
	
	if err != nil {
	log("sending diagnostics with:" +string(output))
	sendDiagnostics(uri, string(output))
	return 
	}
	
	
	// clear diagnostics on success
	log("sending empty diagnostics")
	sendDiagnostics(uri, "")
	
}

func Run() {
	
	initLogger()
	log("Started")
	
	var reader *bufio.Reader  = bufio.NewReader(os.Stdin)
	
	var map_error MapError  = NewMapError()
	var shuttingDown bool  = false
	
	for  {
	
		map_error.Clear()
		
		var raw string; var err error
		raw, err  = readMessage(reader)
		
		if err != nil {
		if err == io.EOF {
			log("Error: end of file, exit")
			return 
		}
		var error_message string  = fmt.Sprintf("read error: %s", err)
		log(error_message)
		fmt.Fprintln(os.Stderr, "Read error:", err)
		return 
		}
		
		var raw_log string  = fmt.Sprintf("Received: %s", raw)
		log(raw_log)
		
		var json_data_map Map  = parseJSON(raw)
		if json_data_map.Data == nil {
			log("Error: json data is nil")
			fmt.Fprintln(os.Stderr, "Error: json data is nil")
			continue
		}
		
		var method string  = json_data_map.GetString("method", &map_error)
		if map_error.IsError() {
			log(map_error.ToString())
			continue
		}
		
		var id int  = 0
		var hasID bool  = false
		if json_data_map.ContainsEntry("id") {
			id = json_data_map.GetInt("id", &map_error)
			if map_error.IsError() {
				log(map_error.ToString())
				continue
			}
			hasID = true
			
		}
		
		switch method {
			
			
			case "initialized":
				log("Initialized")
				
				
			case "initialize":
				if hasID {
					log("Handle initialize")
					handleInitialize(json_data_map, &map_error)
					
				}
				
				
			case "shutdown":
				shuttingDown = true
				if hasID {
					log("send response")
					sendResponse(id, nil)
					
				}
				
				
			case "exit":
				if shuttingDown == false {
					log("Error exit, code 1")
					os.Exit(1)
					
				}
				log("successful exit, code 0")
				os.Exit(0)
				
				
			case "textDocument/didSave":
				log("Handle did save")
				handleDidSave(json_data_map, &map_error)
				
		}
		
			}
	
	
}
