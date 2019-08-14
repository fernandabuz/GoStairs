package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 10.0},
			item{2, 1, 5.0},
			item{3, 6, 98.0},
		},
	}
	fmt.Printf("O valor total do pedido é R$%.2f", pedido.valorTotal())
}
