# concurrency
[concurrency](https://go.dev/tour/concurrency/1)

## goruntimes
[goruntimes](./goruntimes.go)  
goruntime = lightweight runtime for go.  
```go
go f(x, y, z)
```
> starts a new goroutine running

> Goroutines run in the same address space, 
> so access to shared memory must be synchronized. 
> The sync package provides useful primitives, 
> although you won't need them much in Go as there are other primitives.

## Channels
[channels](./channels.go)  

> Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

channels need to be initialized before use
```go
ch := make(chan int)
```

> By default, sends and receives block until the other side is ready.
> This allows goroutines to synchronize without explicit locks or condition variables.

## Buffered Channels
[buffered-channels](./buffered-channels.go)  
```go
// var := make a channel of integer with buffer size 100
ch := make(chan int, 100)
```

> Sends to a buffered channel block only when the buffer is full. 
> Receives block when the buffer is empty.

## Range and Close
[range-and-close](./range-and-close.go)  
> A sender can close a channel to indicate that no more values will be sent.
> Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```go
v, ok := <-ch
```

> `ok` is `false` if there are no more values to receive and the channel is closed.
>
> Note: Only the sender should close a channel, never the receiver. 
> Sending on a closed channel will cause a panic.
> 
> Another note:  Closing is only necessary when the receiver must be told there are no more values coming


## Select
[select](./select.go)  

> The `select` statement lets a goroutine wait on multiple communication operations.
> A `select` blocks until one of its cases can run, then it executes that case. 
> It chooses one at random if multiple are ready.

## Default Selection
[default-selection](./default-selection.go)  
> The `default` case in a `select` is run if no other case is ready.
> Use a `default` case to try a send or receive without blocking:

```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

## Exercise: Equivalent Binary Trees
[exercise-equivalent-binary-trees](./exercise-equivalent-binary-trees.go)  

the solution I found properly uses `defer` to close the channel after tree traversal.  
also, instead of limiting myself to a single channel, solution would concurrently process right and left nodes separaqtely.  

## sync.Mutex
[mutex-counter](./mutex-counter.go)  
mutual exclusion (mutex) = only one goruntime can access a variable to avoid conflicts.  
mutex = conventional data structure for mutual exclusion.  

go standard library has `sync` with `sync.Mutex` obj to achieve mutual exclusion.  
`sync.Mutex` has two methods:  
* `Lock`
* `Unlock`

```go
mu.Lock()
// block of code to be executed in mutual exclusion.
mu.Unlock()
```

can also use `defer` to ensure the mutex will be unlocked later.  
```go
mu.Lock()
// code block
defer mu.Unlock()
```

## Exercise: Web Crawler
[exercise-web-crawler](./exercise-web-crawler.go)  
the most difficult programming problem I've ever faced.  

### Approach 1 
I wanted to use channels and mutex together.  
Using mutex, I thought I had to lock the calls to `.Fetch()`  
then send the resulting body or urls to a channel and receive it outside the method.  

**problem**  
channels were being deadlocked.  
I don't really understand why at this point.  

## Approach 2
Some solutions online uses structure similar to `mutex.counter.go`  
Instead of using mutex on fetch calls, use mutex when getting and setting values to the hashmap of seen urls.  

**problem**
I do not know how to properly receive the results.  
Since multiple instances of goroutines are processing `Crawl()` calls, program terminates before receiving results.  
also tried using channels to receive the results.  

## Approach 3
sync.WaitGroup  
just found a solution online and used it.  
will need to revisit for better understanding.  

moving on so I can code more.  

this is my current solution.  
