# TypeGo

A tool to help improve and give Go some quality of life improvements.
It converts to Go code, similar to Typescript and Javascript.

Example:

TypeGo:
```
  package main

  import "fmt"
  
  struct Person {
      string Name
      int Age
      fn string Greet() {
          return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", Name, Age)
      }
  }
    
  fn main() {
      Person p = {Name: "Alice", Age: 30}
      fmt.Println(p.Greet())
  }
```

Converts to Go:

```
  package main
  
  import "fmt"
  
  type Person struct {
          Name string
          Age int
  }
  
  func (p *Person) Greet() string {
          return fmt.Sprintf("Hi, I'm %s and I'm %d years old." , p.Name , p.Age)
  }
  
  func main() {
          var p Person = Person{ Name : "Alice" , Age : 30 }
          fmt.Println(p.Greet())
  }
```

Questions and Answers:

**Q. What is TypeGo?**

  A. TypeGo is simply a transpiler. It converts .tgo into duplicated .go files.  
  Features:  
  -No inferred types outside of ifs and fors  
  -Return to C style: type name = value  
  -Improved enums with two options  
  -Improved method syntax  

**Q. Why did you make TypeGo?**

  A. From my experience in the language, I found Go has 3 main problems:
  1. Inferred types 99% of the time and the language punishes specifying types or makes it impossible
  2. Enums are terrible
  3. Many things are unintuitive and verbose: append, struct methods, make.  
  The goal was to try and fix these 3 things.

**Q. What's wrong with ':=' and inferred types?**

A. You can't find any Go code that doesn't include ':=', anywhere, except if you read my Go code. It is basically everywhere and considered idiomatic.
Just let the compiler work out the type. What could go wrong, right?
This is literally the worst feature in Go and I would remove it entirely if I could. I'll explain why:

**1. It makes code very hard to read**
Example:

```
func main() {
	data := getData()
	fs := getFuncs()

	ch := extractChan(data)

        createGoRoutine(ch)

	time.Sleep(100 * time.Millisecond)

	select {
		case msg := <-ch:
			for k, v := range msg {
				fmt.Printf("Key: %d, Value: %v\n", k, v)
			}
		default:
			fmt.Println("No message")
	}

	for _, f := range fs {
		fmt.Println("Func result:", f())
	}
}

```

Good luck reading this without an IDE. Compare this to this version:

```
func main() {
	var data []interface{} = getData()
	var fs []func() int = getFuncs()

	var ch chan map[int][]string = extractChan(data)

        createGoRoutine(ch)

	time.Sleep(100 * time.Millisecond)

	select {
                case msg := <-ch:
                        for k, v := range msg {
                                fmt.Printf("Key: %d, Value: %v\n", k, v)
                        }
                default:
                        fmt.Println("No message")
	}

	for _, f := range fs {
                var f func() int = f
		fmt.Println("Func result:", f())
	}
}

```

Now you don't have to hold these var types in your head or use the crutch of an IDE just to be able to actually understand the code you write.

If you think that ':=' doesn't make code harder to read, then try writing a whole Go project in notepad and then see how easy it is to read and write.
You basically need to use an IDE to make anything readable at all.

**2. It creates silent bugs**

Let's say you make a mistake and forget to dereference a pointer:

```
	func foo(index *int) {
		...
		index_before := index
		...
	}
```

You didn't want 'index_before' to be a pointer but guess what the compiler just inferred the type as.
This is the main problem with inferred types. You are literally putting all your hopes in the compiler guessing what you want.
You get minimal compile time checks when you code this way. No warning, no hint you made a mistake. The compiler just blasts ahead thinking you wanted a pointer.

Compare this to:

```
	func foo(index *int) {
		...
		var index_before int := index  //error
		...
	}
```

You now get a compiler error and it tells you that the type is incompatible.
Summary:

```
	index_before := index   	𝕏 - zero compile time checks, infers a pointer
	var index_before = index 	𝕏 - zero compile time checks, infers a pointer
	var index_before int = index	✓ - compiler actually catches the error
	var index_before int = *index	✓ - compiler checks and finds no error

```

So the default and idiomatic style in Go makes code harder to read and introduces silent bugs? This is good how?  
Yes, it's more concise to write ':=' but this comes at the cost of the compiler not actually checking your code and just guessing what you want.

**Q. How does TypeGo change inferred types?**

Firstly TypeGo reverses the order of declarations back to C style.

TypeGo:
``` 
	index := getIndex()  		𝕏 - not allowed
	var index int = getIndex() 	𝕏 - not allowed
	int index = getIndex() 		✓ - Correct

converts to:

	var index int = getIndex()
```
The converted Go code will never use ':=' with two exceptions. This is to actually have the compiler check your code, rather than guess what you are thinking.

The exceptions are the situations where it is impossible or not needed.

**Constants**
```
	const PI = 3.14159
```

It's obvious this is a float, so you don't need to specify here. You can if you want though:

```
	const float32 PI = 3.14159

converts to:

	const PI float32 = 3.14159
```
The other exceptions are in if statements and for loops. Why? Because Go doesn't allow you to use the long form in those.
```
	for var i int = 0; i < 10; i++ { //error
	for i := 0; i < 10; i++ { //Only way
```
So, as you literally can't specify the type in an if statement or for loop, Go and TypeGo use the same syntax here.
In theory I could do this in the conversion to Go:
```
	TypeGo:
	for int i = 0; i < 10; i++ {

	Converted Go:
	var i int
	for i = 0; i < 10; i++ {
```
The problem with doing this is that 'i' is now not just local to the scope of the for loop and could cause conflicts. TypeGo doesn't track declared variables at all, 
so for simplicity, for loops and if statements are unchanged.

**Q. How does TypeGo improve enums?**

Let's start with the problems of Enums.

```
	type UserType int

	const (
		Regular = UserType iota
		Guest
		Admin
	)
```

1. They have weird syntax
2. They are global constants that conflict with other globals
3. You can't print their names without manually creating string arrays or using 'stringer'
4. There is no value checking and you can set any value

TypeGo tries to address these issues by giving you two options.

**1. enum**

TypeGo:
```
	enum UserType {
		Regular
		Guest
		Admin
	}

```
This will convert to:
```
	type UserType int

	const (
		UserTypeRegular = UserType iota
		UserTypeGuest
		UserTypeAdmin
	)

	func UserTypeToString(user_type int) string {
		switch user_type {
    			case UserTypeRegular:
        			return "Regular"
    			case UserTypeGuest:
        			return "Guest"
    			case UserTypeAdmin:
        			return "Admin"
    			default:
        			return "Unknown"
    		}
	}

```

The name of the enum will be placed before each name to avoid conflicts and a function called NameToString will be auto generated.  
This is the first option, to keep the iota way still possible.

**2. enumstruct**

The other option is 'enumstruct', taken as an idea from 'enumclass' in C++.

```
	enumstruct UserType {
		Regular
		Guest
		Admin
	}
```

This will generate the following Go code:

```
	var UserType = struct {
	        Regular int
	        Guest int
	        Admin int
	}{
	        Regular: 0,
	        Guest: 1,
	        Admin: 2,
	}

	func UserTypeToString(user_type int) string {
		switch user_type {
    			case UserType.Regular:
        			return "Regular"
    			case UserType.Guest:
        			return "Guest"
    			case UserType.Admin:
        			return "Admin"
    			default:
        			return "Unknown"
    		}
	}
```

You then access these values like so:

```

	func main() {
		int user_type = UserType.Regular
		fmt.Println("user_type:", UserTypeToString(user_type));
	}

```

So enumstructs will never conflict with any other global variable.

The one downside is that you can't give the type a custom name, but that doesn't even do anything anyway and is only for the programmer.
All custom enum types are just ints anyway and there is no difference when the compiler processes them.

Issue with enums that can't be fixed:
**Incompatible values being set to enums**

```
	int user_type = 99999999
	or:
	var user_type UserType = 9999999
```

There's nothing that can stop this, as the Go compiler doesn't bother checking these things and it's too complicated to make for me.

But either way, this is a big quality of life improvement for enums. 
