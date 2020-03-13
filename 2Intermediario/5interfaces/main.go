package main

import (
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/2Intermediario/5interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "jojo da silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvlirUmpatoNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvlirUmpatoNoLago(p model.Pato) {
	fmt.Println(p.Quack())
}
