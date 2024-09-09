package main

import "fmt"

var x = 10
var y = 10.0

func main() {
	// x := 10
	// y := 10.0

	// Não funciona - x = 10.5
	
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)
	fmt.Printf("Valor: %v, Tipo: %T", y, y)
}