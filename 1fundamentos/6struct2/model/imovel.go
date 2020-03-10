package model

//Imovel representa informações de um imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `joson:"coordenada_y"`
	Nome  string `joson:"nome"`
	valor int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
