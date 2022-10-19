package main

import (
	"fmt"
)

type A int

func (a A) Foo() {
	a++
	fmt.Println("foo", a)
}
func main() {
	var a A
	fmt.Println("before", a) // 0
	a.Foo()                  // 1
	fmt.Println("after", a)  // 0
}
