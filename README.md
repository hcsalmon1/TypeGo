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

Q. What is TypeGo?

  A. TypeGo is simply a transpiler. I got frustrated with certain aspects of the Go language and decided to make this tool to fix those shortcomings.
  It converts .tgo into duplicated .go files.
  Features:
  -No inferred types outside of ifs and fors
  -Return to C style: type name = value
  -Improved enums with two options
  -Improved method syntax

Q. Why did you make TypeGo?

  A. From my experience in the language, Go has 3 main problems:
  1. Inferred types 99% of the time that punishes specifying types or makes it impossible
  2. Enums are terrible
  3. Many things are unintuitive and verbose: append, struct methods, make.
  The goal was to try and fix these 3 things.

Q. How does TypeGo improve inferred types and why are they bad exactly?

A. You can't read any Go code that doesn't include ':=', anywhere. It is basically everywhere and considered idiomatic.
Just let the compiler work out the type. What could go wrong, right?
This is literally the worst feature in Go and I would remove it entirely if I could. I'll explain why:
1. It makes code very hard to read
Example:

```
func main() {
	data := getData()
	fs := getFuncs()

	ch := extractChan(data)

	go func() {
		ch <- map[int][]string{
			1: {"hello", "world"},
			2: {},
		}
	}()

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

If you think that ':=' doesn't make code harder to read than try writing a whole Go project in notepad and then see how it is to read and write.
You basically need to use an IDE to make anything readable at all.

