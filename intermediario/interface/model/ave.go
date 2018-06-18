package model

//Galinha interface
type Galinha interface {
	Cacareja() string
}

//Pata interface
type Pata interface {
	Quack() string
}

//Ave tipo
type Ave struct {
	Nome string
}

//Cacareja coco
func (a Ave) Cacareja() string {
	return "Cóco"
}

//Quack quac quack
func (a Ave) Quack() string {
	return "Quack, Quack"
}
