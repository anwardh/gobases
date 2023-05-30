package main

import (
	"context"
	"fmt"
)

func saudacoesWrapper(ctx context.Context) {
	saudacoes(ctx)
}

func saudacoes(ctx context.Context) {
	fmt.Println(ctx.Value("Saudações"))
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Saudações", "Turma")
	saudacoesWrapper(ctx)
}
