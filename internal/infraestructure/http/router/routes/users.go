package routes

import (
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/usersController"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:        "/users",
		Method:     http.MethodPost,
		Controller: usersController.CreateUser,
		NeedAuth:   false,
	},

	{
		URI:        "/users",
		Method:     http.MethodGet,
		Controller: usersController.GetUsers,
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodGet,
		Controller: usersController.GetUser,
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodPut,
		Controller: usersController.UpdateUser,
		NeedAuth:   true,
	},

	{
		URI:        "/users/{userId}",
		Method:     http.MethodDelete,
		Controller: usersController.DeleteUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/follow",
		Method:     http.MethodPost,
		Controller: usersController.FollowUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/unfollow",
		Method:     http.MethodPost,
		Controller: usersController.UnFollowUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/followers",
		Method:     http.MethodGet,
		Controller: usersController.GetFollowersOfAnUser,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/following",
		Method:     http.MethodGet,
		Controller: usersController.GetWhoAnUserFollow,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/update-password",
		Method:     http.MethodPost,
		Controller: usersController.UpdateUserPassword,
		NeedAuth:   true,
	},
}
