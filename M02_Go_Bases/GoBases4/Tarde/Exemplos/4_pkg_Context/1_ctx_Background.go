package main

import (
	"context"
	"fmt"
)

func main() {

	// Dessa forma a gente cria a nossa variável do tipo Context
	ctx := context.Background()
	fmt.Println(ctx) // pcontext.Background
	//fmt.Printf("%T", ctx)

}
