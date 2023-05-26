package main

import (
	"fmt"
	"math" // Importação do Pacote Math
)

// Definição da Estrutura Círculo
type Circulo struct {
	raio float64
}

// Definição do Método Área, pertencetente a estrutura Círculo
func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// Definição do Método Perímetro, pertencetente a estrutura Círculo
func (c Circulo) perimetro() float64 {
	return 2 * (math.Pi * c.raio)
}

// Uso de Ponteiro sobre a Estrtura Círulo para que seja alterado o valor de sua variável correspontente "c"
func (c *Circulo) setRaio(r float64) {
	c.raio = r
}

func main() {

	// Declaração da variável x recebendo o valor 5 para o campo raio de sua estrutura
	x := Circulo{raio: 5}

	// Execução do Método Área
	fmt.Printf("Área: %.2f\n", x.area()) // 78.54

	// Execuração do Método Perímetro
	fmt.Printf("Perímetro: %.2f\n", x.perimetro()) // 31.42

	// Alteração do valor da variável
	x.setRaio(10)

	fmt.Printf("\nÁrea: %.2f\n", x.area())         // 314.16
	fmt.Printf("Perímetro: %.2f\n", x.perimetro()) // 62.83

}
