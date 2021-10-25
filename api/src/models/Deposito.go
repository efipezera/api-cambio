package models

import (
	"errors"
)

type Deposito struct {
	ID            uint64  `json:"id,omitempty"`
	ValorDeposito float64 `json:"valorDeposito,omitempty"`
}

//Preparar prepara o depósito para ser enviado ao banco de dados.
func (deposito *Deposito) Preparar() error {
	if erro := deposito.validar(); erro != nil {
		return erro
	}
	return nil
}

//validar valida se o depósito foi feito de forma correta.
func (deposito *Deposito) validar() error {
	if deposito.ValorDeposito <= 0 {
		return errors.New("o valor do depósito é obrigatório e deve ser acima de 0")
	}
	return nil
}
