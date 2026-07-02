package main

import "fmt"

func maiordeidade() {
	fmt.Println("sim")
}

func main() {
	var idadeDoclient uint = 20
	// var maiorDeIdade bool = true
	var nomeDoClient string = "carlos"

	fmt.Println(&idadeDoclient)

	nome(nomeDoClient)
	idade(idadeDoclient)
	maiordeidade()
}

func idade(minhaidade uint) {
	fmt.Printf("Minha idade é %d anos ", minhaidade)
}

func nome(nomedoclient string) {
	fmt.Printf("Minha nome é %s", nomedoclient)
}
