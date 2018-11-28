package main

import "fmt"

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	//new pega o endereco de memoria
	casa := new(Imovel)
	//%p mostra o enderecode memoria da casa
	fmt.Printf("Casa é: %p\r\n", &casa)

	chacara := Imovel{17, 28, "blablabla", 280000}
	fmt.Printf("chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	//& passa o endereco de memoria
	fmt.Printf("chacara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
