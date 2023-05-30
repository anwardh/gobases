package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Criação de ctx
	ctx := context.Background()
	// Criação de deadline, com um tempo de duração de 5s
	deadline := time.Now().Add(time.Second * 5)
	fmt.Println("Dormiu demais")
	// Instância de ctx com o parâmetro de deadline
	ctx, _ = context.WithDeadline(ctx, deadline)

	// Passados os 5s, o canal ctx.Done receberá uma mensagem
	// indicando que o contexto foi cancelado
	//<-ctx.Done()
	// Impressão do erro - tempo de context excedido
	fmt.Println(ctx.Err().Error())
}
