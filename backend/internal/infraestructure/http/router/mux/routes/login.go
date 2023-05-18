package routes

import (
	"backend/internal/application/services"
	repository "backend/internal/infraestructure/database/repositories"
	"backend/internal/infraestructure/http/controllers/loginController"
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
