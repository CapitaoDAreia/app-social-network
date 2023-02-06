package routes

import (
	"api-dvbk-socialNetwork/src/controllers/loginController"
	"net/http"
)

var LoginRoute = Route{
	URI:        "/login",
	Method:     http.MethodPost,
	Controller: loginController.Login,
	NeedAuth:   false,
}
