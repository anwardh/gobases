package main

import "fmt"

func main() {

	var a uint = 60 // 60 = 0011 1100
	var b uint = 13 // 13 = 0000 1101
	var c uint = 0

	c = a & b
	fmt.Println("Conjunção - O valor de \"c = a & b\" é: ", c) // 12 0000 1100

	c = a | b
	fmt.Println("Disjunção - O valor de \"c = a | b\" é: ", c) // 61 0011 1101

	c = a ^ b
	fmt.Println("Disjunção Exclusiva - O valor de \"c = a ^ b\" é: ", c) // 49 0011 0001

	c = a << 2                                                                             // Multiplica por 4
	fmt.Println("Deslocamento de bits para a esquerda - O valor de \"c = a << b\" é: ", c) // 240 1111 0000

	c = a >> 2                                                                             // Divide por 4
	fmt.Println("Desclocamento de bits para a direita - o valor de \"c = a >> b\" é: ", c) // 15 0000 1111

}
