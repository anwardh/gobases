package main

import (
	"fmt"
	"net/http"
	"os"
)

// Definição da estrutura meuErroCustomizado com seus campos e valores
// para personalizar a interface Error( )
type meuErroCustomizado struct {
	status int
	msg    string
}

// Definição do método Error (que retorna uma string)
// da estrutura meuErro quando houver um erro (status maior ou igual a 400)
func (m *meuErroCustomizado) Error() string {
	// Retorno do método formatado como string
	return fmt.Sprintf("%d - %v", m.status, m.msg)
}

// Agora criamos uma função para preencher nosso erro personalizado
// com quaisquer dados necessários para descrevê-lo.
func meuTesteDeErroCustomizado(status int) (int, error) {
	if status >= http.StatusBadRequest /* 400 */ {
		// Retorno do stsatus e da função que customizou o erro
		return http.StatusInternalServerError /* 500 */, &meuErroCustomizado{
			status: status,
			msg:    "Algo deu errado",
		}
	}
	// Senão, irá retornar o status 200 (quando dá certo) e 'nulo' para o erro
	return http.StatusOK /* 200 */, nil
}

func main() {
	//Verificação do tratamento de erro com a declaração das varáveis
	// status e erro, que receberão os dois retornos da função meuTesteDeErroCustomizado
	status, erro := meuTesteDeErroCustomizado(http.StatusInternalServerError) //l

	// Validação do erro
	// O erro só será exibido se for diferente de nil
	if erro != nil { // Se o erro não for nulo
		fmt.Println(erro) // Impressão do erro

		os.Exit(1) // Função que encerra o programa com um código,
		// 1 - quando o programa fecha com algum erro
	}
	fmt.Printf("Status: %d, Funciona!", status) // Status 200, Funciona!
}
