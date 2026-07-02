package main

import "fmt"

func main() {

	nomeDoUsuario := "carlos"
	sobreNomeDoUsuario := "soares"
	nomeCompleto := mostraNomeCompleto(nomeDoUsuario, sobreNomeDoUsuario)
	SaldoDoUsuario := 11
	TipoDePessoa := -100

	mensagem := verificaSaldo(nomeCompleto, SaldoDoUsuario, TipoDePessoa)

	fmt.Println(mensagem)
}

func verificaSaldo(nome string, saldo int, tipo int) string {
	// algoritmo
	var tipoPessoa string

	switch tipo {
	case 1, 4, 6:
		tipoPessoa = "juridica"
	case 2:
		tipoPessoa = "fisica"
	case 3:
		tipoPessoa = "microempreendedor"
	default:
		tipoPessoa = "desconecido"
	}

	if saldo > 1000 {
		return fmt.Sprintf("o saldo do cliente %s é alto, %s", nome, tipoPessoa)
	} else if saldo <= 10 {
		return fmt.Sprintf("o saldo do cliente %s é baixo, %s", nome, tipoPessoa)
	}

	return fmt.Sprintf("o saldo do cliente %s é normal, %s", nome, tipoPessoa)
}

func mostraNomeCompleto(nome string, sobrenome string) string {

	mostraNomeCompleto(nome, sobrenome)

	return fmt.Sprintf("%s %s", nome, sobrenome)
}
