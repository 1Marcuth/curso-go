package main

import "fmt"

func main() {
	if x := 10; x > 100 {
		fmt.Printf("%v é maior do que 100!\n", x)
	} else if x < 10 {
		fmt.Printf("%v é menor do que 10!\n", x)
	} else {
		fmt.Printf("%v não é menor do que 10 nem maior do que 100!\n", x)
	}
}