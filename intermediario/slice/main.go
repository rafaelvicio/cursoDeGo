package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "São Paulo"
	cidades[1] = "Brasilia"
	cidades[2] = "João Pessoa"
	cidades[3] = "Rio de Janeiro"
	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {
		fmt.Println("Cidade: ", cidade, indice)
	}
}
