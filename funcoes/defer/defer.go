package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") //atrasa execução do código para até antes de sair do método
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
