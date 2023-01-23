package routes

import "net/http"

//Routes is a struct that represents API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}
