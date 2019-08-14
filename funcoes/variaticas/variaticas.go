package main

import "fmt"

func media(numeros ...float64) float64 { //os ... definem que a função é variática e pode receber quantos parâmetros quiser
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(9.8, 5.8, 7.2))
}
