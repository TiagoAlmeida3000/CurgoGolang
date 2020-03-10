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
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("O apartamento é: %+v\n", apartamento)

	chacara := Imovel{
		y:     85,
		Nome:  "Chacara",
		x:     22,
		Valor: 55,
	}
	fmt.Printf("A Chacara é: %+v\n", chacara)

	casa.Nome = "Casa"
	casa.Valor = 450000
	casa.x = 18
	casa.y = 31
	fmt.Printf("A Casa é: %+v\n", casa)
}
