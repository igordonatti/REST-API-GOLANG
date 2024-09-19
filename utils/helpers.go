package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Converte (`marshaliza`) o `payload` para JSON
	response, _ := json.Marshal(payload)

	// Define o cabeçalho da resposta HTTP.
	w.Header().Set("Content-Type", "application/json")
	// Define o código de status HTTP da resposta
	w.WriteHeader(code)
	// Escreve a resposta JSON
	w.Write(response)
}
