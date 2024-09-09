package main

import (
	"fmt"
)

func main() {
	tôMortoDeCansaço := true
	tenhoCama := true
	
	if tôMortoDeCansaço {
		fmt.Println("um dia eu gostaria de ir pra cama")
	} else if tôMortoDeCansaço && !(tenhoCama) {
		fmt.Println("To cansado e não tenho cama :(")
	} else {
		fmt.Println("Opa! Tenho cama, poderei ir tirar um cochilo")
	}
}