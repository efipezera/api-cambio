package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasOperacoes = []Rota{
	{
		URI:    "/depositar",
		Metodo: http.MethodPost,
		Funcao: controllers.FazerDeposito,
	},
	{
		URI:    "/depositos",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarDepositos,
	},
	{
		URI:    "/saldo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSaldoTotal,
	},
}
