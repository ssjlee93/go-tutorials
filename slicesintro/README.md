# Go Slices: usage and internals
TODO : Complete this note. Will return for better understanding.  
[Blog : Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## Introduction
slices = arrays with convenient and efficient means of working with sequence of typed data

## Arrays
* slice = built on top of array type.  
* array type definition specifies length and data type.  
* arrays do not need to be initialized explicitly. all non-specified values will use zero-values.