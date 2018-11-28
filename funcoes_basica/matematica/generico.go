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
