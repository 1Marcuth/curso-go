package main

import "fmt"

var y = 10

func main() {
	z := 20
	
	anything(z)
}

func anything(x int) {
	// fmt.Println(y) | A variável 'y' não é assessível nesse contexto, é apenas assessível no bloco onde ela foi declarada

	fmt.Println(y)
	fmt.Println(x)
}