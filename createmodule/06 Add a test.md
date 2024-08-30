# Add a test
[Add a test](https://go.dev/doc/tutorial/add-a-test)

Go has built-in support for unit testing  
Naming convention is key  

> Test function names have the form `TestName`, where `Name` says something about the specific test. 
> Also, test functions take a pointer to the testing package's `testing.T type` as a parameter. 
> You use this parameter's methods for reporting and logging from your test.

Naming convention denotes which package to use.  
No need to manually specify like Spring Boot annotations.  

```shell
# run tests
go test

# run test verbosely
go test -v

# run test with coverage
go test -cover

# run test verbosely with coverage
go test -v -cover
```

> The go test command executes test functions (Test*) in test files (*_test.go).