package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x\n", x, x, x)

	// Deslocando um bit para a esquerda
	y := x << 1

	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x", y, y, y)
}