package main

import "fmt"

func Aumentar(v *int) {

	*v++

}

func main() {

	v := 19

	fmt.Println("O endereço de memória de V é:", &v)
	fmt.Println("O valor atual de V, ANTES da função Aumentar é:", v)

	Aumentar(&v)

	fmt.Println("O endereço de memória de V é:", &v)
	fmt.Println("O valor atual de V, DEPOIS da função Aumentar é:", v)

}
