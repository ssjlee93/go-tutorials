# Defer, Panic, and Recover
[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
`Defer`, `Panic`, and `Recover` runs Go code in a separate goruntime.

## Defer
> Defer statement pushes a function call onto a list.  
> The list of saved calls is executed after the surrounding function returns.  
> Defer is commonly used to simplify functions that perform various clean-up actions.

provided example is a clean-up.  
without proper clean-up, application may lead to memory leaks.  

3 rules to `defer` statements : 
> 1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
> 2. Deferred function calls are executed in LIFO order after the surrounding function returns.
> 3. Deferred functions may read and assign to the returning function’s named return values.

3 is useful for modifying error return value of a function.  

## Panic
Panic stops ordinary flow and begins panicking.  

Flow of panic : 
func F() runs
-> F() calls `panic`
-> F execution stops
-> any deferred functions execute normally
-> F returns
-> the caller thinks F as a call to `panic`
-> ... repeat until all functions in `goruntime` stack returns
-> `goruntime` crashes

* `panic` can be initiated by invoking direct
* `panic` can be caused by runtime errors

## Recover
> regains control of a panicking `goroutine`
> only useful inside `defer` functions
> During normal execution, a call to recover will return nil and have no other effect. 
> If the current `goroutine` is panicking, a call to recover will capture the value given to panic and resume normal execution.

recovers from `panic` 

## Further reading
[JSON pkg](https://pkg.go.dev/encoding/json)  
epitome of properly using `defer`, `panic`, and `recover`  

> The convention in the Go libraries is that even when a package uses panic internally, 
> its external API still presents explicit error return values.

## Asked Google Gemini
Q: is `defer` in go similar to async-await in javascript? or Future-promise in Java?
A: No. All 3 are async operations, but purposes and mechanisms are different.

### `defer` in Go

**Purpose**: Primarily used to defer the execution of a function until the enclosing function returns.
**Mechanism**: The deferred function is placed on a stack and executed in LIFO (Last-In-First-Out) order when the enclosing function exits, regardless of whether there's an error or not.
**Common use cases**: Resource management (e.g., closing files), cleanup operations, and ensuring functions are called even if there are errors.

### `async-await` in JS and `Future-Promise` in Java

**Purpose**: async operation.
**Mechanism**: Future-Promise. Execute code asynchronously, check completion, retrieve result, cancel operation
**Common use case**: async tasks

### key differences
**Purpose**: `defer` is primarily for cleanup and resource management, while async-await and Future-promise are for asynchronous operations.
**Mechanism**: `defer` uses a stack-based approach, while async-await and Future-promise leverage Promises and Futures for asynchronous handling.
**Execution timing**: `defer` executes after the enclosing function returns, while async-await and Future-promise can execute asynchronously.
