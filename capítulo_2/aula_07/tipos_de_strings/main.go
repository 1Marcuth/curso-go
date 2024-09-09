package main

import "fmt"

func main() {
	interpretedStringLiteral := "Oi, bom dia\nComo vai?\tEspero \"que\" esteja tudo bem!"

	fmt.Println(interpretedStringLiteral)

	rawStringLiteral := `Olá\n Você\t gosta de
	pipoca?`

	fmt.Println(rawStringLiteral)
}