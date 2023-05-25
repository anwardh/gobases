package main

import "fmt"

func main() {

	x, y := 15, 25

	fmt.Println("O valor de X é: ", x)
	fmt.Println("O valor de Y é: ", y)

	x = y
	fmt.Println("x = y: ", x)

	x += y
	fmt.Println("x += y: ", x)

	x -= y
	fmt.Println("x -= y: ", x)

	x *= y
	fmt.Println("x *= y: ", x)

	x /= y
	fmt.Println("x /= y: ", x)

	x &= y
	fmt.Println("x %= y: ", x)

}
