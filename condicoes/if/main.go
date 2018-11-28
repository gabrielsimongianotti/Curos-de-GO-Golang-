package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "Teste"
	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("meses")
	}
	if situacao {
		fmt.Println("ele ta devendo")
	}
	if !situacao {
		fmt.Println("Não esta devendo")
	}
	if cidade == "Teste" {
		fmt.Println("cidede Teste")
	}
	if descricao, status := haQuantotempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}
	fmt.Println("Obrigado por o consultar")
}

func haQuantotempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo."
		return
	}
	descricao = "O cliente esta com o carne em dia."
	return
}
