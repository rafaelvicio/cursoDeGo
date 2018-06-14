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

	casa := new(Imovel)
	fmt.Println("Endereço de memoria casa", &casa)

	apartamento := Imovel{17, 28, "AP", 200}
	fmt.Println("Endereço de memoria apartamento", &apartamento)

	mudaImovel(&apartamento)
	fmt.Println("Endereço de memoria apartamento", &apartamento)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y + 10
}
