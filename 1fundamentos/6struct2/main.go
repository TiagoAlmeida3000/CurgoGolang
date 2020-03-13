package main

import (
	"encoding/json"
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/1fundamentos/6struct2/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	fmt.Printf("Casa é: %+v\n", casa)
	fmt.Printf("O valor da casa: %d\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)

	fmt.Println("A casa em JSON:", string(objJSON))
}
