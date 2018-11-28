package main

import (
	"alura/udemy/erro/model"
	"encoding/json"
	"fmt"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "casa da Lucy"
	casa.X = 17
	casa.Y = 20
	if err := casa.SetValor(10000000); err != nil {
		fmt.Println("ERRO no valor da casa: ", err.Error())
		return
	}

	fmt.Printf("casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("erro ", err)
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
