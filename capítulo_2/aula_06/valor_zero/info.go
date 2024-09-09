package main

import "fmt"

var x int // Declaração

func main() {
	//x := 10 - Faz tudo ao mesmo tempo
	x = 10 // Inicialização e atribuição
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)

	x = 1 // Atribuição
	fmt.Printf("Valor: %v, Tipo: %T", x, x)
}