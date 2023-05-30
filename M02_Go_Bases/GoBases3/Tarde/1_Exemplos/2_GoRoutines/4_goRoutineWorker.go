package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	// range esperando chegar alguma mensagem no canal que foi passado por parâmetro
	for res := range msg {
		fmt.Println("Worker: ", workerId, "msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {

	// declarando o canal do tipo inteiro
	msg := make(chan int)
	// outra forma de se declarar um canal do tipo inteiro
	// var msg chan int

	/*
	 Crianção das go routines passando o canal "msg" em comum
	 que haverá a troca de informações entre a função principal do programa
	 e as go routines
	*/
	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	go worker(6, msg)
	go worker(7, msg)
	go worker(8, msg)
	go worker(9, msg)
	go worker(10, msg)
	go worker(11, msg)

	// For que envia 100 mil mensagens ao canal "msg"
	for i := 0; i < 100_000; i++ {
		msg <- i
	}
}
