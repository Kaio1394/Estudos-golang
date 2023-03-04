package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	ferrari1 := ferrari{"F40", false, 0}
	ferrari1.ligarTurbo()
	fmt.Println(ferrari1)

	// Para usar um tipo interface, deixando de forma genérica, é necessário passar a estrutura como referência
	var ferrari2 esportivo = &ferrari{"F40", false, 0}
	ferrari2.ligarTurbo()
	fmt.Println(ferrari1, ferrari2)
}
