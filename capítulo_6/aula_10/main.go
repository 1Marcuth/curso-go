package main

import "fmt"

var x interface{}

func main() {
	x = -176
	/*
		--------- Loop com nome -----------------
		Batata: for i := 0; i <= 10; i++ {
			if i == 6 {
				fmt.Println("Batata quebrada!")
				break Batata
			}
		}
	*/

	switch x.(type) {
		case int:
			fmt.Println("É um número inteiro")

		case float64:
			fmt.Println("É um número flutuante")

		case bool:
			fmt.Println("É um booleano")

		case string:
			fmt.Println("É uma string")
	}

	// Com inicialização
	switch y := 2; y {
		case 1:
			fmt.Println("É um 1!")

		case 2:
			fmt.Println("É um 2!")

		case 3:
			fmt.Println("É um 3!")

		case 4:
			fmt.Println("É um 4!")
	}
}