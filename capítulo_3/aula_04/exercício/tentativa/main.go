package main

import "fmt"

type MyInt int

var x MyInt

func main() {
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)

	x = 42

	fmt.Println(x)
}