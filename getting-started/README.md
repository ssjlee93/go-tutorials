# Get Started with Go
Notes for the tutorial  
[Get Started with Go](https://go.dev/doc/tutorial/getting-started)

## Write some code
3. Enable dependency tracking for your code.  
module = dependency management system  
```shell
$ go mod init example/hello
```
this command creates `go.mod` file  
> In actual development, the module path will typically be the repository location where your source code will be kept.   

similar to `pom.xml` or `package.json`

## Call code in an external package¶

find packages from `pkg.go.dev`site  
typically, available functions will be in `Documentation`-> `ìndex` page  