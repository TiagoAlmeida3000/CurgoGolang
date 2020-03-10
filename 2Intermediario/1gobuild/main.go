package main

import (
	"encoding/json"
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/2Intermediario/1gobuild/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	fmt.Printf("Casa é: %+v\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON:", string(objJSON))
}
