package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pato representa uma ave tipo pato
type Pato interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja retorna um som emitido por uma galinha
func (a Ave) Cacareja() string {
	return "Cocorico"
}

//Quack retorna um som emitido por um pato
func (a Ave) Quack() string {
	return "Quack"
}
