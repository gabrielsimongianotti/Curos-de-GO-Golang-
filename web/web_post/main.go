package main

import (
	"Curos-de-GO-Golang/web/web_post/model"
	"bytes"
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
	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Zé"

	conteudoEnviado, err := json.Marshal(usuario)

	if err != nil {
		fmt.Println("[main] Erro ao gerar o json. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/10txg9b1", bytes.NewBuffer(conteudoEnviado))
	if err != nil {
		fmt.Println("[main] Erro ao gerar o request para o requestbin. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main]  Erro ao chamar o servico do requestbin. Erro: ", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo da pagina requestbin. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}
