package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("cidade.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo. Erro", err.Error())
		return
	}
	/*
		//pega o conteudo
		scanner := bufio.NewScanner(arquivo)
		//le linha por linha
		for scanner.Scan() {
			linha := scanner.Text()
			fmt.Println("O conteúdo da linha é: ", linha)
		}
	*/
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("Erro ao abrir arquivo. Erro", err.Error())
		return
	}

	//le linha e imprime linha por linha
	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indiceItem, item)
		}
	}
	arquivo.Close()
}
