package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

func main() {
	r := gin.Default()
	r.POST("/produtos", func(ctx *gin.Context) {
		var req request
		erro := ctx.ShouldBindJSON(&req)
		if erro != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": erro.Error(),
			})
			return
		}
		req.ID = 4
		ctx.JSON(http.StatusOK, req)
	})

	r.Run()
}
