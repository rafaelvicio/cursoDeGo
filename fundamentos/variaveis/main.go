package main

import (
	"fmt"
)

var endereco string //endereco = ""
var telefone = "999-999"
var quantidade int //quantidade = 0
var comprou bool   //comprou = false
var palavras rune

//VariavelPublica e publica
var VariavelPublica string
var variavellPrivada string // privada

//atalhos

var nome, sobreNome string
var (
	idade      int
	maiorIdade bool
)

func main() {
	teste := "Valor de teste"
	fmt.Println(teste)
	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quanidade: %d\r\n", quantidade)

}
