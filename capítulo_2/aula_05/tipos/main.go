package main

import "fmt"

// var x int = 19 - Não é preciso informar o tipo, a própria linguagem "adivinha" o tipo 
// var x = 19 
var x int // Declaração de apenas o tipo da variável
// x = 10 - Não funciona (Quando declaramos a variável sem atribuir o valor, a atribuição de valor só poderá ser feita em um bloco de código)

func main() {
	// x = 20 - Pode
	// x = 20.5 - Não pode, os tipos são imutáveis
	x = 10

	fmt.Printf("Valor: %v, Tipo: %T", x, x)
}