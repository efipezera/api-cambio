package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON retorna uma resposta em JSON para a requisição.
func JSON(rw http.ResponseWriter, statusCode int, dados interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(rw).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(rw http.ResponseWriter, statusCode int, erro error) {
	JSON(rw, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
