package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Curos-de-GO-Golang/web/bancosql/repo"

	"github.com/Curos-de-GO-Golang/web/bancosql/model"
)

//local minipula rota
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		fmt.Println("Não foi enviado um numero valido.")
		return
	}
	sql := "select country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possivel pesquisar esse numero .", http.StatusInternalServerError)
		fmt.Println("não deletou a query: ", sql, " erro ", err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possivel pesquisar esse numero .", http.StatusInternalServerError)
			fmt.Println("não fazer o binding dos dados do banco na struct erro ", err.Error())
			return
		}
	}
	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "houve um erro na rederização da pagina.", http.StatusInternalServerError)
		fmt.Println("erro na execucao do modelo: ", err.Error())
	}
	//inserçao no banco
	sql = "insert into logquery(daterequest) values (?)"
	resultado, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("Erro na inclusao do log: ", sql, " erro:", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("Erro ao pegar o numero de linhas afetadas pelo comando:", sql, " - ", err.Error())
	}
	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
}
