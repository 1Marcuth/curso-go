package main

import "fmt"

func main() {
	p := "joana"

	switch p {
		case "zezinho", "maria":
			fmt.Println("Hoje quem está na firma é o time 1")

		case "joana", "pedrinho":
			fmt.Println("Hoje quem está na firma é o time 2")

		default:
			fmt.Println("tá vazio, ou a balada tava ótima ou é feriado")
	}
}

// func main() {
// 	x := 6

// 	// swicth true - Default
// 	switch {
// 		case x < 5:
// 			fmt.Printf("%v é menor do que 5\n", x)

// 		case x == 5:
// 			fmt.Printf("%v é igual a 5\n", x)

// 		case x > 5:
// 			fmt.Printf("%v é maior do que 5\n", x)
// 	}

// 	switch x {
// 		case 10:
// 			fmt.Printf("x é igual a 10")
// 	}

// 	quemtanoescritoriohoje := "ninguém" // "zezinho" // "joana"

// 	// switch quemtanoescritoriohoje {
// 	// 	case "zezinho":
// 	// 		fmt.Println("hoje quem tá no escritório é o zezinho")

// 	// 	case "marquinhos":
// 	// 		fmt.Println("hoje quem tá no escritério é o marquinhos")

// 	// 	case "joana":
// 	// 		fmt.Println("hoje quem tá no escritório é a joana")

// 	// 	case "maria":
// 	// 		fmt.Println("hoje quem tá no eescritório é a maria")

// 	// }

// 	switch quemtanoescritoriohoje {
// 		case "zezinho":
// 			fmt.Println("hoje quem tá no escritório é o zezinho")
// 			fallthrough

// 		case "marquinhos":
// 			fmt.Println("e sempre que o zezinho vem o marquinhos vem também")

// 		case "joana":
// 			fmt.Println("hoje quem tá no escritório é a joana")

// 		case "maria":
// 			fmt.Println("hoje quem tá no eescritório é a maria")
		
// 		default:
// 			fmt.Println("tá vazio, ou a balada tava ótima ou é feriado")
// 	}

// 	{
// 		fmt.Println("xd")
// 	}
// }