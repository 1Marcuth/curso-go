package main

import (
	"fmt"
)

func main() {
	s := "Hello, world! Olá, mundo!"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
	
	// fmt.Printf("%v\n%T", sb, sb)
}