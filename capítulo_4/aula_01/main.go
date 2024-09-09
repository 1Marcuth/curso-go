package main

import "fmt"

var x bool

/*
Operadores relacionais:
	Igualdade: ==
	Igual ou menor que: <=
	Igual ou maior que: >=
	Diferente: !=
	Menor que: <
	Maior que: >
*/

func main() {
	fmt.Println(x) // Valor zero

	x = true

	fmt.Println(x) // Valor atribuído

	x = (10 < 100)
	y := 10 == 100
	z := 10 > 100
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}