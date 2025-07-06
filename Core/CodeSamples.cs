

namespace TypeGo
{
    public static class CodeSamples
    {
        const string BASIC_STRUCT =
                @"
package main

import ""fmt""

struct Person {
    string Name
    int Age
}

fn main() {
    Person p = {
        Name: ""Alice"",
        Age:  30,
    }

    fmt.Println(""Name:"", p.Name)
    fmt.Println(""Age:"", p.Age)
}

";

		public const string NESTED_STRUCT =
                            @"
package main

import ""fmt""

struct Address {
    string City 
    string State
}

struct Person {
    string Name
    int Age
    Address
}

func main() {
    Person p = {
        Name: ""Alice"",
        Age:  30,
        Address: Address{
            City:  ""New York"",
            State: ""NY"",
        },
    }

    fmt.Println(""Name:"", p.Name)
    fmt.Println(""City:"", p.City)     
    fmt.Println(""State:"", p.State)   
}

";

        public const string STRUCT_AND_METHOD =
@"package main

import ""fmt""

struct User {
    string Name 
}

func (User u) SayHello() {
    fmt.Println(""Hi, my name is"", u.Name)
}

func main() {
    User u = {Name: ""Alice""}
    u.SayHello()
}";

        public const string ERROR_1 =
@"package main

import ""errors""

fn error mayFail() {
    return errors.New(""fail"")
}

func main() {
    error err = mayFail()
    if err != nil {
        return
    }
}";

        public const string FUNCTION_1 =
@"package main

fn int add(int a, int b) {
    return a + b
}";

        public const string ARRAY_SLICE_1 =
@"package main

fn main() {
    []int nums = {1, 2, 3}
    nums.append(4)
}";

        public const string DECLARATION_1 =
@"package main

fn main() {
    int count = 5
    string greeting = ""Hello""
    _ = count
    _ = greeting
}";

        public const string DECLARATION_2 =
@"package main

import ""fmt""

fn main() {
    int number = 10
    string name = ""Alice""
    fmt.Println(number, name)
}";

    

    public const string HARD_EXAMPLE =
            @"
package main

import ""fmt""

fn int sum([]int values) {
    int total = 0
    for _, v := range values {
        total += v
    }
    return total
}

fn []int filterEven([]int values) {
    []int result
    for _, v := range values {
        if v % 2 == 0 {
            result.append(v)
        }
    }
    return result
}

fn main() {
    []int numbers = {1, 2, 3, 4, 5, 6}
    []int evens = filterEven(numbers)
    int total = sum(evens)
    fmt.Println(""Sum of even numbers:"", total)
}";

        public const string FOR_TESTING =
            @"
package main

fn main() {

    for {
        break
    }
    int i = 0
    for i < 10 {
        i += 1
    }
    for int i, int j = 0, 10; i < j; i++ {
        break
    }
}
            ";


        public static class LEETCODE
        {
            public const string TWO_SUM =
                @"
fn []int twoSum([]int nums, int target) {
    map[int]int seen = make()

    for i, num = range nums {
        int complement = target - num
        if idx, ok = seen[complement]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }

    return nil
}";
        }

		public static class OTHER_PEOPLES_CODE
		{
			public const string SOME_WEB_THING =
				@"

fn version(http.ResponseWriter w, *http.Request r) {
	*debug.BuildInfo info, bool ok = debug.ReadBuildInfo()
	if !ok {
		http.Error(w, ""no build information available"", 500)
		return
	}

	fmt.Fprintf(w, ""<!DOCTYPE html>\n<pre>\n"")
	fmt.Fprintf(w, ""%s\n"", html.EscapeString(info.String()))
}

fn greet(http.ResponseWriter w, *http.Request r) {
	string name = strings.Trim(r.URL.Path, ""/"")
	if name == """" {
		name = ""Gopher""
	}

	fmt.Fprintf(w, ""<!DOCTYPE html>\n"")
	fmt.Fprintf(w, ""%s, %s!\n"", *greeting, html.EscapeString(name))
}
";

			public const string WEB_CODE_2 =
				@"package main

import (
    ""encoding/json""
    ""fmt""
    ""net/http""
    ""time""
)

// User represents a simple user structure
struct User {
    int       ID       
    string    Name    
    string    Email    
    time.Time CreatedAt 
}


// Handler for the home route
fn homeHandler(http.ResponseWriter w, *http.Request r) {
    fmt.Fprintln(w, ""Welcome to the Go Web Server!"")
}

// Handler to return JSON response
fn usersHandler(http.ResponseWriter w, *http.Request r) {
    []User users = {
        {ID: 1, Name: ""Alice"", Email: ""alice@example.com"", CreatedAt: time.Now()},
        {ID: 2, Name: ""Bob"", Email: ""bob@example.com"", CreatedAt: time.Now()},
    }

    w.Header().Set(""Content-Type"", ""application/json"")
    json.NewEncoder(w).Encode(users)
}

fn main() {
    *http.ServeMux mux = http.NewServeMux()
    mux.HandleFunc(""/"", homeHandler)
    mux.HandleFunc(""/users"", usersHandler)

    *http.Server server = {
        Addr:    "":8080"",
        Handler: mux,
    }

    fmt.Println(""Starting server on :8080..."")
    error err = server.ListenAndServe();
    if err != nil {
        fmt.Printf(""Server error: %s\n"", err)
    }
}";

			public const string ALGORITHM =
                @"
package main

import (
	""fmt""
)

struct Stack {
	[]int data
}

func Push(*Stack s, int val) {
	s.data = append(s.data, val)
}

func int Pop(*Stack s) {
	if len(s.data) == 0 {
		return -1
	}
	int val = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func int Peek(*Stack s) {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func bool IsEmpty(*Stack s) {
	return len(s.data) == 0
}

func int Fibonacci(int n) {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func int BinarySearch([]int arr, int target) {
	int low = 0
	int high = len(arr) - 1

	for low <= high {
		int mid = (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BubbleSort([]int arr) {
	int n = len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j+1]
				arr[j+1] = arr[j]
			}
		}
	}
}

func int Factorial(int n) {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func PrintMatrix([][]int matrix) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf(""%d "", val)
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(""Fibonacci(10):"", Fibonacci(10))

	Stack stack = {}
	for i := 0; i < 5; i++ {
		Push(&stack, i)
	}
	for !IsEmpty(&stack) {
		fmt.Println(""Popped:"", Pop(&stack))
	}

	[]int arr = {5, 2, 9, 1, 5, 6}
	BubbleSort(arr)
	fmt.Println(""Sorted array:"", arr)

	fmt.Println(""Factorial(5):"", Factorial(5))

	[][]int matrix = {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	PrintMatrix(matrix)

	int index = BinarySearch(arr, 5)
	fmt.Println(""Index of 5 in sorted array:"", index)
}
				";
        }

        public static class MY_CODE
        {
            public const string PARSE =
                @"
package parsing

import (
	""fmt""
	. ""compiler/core""
)

struct ParsingData {
	string code;
	ParseError parse_error;
	[]*Token token_list;
	int line_count;
	int char_count;
	*Token last_token;
	bool was_comment;
}

fn processLetterParse(*ParsingData parsing_data, *int character_index) {
	
	if parsing_data.last_token != nil {

		if parsing_data.last_token.Token_type == TokenType_Comment {
			parsing_data.was_comment = true;
		}
	}

	byte current_char = parsing_data.code[*character_index];

	if (current_char == '\n') {

		if parsing_data.was_comment == true {
			*Token end_comment = NewToken("""", TokenType_EndComment, parsing_data.line_count, parsing_data.char_count);
			parsing_data.token_list = append(parsing_data.token_list, end_comment);
			parsing_data.was_comment = false;
		}
		parsing_data.line_count += 1;
		*character_index += 1;
		return;
	}
	bool is_special_char =
		current_char == '\r' ||
		current_char == '\t' ||
		current_char == ' ' ||
		current_char == '\\';

	if is_special_char == true {
		*character_index += 1;
		return;
	}

	int previous_character_index = *character_index;
	*Token token = getParsedToken(parsing_data, character_index);

	if previous_character_index == *character_index {
		*character_index += 1;
	}
	if parsing_data.parse_error != ParseError_None {
		return;
	}

	AppendSlice(&parsing_data.token_list, token);
	parsing_data.last_token = token;
}

fn []*Token ParseToTokens(*ParseError parse_error, string code) {
	fmt.Print(""\n\t"" + GREY_TEXT + ""Parsing:"" + RESET_TEXT + ""\t\t\t"");

	[]*Token token_list;
	ParsingData parsing_data;
	parsing_data.code = code;
	parsing_data.line_count = 0;
	parsing_data.char_count = 0;
	parsing_data.parse_error = ParseError_None;
	parsing_data.last_token = nil;
	parsing_data.token_list = token_list;
	parsing_data.was_comment = false;

	if (len(code) == 0) {
		*parse_error = ParseError_String_Length_Zero;
		return token_list;
	}

	int STRING_LENGTH = len(code);

	int character_index;
	for character_index = 0; character_index < STRING_LENGTH; {

		processLetterParse(&parsing_data, &character_index);
	}
	fmt.Print(GREEN_TEXT + ""Done\n"" + RESET_TEXT);
	return parsing_data.token_list;
}

fn PotentialType getParsePotentialType(byte c) {
	if (c == ' ') { 
		return PotentialType_Seperator;
	}
	for i := 0; i < len(SEPARATORS); i++ {

		if (c == SEPARATORS[i]) {
			return PotentialType_SeperatorToken;
		}
	}
	if (c == '""') {
		return PotentialType_String;
	}
	if (c == '*') {
		return PotentialType_Multiply;
	}
	return PotentialType_NA;
}

fn *Token getParsedToken(*ParsingData parsing_data, *int current_string_index) {
	
	[250]byte buffer;
	int current_buffer_index = 0;

	if current_buffer_index >= len(parsing_data.code) {
		parsing_data.parse_error = ParseError_Buffer_Index_Over_Max;
		return nil;
	}

	bool is_string = false;

	//fmt.Println(""\tlooping:"");

	for current_buffer_index < 250 {
		if *current_string_index >= len(parsing_data.code) {
			parsing_data.parse_error = ParseError_Current_Index_Out_Of_Range;
			return nil;
		}
		if current_buffer_index >= 250 {
			parsing_data.parse_error = ParseError_Buffer_Index_Over_Max;
			return nil;
		}
		byte current_char = parsing_data.code[*current_string_index];

		//fmt.Printf(""\tcurrent char: %c\n"", current_char);

		PotentialType potential_character_type = getParsePotentialType(current_char); //what type could it be? is it expected?
		if potential_character_type == PotentialType_String {
			if is_string == true {
				potential_character_type = PotentialType_NA;
				is_string = false;
			} else {
				is_string = true;
			}
		}

		CharacterAction action = getParseAction(potential_character_type, is_string); //should we break or continue?
		if potential_character_type == PotentialType_Multiply {

			if current_buffer_index != 0 {
				action = CharacterAction_AddCharAndBreak;
			} else {
				action = CharacterAction_AddChar;
			}
		}

		switch action {
			case CharacterAction_Return:
				return nil;

			case CharacterAction_AddChar:
				buffer[current_buffer_index] = current_char;

			case CharacterAction_AddCharAndBreak:
				if (current_buffer_index == 0) {
					buffer[current_buffer_index] = current_char;
					*current_string_index += 1;
					current_buffer_index += 1;
				}

			case CharacterAction_ReturnError:
				parsing_data.parse_error = ParseError_UnexpectedValue;
				return nil;

			default:

		}

		if action == CharacterAction_Break || action == CharacterAction_AddCharAndBreak {
			break;
		}

		*current_string_index += 1;
		current_buffer_index += 1;
	}
	string tokenText = string(buffer[:current_buffer_index])
	TokenType token_type = getParseTokenType(tokenText);

	return NewToken(tokenText, token_type, parsing_data.line_count, parsing_data.char_count);
}

func TokenType getParseTokenType(string input) {
	// Keywords
	if input == ""fn"" { return TokenType_Fn; }
	if input == ""if"" { return TokenType_If; }
	if input == ""else"" {  return TokenType_Else;  }
	if input == ""for"" { return TokenType_For; }
	if input == ""while"" { return TokenType_While; }
	if input == ""return"" { return TokenType_Return; }
	if input == ""break"" { return TokenType_Break; }
	if input == ""continue"" { return TokenType_Continue; }
	if input == ""print"" { return TokenType_Print; }
	if input == ""println"" { return TokenType_Println; }
	if input == ""true"" { return TokenType_True; }
	if input == ""false"" { return TokenType_False; }
	if input == ""in"" { return TokenType_In; }
	if input == ""defer"" { return TokenType_Defer; }
	if input == ""new"" { return TokenType_New; }

	// Types
	if input == ""u8"" { return  TokenType_u8; }
	if input == ""i8"" { return  TokenType_i8;}
	if input == ""i32"" { return  TokenType_i32;}
	if input == ""f32"" { return  TokenType_f32;}
	if input == ""f64"" { return  TokenType_f64;}
	if input == ""i64"" { return  TokenType_i64;}
	if input == ""u64"" { return  TokenType_u64;}
	if input == ""string"" { return  TokenType_String;}
	if input == ""bool"" { return  TokenType_Bool;}
	if input == ""char"" { return  TokenType_Char;}
	if input == ""void"" { return  TokenType_Void;}
	if input == ""const"" { return  TokenType_Const;}
	if input == ""int"" { return  TokenType_i32;}
	if input == ""usize"" { return  TokenType_Usize;}

	// Operators
	if input == ""+"" { return  TokenType_Plus;}
	if input == ""-"" { return  TokenType_Minus;}
	if input == ""*"" { return  TokenType_Multiply;}
	if input == ""/"" { return  TokenType_Divide;}
	if input == ""="" { return  TokenType_Equals;}
	if input == ""+="" { return  TokenType_PlusEquals;}
	if input == ""-="" { return  TokenType_MinusEquals;}
	if input == ""*="" { return  TokenType_MultiplyEquals;}
	if input == ""/="" { return  TokenType_DivideEquals;}
	if input == "">"" { return  TokenType_GreaterThan;}
	if input == ""<"" { return  TokenType_LessThan;}
	if input == ""=="" { return  TokenType_EqualsEquals;}
	if input == "">="" { return  TokenType_GreaterThanEquals;}
	if input == ""<="" { return  TokenType_LessThanEquals;}
	if input == ""%"" { return  TokenType_Modulus;}
	if input == ""!="" { return  TokenType_NotEquals;}
	if input == ""&"" { return  TokenType_And;}
	if input == ""&&"" { return  TokenType_AndAnd;}
	if input == ""|"" { return  TokenType_Or;}
	if input == ""||"" { return  TokenType_OrOr;}
	if input == ""%="" { return  TokenType_ModulusEquals;}

	if input == ""//"" { return  TokenType_Comment;}
	if input == ""delete"" { return  TokenType_Delete;}

	// Parentheses and Brackets
	if input == ""("" { return  TokenType_LeftParenthesis;}
	if input == "")"" { return  TokenType_RightParenthesis;}
	if input == ""{"" { return  TokenType_LeftBrace;}
	if input == ""}"" { return  TokenType_RightBrace;}
	if input == ""["" { return  TokenType_LeftSquareBracket;}
	if input == ""]"" { return  TokenType_RightSquareBracket;}

	if input == "";"" { return  TokenType_Semicolon;}
	if input == "","" { return  TokenType_Comma;}
	if input == ""."" { return  TokenType_FullStop;}

	// Number literals
	if isIntegerToken(input) == true { return  TokenType_IntegerValue;}
	if isDecimalToken(input) == true { return  TokenType_DecimalValue;}

	// String or Char value
	if input[0] == '""' { return  TokenType_StringValue;}
	if input[0] == '\'' { return  TokenType_CharValue;}

	return  TokenType_PotentialVariable;
}
                ";
        }

    }
}
