# Strings, bytes, runes and characters in Go

## Introduction
prerequisite : [slices-intro](../slices-intro)  

[link to another problem](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

the link above leads to some common problems about characters and programming.  

## What is a string?
prerequisite : [Blog : slices]()  
TODO : read slices blog  

strings hold arbitrary bytes.  

## Printing strings
raw print produces some ugly output.  
some chars are neither ASCII nor UTF-8  

we can loop and print each byte  

`%x` gives us hexadecimal digits  
2 per byte  

`%q` gives us quoted output  
escapes non-printable bytes

`+%q` gives us quoted w plus flag output  
escapes non-printable (%q) and non-ASCII (%+)  

### Exercise 1
Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice.

### Exercise 2
Exercise: Loop over the string using the %q format on each byte. What does the output tell you?

## UTF-8 and string literals
> when we store a character value in a string, we store its byte-at-a-time representation  

> “raw string”, enclosed by back quotes, ... contain only literal text.  
> (Regular strings, enclosed by double quotes, can contain escape sequences as we showed above.)

```go
const placeOfInterest = `⌘`
```

> the UTF-8 representation of the string was created ... when the source code was written.  

> only string literals are UTF-8  
> string values can contain arbitrary bytes  
> string literals always contain UTF-8 text as long as they have no byte-level escapes.  

## Code points, characters, and runes

referring to bytes and strings, not characters.  

> The Unicode standard uses the term “code point” to refer to the item represented by a single value.  

> in general, a character may be represented by a number of different sequences of code points, 
> and therefore different sequences of UTF-8 bytes  

code point = rune = int32  

rune literal is read as int32  
so rune needs to be converted to string to be read properly  

Summary
> Go source code is always UTF-8.  
> A string holds arbitrary bytes.  
> A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.  
> Those sequences represent Unicode code points, called runes.  
> No guarantee is made in Go that characters in strings are normalized.  

## Range loop
Go treats UTF-8 specially only through `for-range` loop  

> A for range loop decodes one UTF-8-encoded rune on each iteration. 
> Each time around the loop, 
> the index of the loop is the starting position of the current rune, measured in bytes, 
> and the code point is its value.

the index within for-rage loop designates byte position.  
this is misleading for reading the position of the string.  
running regular for loop causes the rune slice to iterate through byte position.  
so for special chars in Korea, Japanese, and Chinese, 
the position in the string will not be the position in the index since each character contains multiple unicode code points.

## Libraries
`unicode/utf8` pkg helps w unicode  
