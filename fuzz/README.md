# Tutorial: Getting started with fuzzing
[Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)  

Main documentation on fuzzing.  
[Go Fuzzing](https://go.dev/doc/security/fuzz/#glossary)  

> With fuzzing, random data is run against your test in an attempt to find vulnerabilities or crash-causing inputs. 
> Some examples of vulnerabilities that can be found by fuzzing are 
> SQL injection, 
> buffer overflow, 
> denial of service and 
> cross-site scripting attacks.

many security issues can be prevented with fuzzing.  

## Create a folder for your code¶
Done.  

## Add code to test¶
Done.  

## Add a unit test¶
Done.  

## Add a fuzz test¶
fuzzing comes up with inputs for my code.  
pros of using fuzzing : unpredictable output.  

fuzz tests begin with §FuzzXXX§
§t.Fuzz§ instead of §t.Run§

run fuzz tests:  
§§§go
$ go test -fuzz=Fuzz
§§§

## Fix the invalid string error¶

## Fix the double reverse error¶
