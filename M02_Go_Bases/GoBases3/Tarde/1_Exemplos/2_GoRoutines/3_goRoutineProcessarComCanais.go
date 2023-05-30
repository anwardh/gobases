package main

import (
	"fmt"
	"time"
)

func processar(i int, fila chan int) {
	fmt.Println(i, " - Início")
	/*
		*OBS:*
		A partir do momento que a go routine tenha alguma tarefa que trave a execução da função
		(leitura/escrita ao SO ou chamada ao banco de dados por exemplo), nesse caso
		simulamos usando a função sleep para se criar esse "delay", a próxima go routine é executada
	*/
	time.Sleep(time.Second)
	fmt.Println(i, " - Fim")

	// Enviando mensagem ao canal "fila"
	fila <- i
}

func main() {
	/*
		Um canal é uma estrutura que segue o conceito de fila ou FIFO(First in, First Out),
		primeiro a entrar é o primeiro a sair, como se fosse uma fila para comprar pão
	*/
	// Criando o canal chamado "fila" do tipo int
	canal := make(chan int)

	// Criando 10 go routines e passando o canal em comum para a função a ser executada
	for i := 0; i < 10; i++ {
		go processar(i, canal)
	}

	// guardando o valor da mensagem passada no canal numa variavel
	// variavel := <-canal
	// fmt.Println("Valor recebido do canal: ", variavel)

	// esperando 10 mensagens das 10 mensagens que serão passadas no canal
	for i := 0; i < 10; i++ {
		fmt.Println("o valor do canal passado é:", <-canal)
	}

	fmt.Println("Fim do programa")
}
