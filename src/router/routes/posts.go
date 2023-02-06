package routes

import (
	"api-dvbk-socialNetwork/src/controllers/postsController"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:        "/post",
		Method:     http.MethodPost,
		Controller: postsController.CreatePost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts",
		Method:     http.MethodGet,
		Controller: postsController.GetPosts,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodGet,
		Controller: postsController.GetPost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodPut,
		Controller: postsController.UpdatePost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodDelete,
		Controller: postsController.DeletePost,
		NeedAuth:   true,
	},
	{
		URI:        "/users/{userId}/posts",
		Method:     http.MethodGet,
		Controller: postsController.GetUserPosts,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}/like",
		Method:     http.MethodPost,
		Controller: postsController.LikePost,
		NeedAuth:   true,
	},
	{
		URI:        "/posts/{postId}/unlike",
		Method:     http.MethodPost,
		Controller: postsController.UnlikePost,
		NeedAuth:   true,
	},
}
