package main

import "fmt"

func main() {
	var a int //nesses casos preciso tipar porque não estou dando um valor
	var b float64
	var c bool
	var d string
	var e *int //ponteiro é null por padrao quando nao atribui nada

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
	//<nil> = null
}
