package main

import "fmt"

func main() {
	/*
		&&: e
		||: ou
	*/

	x := 9
	// y := "Batata"

	if x == 2 || x == 3 {
		fmt.Println("x é '2' ou '3'")
	}

	if x % 2 == 0 && x % 3 == 0 {
		fmt.Println("x é múltiplo de '2' e também de '3'")
	}

	if x % 2 == 0 || x % 3 == 0 {
		fmt.Println("x é múltiplo de '2' ou de '3'")
	}

	if !(x % 2 == 0) && x % 3 == 0 {
		fmt.Println("x é não múltiplo de '2' e é múltiplo de '3'")
	}


	/*
		if x == 2 && y == "Batata" {
			fmt.Println("x é '2' e y é 'Batata'")
		}
	*/ 
}