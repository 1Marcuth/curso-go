package main

import "fmt"

func main() {
	x := 10
	y := 5

	a := x == y
	b := x != y
	c := x <= y
	d := x < y
	e := x >= y
	f := x > y

	fmt.Printf("É igual? %v\nÉ diferente? %v\nÉ menor ou igual? %v\nÉ menor? %v\nÉ maior ou igual? %v\nÉ maior? %v", a, b, c, d, e, f)
}