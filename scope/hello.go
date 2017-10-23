package main

import "fmt"

func foo() string {
	return "different≠another"
}

func main() {
	x := "Hello World!"
	if true {
		x := foo()
		fmt.Println(x)
	}
	fmt.Println(x)
}

// The ":=" operator is useful,
// but if you do not pay attention to the scope,
// the result disappears with the above processing.
//
// The above processing is a common processing that
// replaces default values with some processing results
// depending on conditions.
