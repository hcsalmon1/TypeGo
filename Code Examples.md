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

To avoid repetition, the type is not needed to written after the equals or in make().

Structs:

```go

    struct Game {
        int player_count
        bool game_started
        int game_result

        fn StartGame(int _player_count) {
            game_started = true
            player_count = _player_count
        }
        fn EndGame() {
            game_started = false
            player_count = 0
        }
    }

    Game game = {
        player_count: 0,
        game_started: false,
        game_result: GameResult.None,
    }

```

Methods are written in inside the struct.
Again the Struct name can be omitted after '='

**Enums:**

```go

    //Option 1
    enum OsType {
        Windows
        Mac
        Linux
    }

    fn main() {
        OperatingSystem os = OsTypeWindows
    }

    //Option 2
    enumstruct OsType {
        Windows
        Mac
        Linux
    }

    fn main() {
        int os_type = OsType.Windows
    }

```
'enum' creates an iota based integer.  
'enumstruct' creates a struct with those members as integers.  

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

