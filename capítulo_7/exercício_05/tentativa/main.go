package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		res := i % 4

		fmt.Printf("%v %% 4 = %v \n", i, res)
	}
}