package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i > 10 {
			break // Interrupção no momento em que i = 11
		}
		fmt.Println(i) //  1 2 3 4 5 6 7 8 9 10
	}

	fmt.Println(i) // 11

}
