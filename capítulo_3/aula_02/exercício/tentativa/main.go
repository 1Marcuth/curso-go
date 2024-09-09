package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("Se chamam valor zero!")

	fmt.Printf("Tipo: %T, Valor: %v", x, x)
	fmt.Printf("Tipo: %T, Valor: %v", y, y)
	fmt.Printf("Tipo: %T, Valor: %v", z, z)
}