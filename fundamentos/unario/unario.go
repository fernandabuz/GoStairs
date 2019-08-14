//operadores unarios agem em um unico operando
package main

import "fmt"

func main() {
	x := 1
	y := 2

	//apenas postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

	//fmt.Println(x == y--) não pode misturar operações dentro de uma expressão
}
