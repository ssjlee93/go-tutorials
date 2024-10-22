# go-tutorials
Go programming language tutorials
[go.dev tutorials](https://go.dev/doc/tutorial/)

## package structure 
in the order of creation

1. [Getting started](./getting-started)
2. [Create a module](./create-module)
3. [Tour of Go](./tour)
   1. [Basics](./tour/basics)
   2. [Flow Control](./tour/flowcontrol)
      * [Blog : Defer, Panic, Recover](./defer-panic-and-recover)
   3. [More Types](./tour/more-types)
      * [Blog : Go Slices: usage and internals](./slices-intro)
   4. [Methods](./tour/methods)
   5. [Generics](./tour/generics)
   6. [Concurrency](./tour/concurrency)
4. [Getting started with multi-module workspaces](./workspaces)
5. [Accessing relational database](./database-access)
6. [Developing a RESTful API with Go and Gin](./web-service-gin)
7. [Getting started with generics](./generics)
8. 

## Notes

### packages
ideally, in a legitimate application 1 project would have 1 `go.mod`  
similar to Java, you would have 1 entry point `main.go` and import packages to your necessary classes / go files.  

Some observations I found :  
* each package does not need a separate `go.mod` file.  
* each package may contain unique variables and functions.  
  * unlike Java where each class is independent, go files are not independent; go packages are independent.

### multiple mains
GoLand IDE will not allow multiple `package main` lines in the project.  
However, if you still place it and run the following, each file with `main` function can run.  
```shell
go run <path-to-my-file-with-main>
```
 