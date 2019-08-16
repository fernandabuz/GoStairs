package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"` //identificador começando com letra maiúscula tem significado de ser público e minúscula é privado
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", 99.0, []string{"promo", "eletro"}}
	p1Json, _ := json.Marshal(p1) //retorna o json e o erro - sempre dois
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "Celular", "preco": 1800.00, "tags": ["baratinho", "novinho"]}`
	json.Unmarshal([]byte(jsonString), &p2) //primeiro slice de byte passando referencia do p2 - para que ele leia o json e sete os valores pra dentro do p2(por referencia &)
	fmt.Println(p2.Tags[1])
}
