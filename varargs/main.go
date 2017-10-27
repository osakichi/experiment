package main

import "fmt"

func varargs(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	varargs("Are", "you", "receiving", "me", "?")
}
