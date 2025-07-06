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
I will try to fix this in future releases.

**Spacing Problems:**

The spacing isn't always perfect. I'll try to adjust it in future.
For now a workaround is to run the command: go fmt ./...
Just run that in the directory and it will format all the files in a Go style.
