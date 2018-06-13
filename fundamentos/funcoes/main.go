package main

import (
	"cursoDeGo/fundamentos/funcoes/matematica"
	"fmt"
)

func main() {
	resultadoMultiplicacao := multiplicados(1, 2)
	resultadoSoma := matematica.Soma(1, 1)
	resultadoDivisao := matematica.Divisor(10, 1)
	resultadoDivisaoComResto, resto := matematica.DivisorComResto(10, 1)

	fmt.Println("Valor da Multiplicação", resultadoMultiplicacao)
	fmt.Println("Valor da Soma", resultadoSoma)
	fmt.Println("Valor da Divisao", resultadoDivisao)
	fmt.Println("Valor da Divisao Com Resto", resultadoDivisaoComResto)
	fmt.Println("Valor do Resto", resto)
}

func multiplicados(x int, y int) int {
	return x * y
}
