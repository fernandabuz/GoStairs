package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Oi"
	fmt.Println(coisa2)

	var coisa3 dinamico = 3.0
	fmt.Println(coisa3)

	var coisa4 dinamico = true
	fmt.Println(coisa4)

	coisa4 = curso{"Opaaaaa"}
	fmt.Println(coisa4)
}
