package main

import (
	"cursoDeGo/fundamentos/funcoes/matematica"
	"fmt"
)

func main() {
	resultadoMultiplicacao := multiplicados(1, 2)
	resultadoSoma := matematica.Soma(1, 1)

	fmt.Println("Valor da Multiplicação", resultadoMultiplicacao)
	fmt.Println("Valor da Soma", resultadoSoma)
}

func multiplicados(x int, y int) int {
	return x * y
}
