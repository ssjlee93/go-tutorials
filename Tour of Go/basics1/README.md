# Tour of Go
notes for [A Tour of Go](https://go.dev/tour/welcome/1)

## basics
packages, variables, and functions

### Packages
every Go program is bundled in packages
programs start running in package `main`

### Imports
factored import statements > repeated imports

### Exported names
`math.pi` is unexported  
`math.Pi` is exported and can be used

exporting is similar to `public` in Java.  
it is more close to exporting in JavaScript.  

### Functions
Go allows multiple args.  

### Functions continued
Go has shorthand for multiple args.

### Multiple results
Go allows multiple results from one method.  
This is different from traditional statically typed languages 
that require you to bundle multiple data into different data structures.  

### Named return values
instead of initializing return values and returning them manually,
Go allows us to name the return values so that we may skip initializing them.  

`return` without args is called "naked return"  
seems like returning all local variables.  
only use for short functions.  
low readability in long functions.  

### Variables
`var` declares variables  
can be used at package level.  

### Variables with initializers
`var` can have 1 initializer per variable.  
type can be omitted with initializer.  

### Short variable declarations
inside a func, variable can initialize and assign with `:=`  
Outside func requires `var`

### Basic types
* bool

* string

* int  int8  int16  int32  int64
* uint uint8 uint16 uint32 uint64 uintptr

* byte // alias for uint8

* rune // alias for int32 
       // represents a Unicode code point

* float32 float64

* complex64 complex128

### Zero values
variables without initial values are initialized with default zero values.  

### Type conversion
`T<v>` converts a value v to type T.  
```go
var x type = 3
var f float64 = float64(x)
var u unit = unit(f)

// or
x := 3
f := float64(x)
u := unit(f)
```
Similar to type casting in Java.  

### Type inference
Unlike C or Java that only support statical types, Go allows type inferences.  
This allows the programmer to treat Go like dynamically typed language.  
Depending on the syntax of right hand, different numeric types are inferred.  
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants
`const` declares constants.  
supports character, string, boolean, numeric  
can be used globally and locally.  

### Numeric Constants
Numeric constants are high-precision values.  
> An untyped constant takes the type needed by its context.  

Same as variable

list of min and max values.  
unit64 = 2^64  
int64 = +-2^64 / 2  
since it is signed, it extends +2^63 and -2^63  
```go
uint8  : 0 to 255 
uint16 : 0 to 65535 
uint32 : 0 to 4294967295 
uint64 : 0 to 18446744073709551615 
int8   : -128 to 127 
int16  : -32768 to 32767 
int32  : -2147483648 to 2147483647 
int64  : -9223372036854775808 to 9223372036854775807
```
