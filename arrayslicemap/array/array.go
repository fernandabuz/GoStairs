package main

import "fmt"

func main() {
	//homogênea (mesmo tipo) e estática (fixa)
	var nota [3]float64
	fmt.Println(nota)

	nota[0], nota[1], nota[2] = 7.8, 5.8, 9.9
	fmt.Println(nota)

	total := 0.0
	for i := 0; i < len(nota); i++ {
		total += nota[i]
	}

	media := total / float64(len(nota))
	fmt.Printf("Média %.2f\n", media)
}
