package main

import "fmt"

func somar(a int, b int) int { //o int de fora dos () é o tipo de valor que retorna e sou obrigada a ter um return
	return a + b
}

func imprimir(valor int) { //como nao tem tipo de retorno nao sou obrigada a ter um return
	fmt.Println(valor)
}
