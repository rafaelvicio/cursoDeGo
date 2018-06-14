package main

import (
	"fmt"
)

func main() {

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println(len(teste))

	contadores := [2]string{"Rafael", "Augusto"}
	fmt.Println(contadores)

	cidades := [...]string{"Brasilia", "São Paulo"}
	fmt.Println(len(cidades))

	for _, cidade := range cidades {
		fmt.Println("Cidade: ", cidade)
	}

}
