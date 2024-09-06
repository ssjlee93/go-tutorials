package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		counts[words[i]] = counts[words[i]] + 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
