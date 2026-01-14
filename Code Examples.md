**Functions**

```go

    fn int add(int a, int b) {
        return a + b
    }

    fn main() {
        int result = add(10, 10)
    }

    fn (int, bool) getValue() {
        return 10, true
    }

```

'fn' or 'func' are both allowed

**Variables:**

```go

    int number = 10
    string message = "Hello World!"
    bool condition = false
    const MAX_INDEX = 100
    const float32 PI = 3.14159

    *int pointer = &number
    *pointer = 11

    //:= only allowed in certain situations
    for i := 0; i < 10; i++ { 

    }

```
The type is optional in constants.  
':=' is only allowed in for loops, if statements and switch statements.

```go

    []int numbers = make(0)
    numbers.append(10)
	numbers = append(numbers, 20) //.append() is optional

    map[string]int scores = {
        "alice": 100,
        "bob": 85,
    }

    chan int channel = make()

    interface{} object = "Hello"

```

To avoid repetition, the type is not needed to be written after the equals or in make().  
The type is only require with 'make()' when it's never written:

```go

	struct MessageHolder {
		[]string Messages
	}

	fn main() {

		MessageHolder holder = {
			Messages: make([]string, 0) //has to be written here
		}
	}

```

Errors:

```go

	fn openFile(string path) error {

        []byte data, error err := os.ReadFile(path)
    	errreturn err
		*os.File file, err = os.Open(path)
		errcheck {
			return fmt.Errorf("Error opening %s, %s", path, err.Err())
			return err
		}
		return nil
    }

```
If you want to write:  
```go
	if err != nil {
		return err
	}
```
In TypeGo you can just write:
```go
	errreturn err
```
They are completely identical.  
You can also put any return value there:  

```go

	errreturn 0, ""
	errreturn DefaultUser()
	erreturn 0, err

```
'errcheck' is used similarly, but to create a block instead.  

Structs:

```go

    struct Game {
        int player_count
        bool game_started
        IntGameResult game_result

        fn StartGame(int _player_count) {
            self.game_started = true
            self.player_count = _player_count
        }
        fn EndGame() {
            self.game_started = false
            self.player_count = 0
        }
    }

    Game game = {
        player_count: 0,
        game_started: false,
        game_result: GameResult.None,
    }

```

Methods are written inside the struct.  
Again the Struct name can be omitted after '='  
Structs object names will always be self.  

**Enums:**

```go

    //Option 1
    enum OsType {
        Windows
        Mac
        Linux
    }

    fn main() {
        OsType os = OsTypeWindows
    }

    //Option 2
    enumstruct OsType {
        Windows
        Mac
        Linux
        fn bool IsWindows() {
            return self == OsType.Windows
        }
    }

    fn main() {
        IntOsType os_type = OsType.Windows
        fmt.Println("Os type:", os_type.ToString())
    }

```
'enum' creates an iota based integer.  
'enumstruct' creates a struct with those members as integers.  
Enums can also have methods. The self parameter will be passed by value and not pointer.  
ToString() methods are generated automatically.

```go

	enumstruct UserType : UType {
		Regular, Admin, Guest
	}

	fn main() {

		UType u_type = UserType.Regular
		//...
	}

```

The default alias for enumstructs is Int+Name. For example:
enumstruct OsType, alias = IntOsType.  
To specify the alias use:' : Name '
 
**Interfaces**

```go

    interface Logger {
        string Log()
    }

    struct FileLogger {
        fn string Log() {
            //implementation
        }
    }

    fn LogMessage(Logger logger) {
        logger.Log()
    }

```

**Unchanged List:**  
-Switches  
-For loops  
-If/else statements  
-Imports  
-Defer  
-Go routines  
-Goto  

Basically anything that isn't a declaration of some kind is the same.
