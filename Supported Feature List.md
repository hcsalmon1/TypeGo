**Features:**

**Variables**

  Basic var types ✅  
  Maps ✅  
  Channels ✅  
  local var interfaces ✅  
  local var Interfaces with methods ❌  
  Same type multiple declaration ❌  
  ```go
      int a, b
  ```
  Multiple value declarations - eg: int value, bool ok = getValue() ✅  

**Functions:**

  Basic functions ✅   
  Parameters ✅  
  Multiple return types ✅  
  Multiple parameters with one type ❌  
  ```go
     fn foo(int a, b, c)
  ```
  Variadic parameters ❌  
  Anonymous Functions ❌
  Closures ❌ 

**Control flow**

  For loops ✅  
  If statements ✅  
  Switch statements ✅  
  goto ❌  
  go routines ✅

**Interfaces**

  Interface declarations ✅  
  Interface methods ✅  
  Anonymous interface declaration ✅  
  Anonymous interface declaration with methods ❌  

**Built-in Functions**

append - simplified syntax ✅  
```go
    array.append(1)
```
make - simplified syntax ✅  
```go
    chan int ch = make()
    []int slice = make(1)
```
copy ❌  
len() ✅  
cap() ❌  
delete() ❌  
panic() ❌  
recover() ❌  
new() ❌  


