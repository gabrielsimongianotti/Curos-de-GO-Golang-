package main

import (
	"fmt"
	"net/http"

	"github.com/Curos-de-GO-Golang/web/servidor_web/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
