package model

//Imovel - Strucct Imovel
type Imovel struct {
	Nome  string `json:"nome"`
	X     int    `json:"coordeada_x"`
	Y     int    `json:"coordenada_y"`
	valor int
}

//SetValor - Define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor - Retonra o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
