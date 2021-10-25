package repositorios

import (
	"api/src/models"
	"database/sql"
	"log"
)

type Depositos struct {
	db *sql.DB
}

//NovoRepositorioDeDepositos cria um repositório de depósitos.
func NovoRepositorioDeDepositos(db *sql.DB) *Depositos {
	return &Depositos{db}
}

//CriarDeposito insere um depósito no banco de dados.
func (repositorio Depositos) CriarDeposito(deposito models.Deposito) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO depositos (valorDeposito)VALUES (?)")
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()
	resultado, erro := statement.Exec(deposito.ValorDeposito)
	if erro != nil {
		return 0, nil
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(ultimoIDInserido), nil
}

//BuscarDepositos traz todos os depósitos feitos.
func (repositorio Depositos) BuscarDepositos() ([]models.Deposito, error) {
	linhas, erro := repositorio.db.Query("select valorDeposito from depositos")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var depositos []models.Deposito
	for linhas.Next() {
		var deposito models.Deposito
		if erro = linhas.Scan(&deposito.ValorDeposito); erro != nil {
			return nil, erro
		}
		depositos = append(depositos, deposito)
	}
	return depositos, nil
}

//BuscarSaldoTotal soma os depositos e traz o saldo total.
func (repositorio Depositos) BuscarSaldoTotal(moeda string) (models.Saldo, error) {
	linhas, erro := repositorio.db.Query("select valorDeposito from depositos")
	if erro != nil {
		return models.Saldo{}, erro
	}
	defer linhas.Close()
	valoresEmSlice := make([]float64, 3)
	for linhas.Next() {
		var deposito models.Deposito
		if erro = linhas.Scan(&deposito.ValorDeposito); erro != nil {
			return models.Saldo{}, erro
		}
		valoresEmSlice = append(valoresEmSlice, deposito.ValorDeposito)
	}
	var valorTotalEmBRL float64
	for index, _ := range valoresEmSlice {
		valorTotalEmBRL += valoresEmSlice[index]
	}
	var valores models.Saldo
	switch moeda {
	case "USD":
		valores.ValorMoedaUSD = valorTotalEmBRL / calculaCambio(5.552)
	case "EUR":
		valores.ValorMoedaEUR = valorTotalEmBRL / calculaCambio(6.444)
	case "GBP":
		valores.ValorMoedaGBP = valorTotalEmBRL / calculaCambio(7.638)
	default:
		valores.ValorTotal = valorTotalEmBRL
	}
	return valores, nil
}

//calculaCambio faz as operações de calcular Spread, IOF e Taxa de Câmbio e retorna o valor real da moeda com impostos.
func calculaCambio(valorMoeda float64) float64 {
	if valorMoeda <= 0 {
		log.Fatal("o valor da moeda não pode ser menor ou igual a zero")
	}
	spread := 0.04
	iof := 0.0638
	taxaCambio := 0.16
	valorFinal := valorMoeda + (valorMoeda * spread) + (valorMoeda * iof) + (valorMoeda * taxaCambio)
	return valorFinal
}
