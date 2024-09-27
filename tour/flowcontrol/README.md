# Flow control statements : for, if, else, switch, and defer
notes [Flow Control](https://go.dev/tour/flowcontrol/1)

## for
the only loop for Go  
no () but requires {}

## for continued
init and post statements are optional  
no int for loop and no increment marker required

## For is Go's "while"

## Forever
no condition in for-loop = loops forever

## if
no () but requires {]

## If with a short statement
can have short init statement  
scope blocked for that specific if block

## If and else
short init statements are also valid in else block

## Exercise: Loops and Functions
How computer computes square root
> Computers typically compute the square root of x using a loop. 
> Starting with some guess z, 
> we can adjust z based on how close z² is to x, 
> producing a better guess:
```go
z -= (z*z - x) / (2*z)
```
> Repeating this adjustment makes the guess better and better 
> until we reach an answer that is as close to the actual square root as can be.

added convenience print statements to show number of tries and difference with math.Sqrt method.

## Switch
shorter way to write multiple if-else statements.  
runs first case that matches the condition.  
Unlike other languages, no need for `break`

for switch cases with multiple values
```go
switch var {
case v1, v2, v3 ... :
	// code block
}
```
since Go does not use `break` to control switch case flow, we use commas to list values

## Switch evaluation order
top to bottom  
stops when a condition is met

### time
`time` package offers APIs for time

## Switch with no condition
considered true  
allows `switch` case to be used as `if-else` alternatives

## Defer
> A defer statement defers the execution of a function until the surrounding function returns.

`defer` resolves all lines of code within its function, then executes.

> The deferred call's arguments are evaluated immediately, 
> but the function call is not executed until the surrounding function returns.

args are evaluated immediately,  
but execution only after all surrounding function returns

## Stacking defers
Deferred calls are in stack(LIFO)

Further reading:  
[Defer, Panic, and Recover](../../defer-panic-and-recover)



