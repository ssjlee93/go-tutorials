# Methods and Interfaces
[Methods and Interfaces](https://go.dev/tour/methods/1)

## Methods
[methods](./methods.go)  
no classes in Go.  
but `type` has `methods`  
method = function with special receiver arg  
receiver args appears between `func` and function name.

## Methods are functions
[methods-funcs](./methods-funcs.go)  
same as function that takes a `type` as an arg.  

## Methods continued
[methods-continued](./methods-continued.go)  
can declare methods with non-struct types.  
method and its receiver type must be in the same package.  
including primitive types.  

## Pointer receivers\
[methods-pointers](./methods-pointers.go)  
can declare methods with pointer receivers.  
```go
func (t *T) Name() {
    // code block
}
```

modifies the value of T through pointer.  
>  Since methods often need to modify their receiver, 
>  pointer receivers are more common than value receivers.

if `*` is removed on the file, 
file only executes `v.Abs()`
and returns `5`.  

With a value receiver, the method operates on a copy of the original value.  

## Pointers and functions
[methods-pointers-explained](./methods-pointers-explained.go)  
rewritten methods as functions.  

if `*` is removed on the file, 
file panics with an error:  
```shell
cannot use &v (value of type *Vertex) as Vertex value in argument to Scale
```
`*T` receiver will only take `&t`.  
pointer receiver will only expect pointer value.  

## Methods and pointer indirection
[indirection](./indirection.go)  
functions with pointer arg must take a pointer.  
but methods with pointer receivers take either a value or pointer as the receiver when called.  

## Methods and pointer indirection (2)
[indirection-values](./indirection-values.go)

## Choosing a value or pointer receiver
[methods-with-pointer-receivers](./methods-with-pointer-receivers.go)  
reasons for using pointer receivers :  
* modify value for the pointer
* avoid copying the value on each method call  
  more efficient if receiver is a large struct

> In general, all methods on a given type should have either value or pointer receivers, 
> but not a mixture of both. 

## Interfaces
[interfaces](./interfaces.go)  
interface = set of method signatures.  

similar to Java interfaces.  

> A value of interface type can hold any value that implements those methods.

In other words, implementation of the interface can hold any value (struct or else) that implements those methods.  
"a value of interface" means a value that implements the interface.  

### note on line 22
implementation of `Abs()` in `Vertex` receives `*Vertex`.  
because the specific method expects pointer receiver, 
`Abs()` only exists in `*Vertex`,
not in `Vertex`.  
so `a=v` panics an error
```shell
cannot use v (variable of type Vertex) as Abser value in assignment: Vertex does not implement Abser (method Abs has pointer receiver)
```
changing to `&v` resolves the error

## Interfaces are implemented implicitly
[interfaces-are-satisfied-implicitly](./interfaces-are-satisfied-implicitly.go)  
in Go, a type implements an interface by implementing its methods.  
no explicit inheritance necessary.  
> Implicit interfaces decouple the definition of an interface from its implementation, 
> which could then appear in any package without prearrangement.

implicit implementation decouples interfaces.  

instead of using keywords like §implements§ in Java, 
Go allows its children to inherit by implementing just the methods.  

## Interface values
[interface-values](./interface-values.go)  
interface values = tuple of a value and a concrete type.  
```go
(value, type)
```

## Interface values with nil underlying values
[interface-values-with-nil](./interface-values-with-nil.go)  
if concrete value for the interface is `nil`, 
then the interface method will be called with nil receiver.  
no `NullPointerException`.  
handle `nil` gracefully.  
interface value that holds a nil concrete value is non-nil.  
in other words, an interface value is considered non-nil even if it holds a nil concrete value

## Nil interface values
[nil-interface-values](./nil-interface-values.go)  
nil interface values = no value, no concrete type.  
calling a method on nil interface is run-time error  
no type in interface to call the concrete method.  
```shell
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1020faa88]
```

## The empty interface
[empty-interface](./empty-interface.go)  
empty interface = zero methods in interface.  
> An empty interface may hold values of any type.  
> Empty interfaces are used by code that handles values of unknown type.


