package main

import (
	"fmt"
)

func redis() {
	fmt.Println("Fechando a conexão do redis")
}

func mysql() {
	fmt.Println("Fechando a conexão do mysql")
}

func mongo() {
	fmt.Println("Fechando a conexão do mongo")
}

func main() {
	defer redis()
	defer mysql()
	defer mongo()
}
