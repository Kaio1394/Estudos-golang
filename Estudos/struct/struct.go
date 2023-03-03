package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
// Método criado a partir parte de uma estrutura
func (p produto) precomComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	// alternativa: produt := produto{"Lapis", 1.79, 0.05}

	fmt.Println(produto1, produto1.precomComDesconto())
}
