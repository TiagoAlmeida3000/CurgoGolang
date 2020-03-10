package main

import "fmt"

var num int
var a string
var float float64
var boolean bool
var palavras rune
var num1, num2 int

var (
	//Valor1 è um valor importante
	Valor1  int
	valor2  int //Valor1 publica, valor2 é privada
	nome    string
	comprou bool
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("numero %d\r\n", num)
	fmt.Printf("comprou %v\r\n", comprou)
	fmt.Println(teste)
}
