package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	reader := strings.NewReader("Mensagem no Terminal")

	_, erro := io.Copy(os.Stdout, reader)

	if erro != nil {
		fmt.Println(erro)
	}

}
