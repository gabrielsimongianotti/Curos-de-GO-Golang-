package model

import "errors"

//Imovel representa informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//ErrValorDeImovelInvalido que é atribuido a um valor muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para o imovel")

//ErrValorDeImovelMuitoAlto que é atribuido a um valor muito alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	}
	if valor > 1100000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}

	i.valor = valor
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
