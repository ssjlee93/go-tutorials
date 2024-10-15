# Getting started with multi-module workspaces
[Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)  

## Create a module for your code¶

```shell
// initialize module
go mod init <module>

// add a dependency
go get <dependency>
```

## Create the workspace¶

```shell
// initialize workspace
go work init <directory>
// create go.work for a workspace containing the modules in the <directory>
```

```go
go 1.23.0

// this line allows go to run ./hello module
use ./hello
```

```shell
// in workspaces
go run ./hello
```
now, we can run the module from the workspace.  

## Download and modify the golang.org/x/example/hello module¶
used `git submodule`  

```shell
git submodule add <repo-url>
```

```go
// use the cloned module
go work use <module-directory>
```
