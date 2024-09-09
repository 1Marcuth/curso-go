package main

import "fmt"

func main() {
	//x := 10

	// ! operador de negação
	// if !(x < 100) {
	// 	fmt.Printf("x(%v) é maior que 100", x)
	// }

	/* ---------------------------- */ 

	// "if" com inicialização
	if x := 10; (x < 100) {
		fmt.Printf("x(%v) é maior que 100", x)
	}
}