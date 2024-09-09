package main

import "fmt"

func main() {
	myYearOfBirth := 2006
	currentYear := 2023

	i := int(myYearOfBirth)

	for i <= currentYear {
		fmt.Println(i)

		i++
	}
}