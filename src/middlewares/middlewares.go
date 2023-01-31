package middlewares

import (
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
		receivedFunc(w, r)
	}
}
