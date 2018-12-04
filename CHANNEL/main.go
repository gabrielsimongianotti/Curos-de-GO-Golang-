package main

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	for {
		//para onde aponta a seta é onde recebe a informação
		canal <- "ping"
	}
}
func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//os canal podem receber tipos primitivos
	var canal chan string
	canal = make(chan string)
	go pinger(canal)
	go ponger(canal)
	go impressora(canal)

	var entrada string
	fmt.Scanln(&entrada)

}
