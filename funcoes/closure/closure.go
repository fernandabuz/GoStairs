package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() //vem o valor do x la de closure, porque ele respeita o escopo da função - sabe que foi definida lá então mantém na memória
}
