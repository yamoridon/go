package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
