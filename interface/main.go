package main

import (
	"alura/udemy/interface/model"
	"fmt"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da silva"
	Cacarejo(jojo)
	Pato(jojo)
}

func Cacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func Pato(p model.Pata) {
	fmt.Println(p.Quack())
}
