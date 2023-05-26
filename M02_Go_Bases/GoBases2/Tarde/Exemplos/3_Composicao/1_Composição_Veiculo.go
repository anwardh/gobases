package main

import "fmt"

// Definição da estrutura base Veiculo
type Veiculo struct {
	hora float64
	km   float64
}

// Declaração do Método Detalhe para impressão dos valores correspondentes a Veiculo
func (v Veiculo) detalhe() {
	fmt.Printf("km:\t%.2f \nhora:\t%.2f\n", v.km, v.hora)
}

// Definição da Estrutura Composta (Embutida) base Automovel, onde será declarado o tipo do veículo
type Automovel struct {
	v Veiculo
}

// Definição do Método que recebe o tempo em minutos (60min) e calcula a distância com base numa velocidade de 100Km/h
func (a *Automovel) Correr(minutos int) {
	a.v.hora = float64(minutos) / 60 // Tempo
	a.v.km = a.v.hora * 100          // Velocidade
}

// Definição do Método Detalhe que 'chama' o Método da Estrutura Pai - detalhe
func (a *Automovel) Detalhe() {
	fmt.Println("\nV:\tAutomovel") // V:      Automovel
	a.v.detalhe()
}

type Moto struct {
	v Veiculo
}

func (m *Moto) Correr(minutos int) {
	m.v.hora = float64(minutos) / 60
	m.v.km = m.v.hora * 80
}

func (m *Moto) Detalhe() {
	fmt.Println("\nV:\tMoto") // V:      Moto
	m.v.detalhe()
}

func main() {

	auto := Automovel{}
	auto.Correr(60)
	auto.Detalhe()

	moto := Moto{}
	moto.Correr(60)
	moto.Detalhe()
}
