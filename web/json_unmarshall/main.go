package main

import (
	"Curos-de-GO-Golang/web/json_unmarshall/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o jsonplaceholder. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do jsonplaceholder. Erro: ", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo da pagina do jsonplaceholder. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno json do jsonplaceholder. Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é : %+v/r/n", post)
	}

}
