package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarVariaveisDeAmbiente()
	roteador := router.GerarRouter()
	fmt.Printf("Rodando a API na porta :%d \n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), roteador))
}
