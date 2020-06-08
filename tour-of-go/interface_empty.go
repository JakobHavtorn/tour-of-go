package main

import "fmt"

// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
