package main

import (
	"fmt"
	"time"
)

func processar(i int) {

	fmt.Println(i, " - Início")
	time.Sleep(time.Second)
	fmt.Println(i, " - Fim")
}

func main() {

	for i := 0; i < 10; i++ {
		go processar(i)
	}

	//Adicionando uma "PAUSA" até o encerramento do programa

	//time.Sleep(5 * time.Second)
	fmt.Println("Fim do Progrma")
}
