package main

import "fmt"

type Hotdog int

var b Hotdog

func main() {
	x := 10

	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)

	// b = x // Error
}