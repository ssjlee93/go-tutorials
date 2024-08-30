# Create a Go Module

[Create a Go module](https://go.dev/doc/tutorial/create-module)

## Start a module that others can use¶
> * In a module, you collect one or more related packages for a discrete and useful set of functions  
> * Go code is grouped into packages, and packages are grouped into modules. 
> * Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires. 

Dependency Management  

> * As you add or improve functionality in your module, you publish new versions of the module

> * In Go, a function whose name starts with a capital letter can be called by a function not in the same package.  

similar to `public` in Java


### Gemini : is go modules the same as Maven or Gradle?
> Gemini :  
> Key Features of Go Modules:
> 
> * go.mod file: Defines the module path, version, and dependencies.
   
> * go.sum file: Stores checksums of downloaded dependencies for integrity verification.   
> * Versioning: Supports semantic versioning for dependencies.   
> * Dependency resolution: Uses a relatively simple algorithm compared to Maven and Gradle.
> * Integration with version control: Leverages Git for dependency management.

Go modules = dependency management  
Previously used `GOPATH`  
Asked Google Gemini to confirm.  
Go modules : dependency management  
Maven / Gradle : dependency management, central repository, compilation, and testing

## Dependency Management
```shell
go mod init <repository-location>
```

> The go mod init command’s argument is your module’s module path. 
> If possible, the module path should be the repository location of your source code.

Typically, `repository-location` would be `github.com/<project-name>`  
Two reserved prefixes are : `test/*` and `example/*` or `example.com/*`  

