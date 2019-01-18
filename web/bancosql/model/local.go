package model

//Local armazena os dados da localidade pelo seu codigotelefonico
type Local struct {
	Pais             string `json:"pais" db:"country"`
	Cidade           string `json:"cidade" db:"city"`
	Codigotelefonico int    `json:"cod_telefone" db:"telcode"`
}
