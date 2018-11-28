package main

import (
	"fmt"
)

func main() {
	var teste [3]int

	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("capacidade do array ", len(teste))

	cantores := [20]string{"ze", "blablablabla"}

	fmt.Printf("O que há nesse  Array? \n\r%v\r\n", cantores)

	capitais := [...]string{"lisboa", "Brasilia", "luanda", "maputo"}
	fmt.Println("Qual a capacidade deste array? ", len(capitais))

	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\n\r", indice, cidade)
	}
}
