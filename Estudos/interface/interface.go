package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}
func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func main() {
	var coisa imprimivel = produto{"TV32", 1563.00}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = pessoa{"Kaio", "Santiago"}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
