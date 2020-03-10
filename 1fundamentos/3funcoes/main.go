package main

import (
	"fmt"

	"github.com/TiagoAlmeida3000/cursogo2/fundamentos/3funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Println(resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Println(resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Println(resultado)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Println(resultado, resto)
}
