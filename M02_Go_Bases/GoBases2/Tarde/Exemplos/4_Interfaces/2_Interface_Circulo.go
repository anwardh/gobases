package main

import (
	"fmt"
	"math"
)

// Defiinição da Estrutura Circulo
type circulo struct {
	raio float64
}

// Definição da Função e método para o cálculo da área
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// Definição da Função e método para o cálculo do perímetro
func (c circulo) perimetro() float64 {
	return 2 * (math.Pi * c.raio)
}

// Definição da Função que executará a impressão do valor do raio, da área e do perímetro
func detalhe(c circulo) {
	fmt.Println(c)
	fmt.Printf("Área:%.2f\n", c.area())
	fmt.Printf("Perímetro:%.2f", c.perimetro())
}

func main() {

	// Execução da função - detalhe (que carrega todas as demais - área e perímetro)
	c := circulo{raio: 5}
	detalhe(c)
}
