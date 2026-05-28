package respostas

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
		return
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error){
	JSON(w, statusCode, struct {
	Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}