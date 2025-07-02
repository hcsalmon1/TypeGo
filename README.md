# TypeGo

A tool to help 

Example:

TypeGo:

  package main

  import "fmt"
  
  struct Person {
      string Name
      int Age
      func string Greet() {
          return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", Name, Age)
      }
  }
    
  func main() {
      Person p = {Name: "Alice", Age: 30}
      fmt.Println(p.Greet())
  }

Converts to Go:

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
