package repositorios

import "testing"

func TestCalculaCambio(t *testing.T) {
	esperado := 71.25940567147339
	recebido := 500 / calculaCambio(5.552)
	if recebido != esperado {
		t.Errorf("Função esperava: %f, recebeu: %f", esperado, recebido)
	}
}
