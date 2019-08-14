package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Lucas"} //não poderia ser um array nesse contexto nem um [...]
	imprimirAprovados(aprovados...)                  //espalha cada um dos elementos do slice pra passar como função variática
}
