package main

import "fmt"

type produto struct { //estrutura agrupadora de dados ~como uma classe mas não tem herança, etc
	nome     string
	preco    float64
	desconto float64
}

//Método: funçao com receiver (receptor)
func (p produto) precoComDesconto() float64 { //produto é o receiver
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    2.09,
		desconto: 0.02,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Tênis", 500.0, 0.5}
	fmt.Println(produto2, produto2.precoComDesconto())
}
