package main

import "math"

//inicializando com letra maiúscula é PÚBLICO (visibilidade dentro e fora do pacote)
//não pode ter funções duplicadas em vários arquivos dentro de um pacote porque compila e vira um único pacote
//iniciando com letra minúscula é PRIVADO (visível apenas dentro do pacote)

//Ponto - público
//ponto ou _Ponto - privado

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2)) //pow é elevado a potência e sqrt é raiz quadrada
}
