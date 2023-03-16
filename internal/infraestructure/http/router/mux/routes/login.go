package routes

import (
	"api-dvbk-socialNetwork/internal/application/services"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/loginController"
	"database/sql"
	"net/http"
)

func ConfigLoginRoutes(db *sql.DB) Route {

	reposiry := repository.NewUsersRepository(db)
	services := services.NewUsersServices(reposiry)
	controllers := loginController.NewLoginController(services)

	var LoginRoute = Route{
		URI:        "/login",
		Method:     http.MethodPost,
		Controller: controllers.Login,
		NeedAuth:   false,
	}
	return LoginRoute
}
