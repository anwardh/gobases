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

// Definição da Intetrface Geometria com dois métodos qu serão utilizados pelos objetos (círculo e retângulo)
type geometria interface {
	area() float64
	perimetro() float64
}

// Criação da Estruta Reângulo
type retangulo struct {
	base, altura float64
}

// Método Área do Retângulo
func (r retangulo) area() float64 {
	return r.base * r.altura
}

// Método Perímetro do Retângulo
func (r retangulo) perimetro() float64 {
	return 2 * (r.base + r.altura)
}

// Modificação da Função Detalhes para, não mais, receber os parâmetros de um círculo, mas sim, de uma figura geométrica

func detalhes(g geometria) {
	fmt.Println(g)
	fmt.Printf("Área:%.2f\n", g.area())
	fmt.Printf("Perímetro:%.2f\n", g.perimetro())
}

// Criação de uma nova função para gerar um novo objeto novoCirculo com todas as propriedades da Estrutura Círculo
func novoCirculo(v float64) circulo {
	return circulo{raio: v}
}

func main() {

	r := retangulo{base: 3, altura: 4}
	detalhes(r) // Detalhes do Retângulo

	fmt.Println()

	c := circulo{raio: 5}
	detalhes(c) // Detalhes do Círculo

	v := novoCirculo(2)
	detalhes(v)
}
