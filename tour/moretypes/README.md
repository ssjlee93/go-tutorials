# More types : structs, slices, and maps
[More Types](https://go.dev/tour/moretypes/1)

## Pointers
> A pointer holds the memory address of a value.
> The type *T is a pointer to a T value. Its zero value is nil.
```go
var p *int
```
> The & operator generates a pointer to its operand.
```go
i := 42
p = &i
```
> The * operator denotes the pointer's underlying value.
```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
> This is known as "dereferencing" or "indirecting".

## Structs
collection of fields

## Struct fields
fields are accessed by a dot  
Similar to Java classes, JS objects, python dictionary

## Pointers to structs
> Struct fields can be accessed through a struct pointer.

pointer requires `(*p)`  
so accessing fields through pointer becomes `(*p).X`  
too cumbersome  
hence, just `p.X`

## Struct Literals
newly allocated struct values  
can list a subset of fields by stating `Name:`  
order does not matter  
pointer to struct with struct literals
```go
&Struct{val1, val2}
```

## Arrays
`[n]T` is `n` values of type `T`
```go
var a [n]int
```
Go arrays have fixed size like Java.  
no need to loop through all elements to see values of arrays

## Slices
Slices = dynamically sized arrays.  
same as Java ArrayList.  
`[]T` is elements of type `T`  

form slices by specifying low and high indices.  
```go
a[low : high]
```
half-open range : includes low but excludes high.

## Slices are like references to arrays
> A slice does not store any data, it just describes a section of an underlying array.  
> Changing the elements of a slice modifies the corresponding elements of its underlying array.

## Slice literals
like array literals without length
```go
// array literal
[3]bool{true, true, false}
// slice literal
// creates array, then builds slice
[]bool{true, true, false}
```

## Slice defaults
if no low or high bound is specified, defaults to 0 and Len(slice)
```go
a[:10]
a[0:]
a[:]
```
this is similar to Python and its subarrays

## Slice length and capacity
>A slice has both a length and a capacity.  
> The length of a slice is the number of elements it contains.  
> The capacity of a slice is the number of elements in the underlying array, 
> counting from the first element in the slice.
> extend a slice's length by re-slicing it, provided it has sufficient capacity.

```go
// length of a collection
len(x)
// capacity of a collection
cap(x)
```
capacity tells us the maximum capacity of the object.  
for non-dynamic data structures, this is useful.  
re-slicing will change the length, but the capacity remains the same  
because the capacity depends on the underlying array

## Nil slices
zero value of slice = `nil`  
`len(nil)` = 0 and `cap(nil)` = 0 and no underlying array

## Creating a slice with make
`make` creates dynamically sized slices.  
> The `make` function allocates a zeroed array
> and returns a slice that refers to that array:
```go
a := make([]int, 5)  // len(a)=5
// specify capacity
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

## Slices of slices
Slices can contain other slices  
2D Array

## Appending to a slice
[Go builtin](https://pkg.go.dev/builtin#pkg-overview)  
Go append
```go
func append(s []T, vs ...T) []T
```
> If the backing array of s is too small to fit all the given values a bigger array will be allocated. 
> The returned slice will point to the newly allocated array.

Dynamically re-sizing

Further reading :  
[Go Slices: usage and internals](../../slicesintro)

## [sidenote] Why use pointers?
Google offers Go style guide to follow the best practices.  
[Go style guide](https://google.github.io/styleguide/go/)

[receiver type](https://google.github.io/styleguide/go/decisions#receiver-type)

key takeaway : 
> correctness wins over speed and simplicity

if we need do not modify the value, use the value.  
if we need to modify or mutate the receiver, use pointer
