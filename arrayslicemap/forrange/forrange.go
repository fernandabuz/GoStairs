package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //é um array e o compilador que conta o numero de elementos de acordo com o que inseri

	for i, numero := range numeros { //equivalente ao foreach
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { //tiro o índice
		fmt.Println(num)
	}
}
