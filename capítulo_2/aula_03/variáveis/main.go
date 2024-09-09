package main

import "fmt"

// Não se pode usar o operador ':=' (marmota) fora de um bloco de código
// Variável global - Abrange todo o pacote
var y = "Batatas, bom dia"

func main() {
	// O operador ':=' (marmota) atrbui um tipo e um valor
	x := 10 
	// y := "Olá, bom dia!"
	z := 10.0

	// [%v] -> Valor | [%T] -> Tipo
	fmt.Printf("x: %v, %T\n", x, x) 
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	// Aqui é atribuídio um novo valor para a vaiável 'x' - Não se pode redeclarar um tipo
	x = 20

	fmt.Printf("x: %v, %T", x, x)
}