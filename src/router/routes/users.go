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
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodPut,
		Controller: controllers.UpdateUser,
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodDelete,
		Controller: controllers.DeleteUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/follow",
		Method:     http.MethodPost,
		Controller: controllers.FollowUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/unfollow",
		Method:     http.MethodPost,
		Controller: controllers.UnFollowUser,
		NeedAuth:   true,
	},
}
