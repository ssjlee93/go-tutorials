# Call your code from another module
[Call your code from another module](https://go.dev/doc/tutorial/call-module-code)

if module is not published, we must resolve it manually  
```shell
go mod edit -replace <module-path>=<location-of-module>
```

Once `replace` directive is added to the `go.mod` file, run :
```shell
go mod tidy
```
This will resolve classpath conflicts  

> To reference a published module, a `go.mod`file would typically omit the `replace` directive and use a `require` directive with a tagged version number at the end.

```go
require example.com/<module-name> v1.0.0
```
Once I make changes and publish go module, Go will automatically use the published version.  
