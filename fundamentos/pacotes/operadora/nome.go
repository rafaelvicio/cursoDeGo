package operadora

import (
	"cursoDeGo/fundamentos/pacotes/prefixo"
	"strconv"
)

var nomeOperadora = "Minha Operadora"

//PrefixoOperadora prefixo e nome da operadora
var PrefixoOperadora = strconv.Itoa(prefixo.Capital) + " " + nomeOperadora
