package main

import (
	"api-dvbk-socialNetwork/src/router"
	"log"
	"net/http"
)

func main() {
	const PORT string = `:5000`
	r := router.Generate()

	log.Fatal(http.ListenAndServe(PORT, r))
}
