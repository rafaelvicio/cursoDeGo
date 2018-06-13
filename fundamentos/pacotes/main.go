package main

import (
	"cursoDeGo/fundamentos/pacotes/operadora"
	"cursoDeGo/fundamentos/pacotes/prefixo"
	"fmt"
)

//NomeDoUsuario é o nome od usuário
var NomeDoUsuario = "Rafael Augusto"

func main() {
	fmt.Printf("Nome do Usuário %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital %d\r\n", prefixo.Capital)
	fmt.Printf("Nome e Prefixo %s\r\n", operadora.PrefixoOperadora)
}
