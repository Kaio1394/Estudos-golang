package main

import "fmt"

// interface 1
type esportivo interface {
	ligarTurbo()
}

// interface 2
type luxuoso interface {
	fazerBaliza()
}

// fazendo a composição da interface 1 e 2 com a interface 3
type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw) fazerBaliza() {
	fmt.Println("Baliza...")
}
func main() {
	var b esportivoLuxuoso = bmw{}
	b.ligarTurbo()
	b.fazerBaliza()
}
