package routes

import (
	"backend/internal/application/services"
	repository "backend/internal/infraestructure/database/repositories"
	"backend/internal/infraestructure/http/controllers/loginController"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigLoginRoutes(db *mongo.Database) Route {

	repository := repository.NewUsersRepository(db)
	services := services.NewUsersServices(repository)
	controllers := loginController.NewLoginController(services)

	var LoginRoute = Route{
		URI:        "/login",
		Method:     http.MethodPost,
		Controller: controllers.Login,
		NeedAuth:   false,
	}
	return LoginRoute
}
