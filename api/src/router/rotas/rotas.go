package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

func ConfigurarRotas(roteador *mux.Router) *mux.Router {
	rotas := rotasOperacoes
	for _, rota := range rotas {
		roteador.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return roteador
}
