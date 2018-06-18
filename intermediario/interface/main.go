package main

import (
	"cursoDeGo/intermediario/interface/model"
	"fmt"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da silva"

	queroAcordarComoGalinha(jojo)
	queroOuvirUmaPata(jojo)
}

func queroAcordarComoGalinha(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPata(g model.Pata) {
	fmt.Println(g.Quack())
}
