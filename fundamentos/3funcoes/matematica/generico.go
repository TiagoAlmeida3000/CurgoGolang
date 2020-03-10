package matematica

//Calculo executa qualquer tipo de calculo basta enviar a funcao desejada
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica o numero x e y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetua a divisao do numeroA e numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto efetua a divisao do numeraA pelo numeroB e retorna o resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
