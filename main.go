package main

import (
	"github.com/igordonatti/REST-API/startup"
)

func main() {
	startup.Server()
}

// To Do
// Fazer verificação para o não cadastro de dois emails
// Autenticação
// Autorização
// Corrigir o envio de erro para usuario ja existente
// Criptografia da senha
