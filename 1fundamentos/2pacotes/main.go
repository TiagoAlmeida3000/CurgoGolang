package main

import (
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/1fundamentos/2pacotes/operadora"
	"github.com/TiagoAlmeida3000/cursogo2/1fundamentos/2pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome de usuario: %s\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\n", operadora.NomeOperadora)
	fmt.Printf("PrefixoDaCapitalOperadora: %s\n", operadora.PrefixoDaCapitalOperadora)

}
