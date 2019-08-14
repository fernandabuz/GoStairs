package main

import "fmt"

func main() {
	array1 := ([3]int{1, 2, 3})
	var slice1 []int

	//array1 = append(array1, 4, 5, 6) não dá pra fazer append em um array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) //primeiro item recebe o valor copiado e segundo item é a fonte
	fmt.Println(slice2)  //como tem dois itens no slice2, ele copia os dois primeiros do slice1 - 4, 5
}
