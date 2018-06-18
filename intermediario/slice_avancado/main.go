package main

import (
	"fmt"
)

func main() {

	cidades := make([]string, 5)
	cidades[0] = "São Paulo"
	cidades[1] = "Brasilia"
	cidades[2] = "João Pessoa"
	cidades[3] = "Rio de Janeiro"
	cidades[4] = "Sao Luiz"

	cidadesVizinhas := cidades[1:3]
	fmt.Println(cidadesVizinhas)

	primeiras := cidades[:2]
	fmt.Println(primeiras)

	ultimas := cidades[3:]
	fmt.Println(ultimas)

}
