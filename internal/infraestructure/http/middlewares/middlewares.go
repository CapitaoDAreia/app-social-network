package middlewares

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/auth"
	"api-dvbk-socialNetwork/internal/infraestructure/http/responses"
	"log"
	"net/http"
)

func Logger(receivedFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Logger info: %v\n %v\n %v\n", r.Method, r.RequestURI, r.Host)
		receivedFunc(w, r)
	}
}

func Authenticate(receivedFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.FormatResponseToCustomError(w, http.StatusUnauthorized, err)
			return
		}
		receivedFunc(w, r)
	}
}
