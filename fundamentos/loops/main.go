package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("Valor", valor)
	}

	for {
		valor--
		fmt.Println("Valor", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Aprendendo go"
	for indice, letra := range texto {
		fmt.Println(indice)
		fmt.Println(letra)
	}

}
