package matematica

/*
calculo faz qualquer tipo de calculo bastaa vc enviar a funcao
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)

}

// Multiplicador multiplica
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor restorna a Divisao de numeroA com numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//retorna divisao com resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
