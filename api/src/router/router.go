package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//GerarRouter retorna um roteador com as rotas configuradas.
func GerarRouter() *mux.Router {
	roteador := mux.NewRouter()
	return rotas.ConfigurarRotas(roteador)
}
