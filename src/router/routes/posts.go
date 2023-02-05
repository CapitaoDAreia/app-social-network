package routes

import (
	"api-dvbk-socialNetwork/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:        "/post",
		Method:     http.MethodPost,
		Controller: controllers.CreatePost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts",
		Method:     http.MethodGet,
		Controller: controllers.GetPosts,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodGet,
		Controller: controllers.GetPost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodPut,
		Controller: controllers.UpdatePost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodDelete,
		Controller: controllers.DeletePost,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/posts",
		Method:     http.MethodGet,
		Controller: controllers.GetUserPosts,
		NeedAuth:   true,
	},
}
