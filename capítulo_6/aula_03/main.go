package main

import "fmt"

func main() {
	// for hours := 0; hours < 24; hours++ {
	// 	for minutes := 0; minutes < 60; minutes++ {
	// 		for seconds := 0; seconds < 60; seconds++ {
	// 			fmt.Printf("%v:%v:%v\n", hours, minutes, seconds)
	// 		}
	// 	}
	// }

	// for hours := 0; hours < 24; hours++ {
	// 	fmt.Println("- Hora:", hours)

	// 	for minutes := 0; minutes < 60; minutes++ {
	// 		fmt.Print(" ", minutes)
	// 	}

	// 	fmt.Println("\n")
	// }
	
	for month := 1; month <= 12; month++ {
		fmt.Println("- Mês:", month)

		for day := 1; day <= 30; day++ {
			fmt.Printf("Dia: %v, ", day)
		}

		fmt.Println("\n")
	}
}