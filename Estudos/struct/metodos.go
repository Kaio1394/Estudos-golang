package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// recieve como ponteiro
// necessário para alterar o nome passado para a estrutura
func (p *pessoa) setNomeCOmpleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")

	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Kaio", "Santiago"}

	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCOmpleto("Yuri Santiago")

	fmt.Println(p1.getNomeCompleto())

}
