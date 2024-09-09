package main

import "fmt"

func main() {
	PrintAndPrintln()
	SprintAndSprintln()
}

/*
fmt.Printf() -> Formatar string
*/
func PrintAndPrintln() {
	fmt.Println("Println: Adiciona uma nova linha no final")
	fmt.Print("Print: Não adiciona uma nova linha no final\n")
}

/*
Serve mais para algo como formatação ou junção de strings já que o objetivo dele não é imprimir na tela, mas sim retornar o valor

fmt.Sprintf() -> Formatar string
*/
func SprintAndSprintln() {
	w := "Olá,"
	x := "bom dia"

	y := fmt.Sprint(w, " ", x) // Não adicona uma nova linha no final
	z := fmt.Sprintln(w, x) // Adiciona uma nova linha no final

	fmt.Println(y)
	fmt.Print(z)
}

/*
fmt.Fprint() -> File print (Não quer dizer que é necessáriamente um arquivo)
*/
func Fprint() {

}