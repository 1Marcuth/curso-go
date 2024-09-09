package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	x := 2

	switch {
		case x == a:
			fmt.Println("x == a")

		case x == b:
			fmt.Println("x == b")

		case x == c:
			fmt.Println("x == c")
	}
}