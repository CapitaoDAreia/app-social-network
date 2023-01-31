package routes

import (
	"api-dvbk-socialNetwork/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	URI:        "/login",
	Method:     http.MethodPost,
	Controller: controllers.Login,
	NeedAuth:   false,
}
