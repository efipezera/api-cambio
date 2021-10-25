package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//FazerDeposito faz o depósito de um valor.
func FazerDeposito(rw http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}
	var deposito models.Deposito
	if erro = json.Unmarshal(corpoRequisicao, &deposito); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}
	if erro = deposito.Preparar(); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	deposito.ID, erro = repositorio.CriarDeposito(deposito)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(rw, http.StatusCreated, deposito)
	fmt.Printf("Foi depositado o valor de R$%.2f \n", deposito.ValorDeposito)
}

//BuscarDepositos lista todos os depósitos feitos.
func BuscarDepositos(rw http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	depositos, erro := repositorio.BuscarDepositos()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(rw, http.StatusOK, depositos)
}

//BuscarSaldoTotal busca o saldo total.
func BuscarSaldoTotal(rw http.ResponseWriter, r *http.Request) {
	moeda := r.URL.Query().Get("moeda")
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	saldoTotal, erro := repositorio.BuscarSaldoTotal(moeda)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(rw, http.StatusOK, saldoTotal)
}
