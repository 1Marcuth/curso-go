package main

import "fmt"

func main() {
	x := 100

	fmt.Printf("Valor decimal: %d, Valor hexadecimal: %#x, Valor binário: %b\n", x, x, x)
	fmt.Printf("Valor hexadecimal sem hash: %x", x)
}