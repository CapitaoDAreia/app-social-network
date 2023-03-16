package routes

import (
	"api-dvbk-socialNetwork/internal/application/services"
	repository "api-dvbk-socialNetwork/internal/infraestructure/database/repositories"
	"api-dvbk-socialNetwork/internal/infraestructure/http/controllers/postsController"
	"database/sql"
	"net/http"
)

func ConfigPostsRoutes(db *sql.DB) []Route {

	repository := repository.NewPostsRepository(db)
	services := services.NewPostsServices(*repository)
	controllers := postsController.NewPostsController(services)

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
		{
			URI:        "/posts/{postId}/like",
			Method:     http.MethodPost,
			Controller: controllers.LikePost,
			NeedAuth:   true,
		},
		{
			URI:        "/posts/{postId}/unlike",
			Method:     http.MethodPost,
			Controller: controllers.UnlikePost,
			NeedAuth:   true,
		},
	}

	return postRoutes
}
