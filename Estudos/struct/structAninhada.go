package main

import "fmt"

type item struct {
	produtoID int
	qtdt      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco
	}
	return total
}

func main() {
	item1 := item{1, 10, 50.22}
	item2 := item{1, 2, 150.00}
	item3 := item{1, 20, 356.88}
	itens := pedido{123, []item{item1, item2, item3}}
	fmt.Println(itens.valorTotal())
}
