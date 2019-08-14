//não tem operador ternário em go

package main

import "fmt"

func obterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovado" : "Reprovado" ===> mais ou menos assim seria ternário
	if nota >= 6 {
		return "Aprovado :)"
	}
	return "Reprovado :("
}

func main() {
	fmt.Println(obterResultado(5.2))
}
