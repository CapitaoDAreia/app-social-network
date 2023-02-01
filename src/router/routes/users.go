package routes

import (
	"api-dvbk-socialNetwork/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:        "/users",
		Method:     http.MethodPost,
		Controller: controllers.CreateUser,
		NeedAuth:   false,
	},

	{
		URI:        "/users",
		Method:     http.MethodGet,
		Controller: controllers.SearchUsers,
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodGet,
		Controller: controllers.SearchUser,
		NeedAuth:   false,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodPut,
		Controller: controllers.UpdateUser,
		NeedAuth:   false,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodDelete,
		Controller: controllers.DeleteUser,
		NeedAuth:   false,
	},
}
