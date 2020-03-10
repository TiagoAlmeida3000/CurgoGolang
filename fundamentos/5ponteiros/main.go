package main

import "fmt"

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	x     int
	y     int
	Nome  string
	Valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara", 280000}
	fmt.Printf("Casa é: %p - %+v\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Casa é: %p - %+v\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.x = imovel.x + 10
	imovel.y = imovel.y - 5
}
