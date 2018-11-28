package main

import (
	"fmt"
)

// Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	x     int
	y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	//%+v mostra o nome da variavel e os seus valores
	fmt.Printf("A casa é %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 76000000}
	fmt.Printf("O Apatamento é : %+v\r\n", apartamento)

	chacara := Imovel{
		y:     85,
		Nome:  "Chacara",
		valor: 55,
		x:     22,
	}
	fmt.Printf("A Chacara é : %+v\r\n", chacara)

	casa.Nome = "casa"
	casa.y = 20
	casa.x = 20
	casa.valor = 10
	fmt.Printf("A casa é %+v\r\n", casa)

}
