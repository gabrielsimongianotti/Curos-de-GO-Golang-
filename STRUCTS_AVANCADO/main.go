package main

import (
	"alura/udemy/STRUCTS_AVANCADO/model"
	"encoding/json"
	"fmt"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 19
	casa.Y = 27
	casa.SetValor(60000)

	fmt.Printf("casa é: %+v\r\n", casa)

	// _ isso não é um boa pratica
	fmt.Printf("O valor da casa é: %+v\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println(string(objJSON))
}
