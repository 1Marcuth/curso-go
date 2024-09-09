package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	// Imprimindo todas as variáveis com apenas uma invocação do Print
	fmt.Println(x, y, z)

	// Imprimindo todas as variáveis com várias invocações do Print
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}