package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func inList(s string, l []string) bool {

	contains := false

	for _, word := range l {
		if word == s {
			contains = true
		}
	}
	return contains
}

func WordCount(s string) map[string]int {

	slist := strings.Fields(s)

	m := make(map[string]int)

	for _, word := range slist {
		if inList(word, slist) {
			m[word] += 1
		} else {
			m[word] = 1
		}

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
