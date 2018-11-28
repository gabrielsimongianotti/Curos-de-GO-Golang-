package main

import (
	"alura/udemy/funcoes_avancadas/matematica"
	"fmt"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicacao é de %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 9, 3)
	fmt.Printf("O resultado da divisão é de %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 9)
	fmt.Printf("O resultado da divisão com resultado é de %d e o resto é de %d\r\n", resultado, resto)
	resultado = matematica.Soma(2, 3)
	fmt.Printf("O resultado da soma é de %d\r\n", resultado)
}
