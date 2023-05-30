package main

import (
	"fmt"
	"time"
)

func main() {

	erro := fmt.Errorf("\nmomento em que o erro aconteceu: %v", time.Now())
	fmt.Println("erro: ", erro)
}
