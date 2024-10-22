# Tutorial: Getting started with generics
[Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)

## Create a folder for your code¶
Done.  

## Add non-generic functions¶
Done.  

## Add a generic function to handle multiple types¶
generic function = accepts more than one type for its parameters(s).  
this is Go's way of resolving method overloading.  
Perhaps not the concept of inheritance, but it behaves like method overloading in Java.  

function will declare supported types.  
function call will specify one of the data types.  

generic functions = functions + type parameters + parameters + return types.  
type params = params for supported types.  
calling generic functions = call + type args + args.  

> Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter. 
> Each type constraint specifies the permissible type arguments that calling code can use for the respective type parameter.

type param and type constraint work together.  

> While a type parameter’s constraint typically represents a set of types, 
> at compile time the type parameter stands for a single type – 
> the type provided as a type argument by the calling code. 
> If the type argument’s type isn’t allowed by the type parameter’s constraint, 
> the code won’t compile.

type arg being called work together with type param + type constraint.  

> Keep in mind that a type parameter must support all the operations the generic code is performing on it. 
> For example, if your function’s code were to try to perform string operations (such as indexing) on a type parameter whose constraint included numeric types, 
> the code wouldn’t compile.

my understanding  
```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
```

`K` = type param.  
`comparable` = type constraint.

`V` = type param.  
`int64 | float64`  = a set of type constraints.  

`comparable` constraint 
> allows any type whose values may be used as an operand of the comparison operators == and !=.  

## Remove type arguments when calling the generic function¶
types in type args can be referred.  

## Declare a type constraint¶
