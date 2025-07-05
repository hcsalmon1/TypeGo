**Known Issues**

**Declaring a struct with a pointer:**

```go
	*User user = {
		Username: "",
		Email: "",
	}
```

If you write this then it will convert incorrect Go code.

Converted code:
```go
	var user *User = *User{
		Username: "",
		Email: "",
	}
```
This should be:
```go
	var user *User = &User{
		Username: "",
		Email: "",
	}
```
This is not trivial at all to fix.  
TypeGo doesn't track declared variables and to work out when something is a declaration of a pointer to a struct and then insert '&' before it
is difficult.

For now you can simply write:
```go
	*User user = &User{
		Username: "",
		Email: "",
	}
```
This will fix the problem but requires you to write the type twice.

**Nested Structs and Methods**

If you nest another struct inside a struct and try to use it with a method, it won't work.

```go

	struct User {
		string UserName
		string Email
	}

	struct Account {
		User

		func PrintUser() {
			fmt.Println("Username: ", UserName, "Email: ", Email);
		}
	}
```

If you try convert this, it will do this:

Converted Go:
```go

	type Account struct {
	        User
	}
	
	func (a *Account) PrintUser() {
	        fmt.Println("Username: ", UserName, "Email: ", Email);
	}

```

The way I detect and insert names before struct members is to simply use a list of varnames from this struct.  
However the struct 'Account' has no variable names. They are all inside 'User', which I don't track at all.

I'm not sure the best way to fix this. The simple workaround is just to generate the methods and then manually type 'a.User.Username' and 'a.User.Email' where needed.
