package main

import "fmt"

type MyInt int

var x MyInt
var y int

func main() {
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)

	x = 42

	fmt.Println(x)

	y = int(x)
	fmt.Printf("Valor: %v, Tipo: %T\n", y, y)
}