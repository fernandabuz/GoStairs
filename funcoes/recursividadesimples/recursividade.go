package main

import "fmt"

func fatorial(n uint) uint { //uint não permite números negativos
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
