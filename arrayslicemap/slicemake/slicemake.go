package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        //slice com 10 posições e array interno com 20 posições
	fmt.Println(s, len(s), cap(s)) //cap é a capacidade máxima do slice (no caso o tamanho do array interno)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) //append é para anexar no slice além do definido - ter esse novo slice não significa que mudei o arrey interno
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) //como adicionou um slice a mais que o numero do array, o array dobrou a capacidade
}
