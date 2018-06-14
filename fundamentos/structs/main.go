package main

import (
	"fmt"
)

//Imovel - Tipo Imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Println("A casa", casa)

	apartamento := Imovel{17, 56, "Meu AP", 100}
	fmt.Println("O apartamento", apartamento)

	chacara := Imovel{
		X:     85,
		Y:     10,
		Nome:  "teste",
		valor: 100,
	}
	fmt.Println("chacara", chacara)

	casa.Nome = "Minha casa"
	fmt.Println("Nova casa", casa)
}
