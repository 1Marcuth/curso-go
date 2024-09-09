package main

import "fmt"

func main() {
	x := 0

	for {
		if x < 10 {
			fmt.Println("x é menor do que 10!", x)
			x++
		} else {
			fmt.Println("x é maior do que 10!", x)
			break
		}
	}

	fmt.Println("O loop for foi quebrado!")


	// Equivalente ao "while (true)"
	// for {
	// 	fmt.Println(1)
	// }
}