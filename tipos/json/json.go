package main

type produto struct {
	ID    int      `json:"id"` //identificador começando com letra maiúscula tem significado de ser público e minúscula é privado
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

}
