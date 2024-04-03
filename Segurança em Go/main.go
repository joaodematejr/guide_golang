package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Função de autenticação básica
func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrai as credenciais do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// Se não houver cabeçalho de autorização, retorna um erro de não autorizado
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// Verifica se o cabeçalho de autorização está no formato correto
		auth := strings.SplitN(authHeader, " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		// Decodifica as credenciais do cabeçalho
		payload, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		// Verifica se as credenciais são válidas (nesse exemplo, usuário:senha é "admin:password")
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 || credentials[0] != "admin" || credentials[1] != "password" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// Se as credenciais forem válidas, chama o próximo manipulador
		next.ServeHTTP(w, r)
	})
}

// Manipulador protegido por autenticação básica
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página protegida! Acesso permitido."))
}

func main() {
	// Cria um servidor HTTP
	server := http.NewServeMux()

	// Registra o middleware de autenticação básica para a rota protegida
	server.Handle("/protected", basicAuthMiddleware(http.HandlerFunc(protectedHandler)))

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", server)
}
