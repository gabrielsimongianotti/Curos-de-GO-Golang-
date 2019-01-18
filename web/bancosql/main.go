package main

import (
	"fmt"
	"net/http"

	"github.com/Curos-de-GO-Golang/web/bancosql/manipulador"
	"github.com/Curos-de-GO-Golang/web/bancosql/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor")
}
func main() {
	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println(" [main] ao erro abrir banco", err.Error())
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
