package main

import (
	"fmt"
	"time"
)

//irc e sms são 2 canais
var irc = make(chan string)
var sms = make(chan string)

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

func eai(canal chan string) {
	for {
		canal <- "e ai"
	}
}

func blz(canal chan string) {
	for {
		canal <- "blz"
	}
}
func impressora(canal chan string) {
	var msg string
	for {
		//seleciona cada canal e toma decisão
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("zzz...ZZZ", msg)
		}

		time.Sleep(time.Second * 1)
	}
}

func main() {
	//os canal podem receber tipos primitivos
	var canal chan string
	canal = make(chan string)
	go pinger(irc)
	go ponger(irc)
	go eai(sms)
	go blz(sms)
	go impressora(canal)

	var entrada string
	fmt.Scanln(&entrada)

}
