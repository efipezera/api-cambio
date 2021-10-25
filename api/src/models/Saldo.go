package models

type Saldo struct {
	ValorTotal    float64 `json:"saldoTotal,omitempty"`
	ValorMoedaUSD float64 `json:"saldoEmDolar,omitempty"`
	ValorMoedaEUR float64 `json:"saldoEmEuro,omitempty"`
	ValorMoedaGBP float64 `json:"saldoEmLibra,omitempty"`
}
