# Writing Web Applications with Go
[Writing Web Applications with Go](https://go.dev/doc/articles/wiki/)
The project to walk through is called "wiki".

## Introduction¶
This article walks me through a full stack servlet where I would have server-side-rendering (SSR).  

## Getting Started¶

## Data Structures
probably a good idea to modularize this portion.  

## Introducing the net/http package
def did not need this in §wiki.go§ page.  

## Using net/http to serve wiki pages¶
finished.  

## Using net/http to serve wiki pages¶
finished.  

## Editing Pages¶
finished.  

## The html/template package¶
template engine.  
Skimming.  
Out of my scope.  

## Handling non-existent pages¶
using this feature, I should be able to handle full path match or non-existent pages "**".  
crucial for production-ready routing.  

## Saving Pages¶
finished.  

## Error handling¶
finished.  
always handle errors.  

## Template caching¶
finished.  

## Validation¶
disallow unexpected path variables.  
prevents arbitrary injections.  
should sit in controllers.  

## Introducing Function Literals and Closures¶
function literals := anonymous functions.  
closure := encloses values defined outside of it.  
