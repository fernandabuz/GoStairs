package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota { //no switch eu nao tenho uma expressão que retorna verdadeiro ou falso
	case 10:
		fallthrough //continua descendo para o próximo case - go não usa break porque se encontra o valor sai do switch
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.0))
}
