package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro // Campo anônimo. Dessa forma será possível acessar os atributor
	// da estrutura 'carro'. Pseudo herança
	turboLigado bool
}

func (c *carro) setNomeCarro(nome string) {
	c.nome = nome
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidade = 0
	f.turboLigado = true
	fmt.Println(f.nome)
	f.setNomeCarro("f80")
	fmt.Println(f.nome)
}
