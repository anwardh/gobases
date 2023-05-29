package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	bytesString, erro := io.WriteString(os.Stdout, "Anwar:\t")

	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(bytesString)
}
