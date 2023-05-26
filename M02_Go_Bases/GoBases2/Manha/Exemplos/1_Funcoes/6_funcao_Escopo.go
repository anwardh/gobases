package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 0; j < 3; j++ {
			//var num = 3 //-: Escopo Amarelo
			//fmt.Printf("num: %d\n", num)
			fmt.Printf("\tnum: %d\n", j)
		}

		// Erro 1 - num fora de escopo -> Escopo Vermelho
		//fmt.Println(num) // undefined: num

	}
	// Erro 2 - i fora de escopo -> Escopo Verde
	//fmt.Println(i) // undefined: i

}
