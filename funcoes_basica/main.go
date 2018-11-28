package main

import (
	"alura/udemy/funcoes_basica/matematica"
	"fmt"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicacao é de %d\r\n", resultado)
	resultado = matematica.Soma(2, 3)
	fmt.Printf("O resultado da soma é de %d\r\n", resultado)
}
