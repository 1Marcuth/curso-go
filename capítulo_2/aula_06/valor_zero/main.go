package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Println("Valores zero")
	
	fmt.Printf("Tipo: '%T', Valor: '%v'\n", a, a)
	fmt.Printf("Tipo: '%T', Valor: '%v'\n", b, b)
	fmt.Printf("Tipo: '%T', Valor: '%v'\n", c, c)
	fmt.Printf("Tipo: '%T', Valor: '%v'\n", d, d)
}