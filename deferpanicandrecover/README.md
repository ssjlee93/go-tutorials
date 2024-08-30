# Defer, Panic, and Recover
[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
`Defer`, `Panic`, and `Recover` runs Go code in a separate goruntime.

> Defer statement pushes a function call onto a list.  
> The list of saved calls is executed after the surrounding function returns.  
> Defer is commonly used to simplify functions that perform various clean-up actions.

provided example is a clean-up.  
without proper clean-up, application may lead to memory leaks.  

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
