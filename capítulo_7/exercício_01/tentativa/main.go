package main

import "fmt"

func main() {
	limit := 10_000

	for i := 1; i <= limit; i++ {
		fmt.Printf("%v; ", i)
	}
}