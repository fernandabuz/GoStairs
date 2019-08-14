package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice não é um array! ele define um pedaço de um array
	s2 := a2[1:3] //do 1 ao 3 sem incluir o 3
	fmt.Println(a2, s2)

	s3 := a2[:2] //novo slice mas aponta pro mesmo array
	fmt.Println(a2, s3)

	//slice tem tamanho e um ponteiro para um elemento de um array

}
