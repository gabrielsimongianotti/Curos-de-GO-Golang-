package main

import (
	"alura/udemy/mapas/model"
	"fmt"
)

var cache map[string]model.Imovel

func main() {
	//inicializa o map
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi:%+v\r\n", imovel)
	}

	ap := model.Imovel{}
	ap.Nome = "Casa azul"
	ap.X = 19
	ap.Y = 24
	ap.SetValor(70000)

	cache[ap.Nome] = ap

	fmt.Println("Quantos itens ha no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens ha no cache? ", len(cache))
}
