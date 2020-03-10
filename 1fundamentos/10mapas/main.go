package main

import (
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/fundamentos/6struct2/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	cache["Casa Amarela"] = casa

	fmt.Println("Ha uma casa amerela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, o que achei foi : %v\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(700000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens ha no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens ha no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n", chave, imovel)
	}

}
