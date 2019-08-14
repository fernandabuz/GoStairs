package main

import "fmt"

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo (se o retorno é nomeado não precisa citar aqui)
}

func main() {
	x, y := troca(5, 8)
	fmt.Println(x, y)

	segundo, primeiro := troca(9, 1)
	fmt.Println(segundo, primeiro)
}
