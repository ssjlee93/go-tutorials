# Generics

## Type parameters
> Go functions can be written to work on multiple types using type parameters. 
> The type parameters of a function appear between brackets, before the function's arguments.

```go
func Index[T comparable](s []T, x T) int
```

This is different from the generics I know.  
Generics in Go generalizes the type as opposed to specifying a generic collection to use a type.  
