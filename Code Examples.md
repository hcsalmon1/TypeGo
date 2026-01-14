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

    map[string]int scores = {
        "alice": 100,
        "bob": 85,
    }

    chan int channel = make()

    interface{} object = "Hello"

```

To avoid repetition, the type is not needed to be written after the equals or in make().

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

**Make and Append**:

```

    []int slice = make(0)
    slice.append(1)
    slice.append(2)
    slice.append(3)
    slice = append(slice, 4)

    for i := 0; i < 3; i++ {
		fmt.Println(slice[0])
	}

```
You only need to write the type in 'make()' when it isn't written at all.  
'.append()' shortcut is also completely optional and the long form is still possible.  

**Unchanged List:**  
-Switches  
-For loops  
-If/else statements  
-Imports  
-Defer  
-Go routines  
-Goto  

Basically anything that isn't a declaration of some kind is the same.
