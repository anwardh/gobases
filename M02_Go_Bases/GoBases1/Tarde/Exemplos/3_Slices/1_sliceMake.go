package main

import "fmt"

func main() {

	a := make([]int, 5)

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	a[6] = 6

	fmt.Println(a)      // [1 2 3 4 5]
	fmt.Println(len(a)) // Comprimento 5
	fmt.Println(cap(a)) // Capacidade 5

}
