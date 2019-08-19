//teste unitário é testar uma unidade por exemplo uma função - isola totalmente uma função e testa APENAS ela
//quanto mais você isola essa função pra testar ela isolada do exterior é melhor porque testa de forma única
//é bom porque se eu mudar a função de alguma forma, o teste unitário avisa se der algum problema
//quando cria primeiro o teste é melhor porque a função no fim de ser criada já vai estar muito mais robusta do que se o teste fosse feito depois
//mocar é dar um comportamento artifical pra uma dependência que o método tenha

package matematica

import (
	"fmt"
	"strconv"
)

//Media é responsável por calcular a média kk
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
