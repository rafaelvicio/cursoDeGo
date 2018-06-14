package main

import (
	"fmt"
)

func main() {
	meses := 6
	situacao := true
	cidade := "Brasilia"

	// < > != == >= <= && ||
	if meses > 5 {
		fmt.Println("Mes maior que 5")
	}

	if situacao {
		fmt.Println("Tudo certo!")
	}

	if cidade == "Brasilia" {
		fmt.Println("Tem cidade")
	}

	if descricao, status := verificaInformacoes(meses); status {
		fmt.Println("Qual o status", descricao)
		return
	}

	//Variavel descricao não existe mais
	//fmt.Println("Descricao", descricao)
	fmt.Println("Acacbou!")
}

func verificaInformacoes(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "esta devendo"
		return
	}
	descricao = "esta tudo certo"
	return
}
