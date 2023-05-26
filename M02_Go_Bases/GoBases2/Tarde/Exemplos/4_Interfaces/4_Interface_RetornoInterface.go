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

// Constantes ques definirão o objeto
const (
	tipoRetangulo = "retangulo"
	tipoCirculo   = "circulo"
)

// Definição da função novaFigura que retornará a Interface Geometria, usada para implementar todos os objetos geométricos
func novaFigura(geoTipo string, valor ...float64) geometria {
	switch geoTipo {
	case tipoRetangulo:
		return retangulo{base: valor[0], altura: valor[1]}

	case tipoCirculo:
		return circulo{raio: valor[0]}
	}
	return nil
}

func main() {

	r := novaFigura(tipoRetangulo, 2, 3)
	fmt.Println("Retângulo")
	fmt.Printf("\tÁrea: %.2f\n\t", r.area())
	fmt.Printf("Perímetro: %.2f\n", r.perimetro())

	fmt.Println()

	c := novaFigura(tipoCirculo, 2)
	fmt.Println("Círculo")
	fmt.Printf("\tÁrea: %.2f\n\t", c.area())
	fmt.Printf("Perímetro: %.2f", c.perimetro())
}
