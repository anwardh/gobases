package main

import "fmt"

func main() {

	x, y, z := 10, 20, 30

	fmt.Println("O valor de X é: ", x)
	fmt.Println("O valor de Y é: ", y)
	fmt.Println("O valor de Z é: ", z)

	fmt.Println("\nx < y && x > z", x < y && x > z)
	fmt.Println("x < y || x > z", x < y || x > z)
	fmt.Println("!(x < y && x > z)", !(x < y && x > z))
}
