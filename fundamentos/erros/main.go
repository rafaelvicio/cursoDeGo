package main

import (
	"cursoDeGo/fundamentos/structs_avancados/model"
	"encoding/json"
	"fmt"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa"
	casa.X = 10
	casa.Y = 20
	casa.SetValor(60000)

	fmt.Println("Minha casa", casa)
	fmt.Println("Valor da casa", casa.GetValor())

	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("Deu erro!", err.Error())
		return
	}

	fmt.Println("Casa Json", string(objJSON))

}
