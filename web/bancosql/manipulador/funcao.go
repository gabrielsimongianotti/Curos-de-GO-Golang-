package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao manipula a requisição HTTP na rota /funcao
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando função num pacote")
}
