package model

import "errors"

//Imovel representa informações de um imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `joson:"coordenada_y"`
	Nome  string `joson:"nome"`
	valor int
}

//ErrValorDeImovelInvalido é um erro que e apresentado quando é atribuido a imovel um valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado nao é valido para um imovel")

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
