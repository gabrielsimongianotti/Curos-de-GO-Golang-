package repo

import (
	"fmt"
	/*
		_ "github.com/go-sql-driver/mysql" é usado pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

//armaxena conexão com banco
var Db *sqlx.DB

func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:123g@tcp(127.0.0.1:3000)/cursodego")
	if err != nil {
		fmt.Println(" sqlx erro ao conectar ao banco")
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
