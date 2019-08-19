package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() //executa de forma paralela - bom quando tem varios testes aí ganha tempo de execução
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	//...
	t.Fail()
}
