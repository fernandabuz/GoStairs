package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) //preciso tipar como float64 o y senao nao executa porque ele é inteiro

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(97)) //vai ser a letra referente na tabela a esse número

	//int para string
	fmt.Println("Teste" + strconv.Itoa(123)) //conversao de string ~itoa = int to string

	//string para int
	num, _ := strconv.Atoi("123") //retorna dois valores - primeiro um numero, depois um erro(_)
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
