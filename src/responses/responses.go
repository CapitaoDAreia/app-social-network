package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Format response to JSON format and handle response statusCode
func FormatResponseToJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Format error to JSON based on FormatResponseToJSON
func FormatResponseToCustomError(w http.ResponseWriter, statusCode int, err error) {
	FormatResponseToJSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}
