package main

import (
	"fmt"

	"github.com/anwardh/bootcampGo/M02_Go_Bases/GoBases1/Manha/calculadora/calc1"
	"github.com/anwardh/bootcampGo/M02_Go_Bases/GoBases1/Manha/calculadora/calc2"
)

func main() {

	fmt.Println("Soma: ", calc1.Soma(2, 2))
	fmt.Println("Subtração: ", calc1.Subtracao(2, 2))

	fmt.Println("Multiplicação: ", calc2.Multiplicacao(6, 2))
	fmt.Println("Divisão: ", calc2.Divisao(6, 2))

}
