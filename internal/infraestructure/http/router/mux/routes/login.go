package routes

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/loginController"
	"net/http"
)

var LoginRoute = Route{
	URI:        "/login",
	Method:     http.MethodPost,
	Controller: loginController.Login,
	NeedAuth:   false,
}
