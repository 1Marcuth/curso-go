package main

import "fmt"

func main() {
	myYearOfBirth := 2006
	currentYear := 2023

	i := int(myYearOfBirth)

	for {
		if i >= currentYear {
			break
		}

		fmt.Println(i)

		i++
	}
}