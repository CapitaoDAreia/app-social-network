package repositories

import "backend/internal/domain/entities"

type PostsRepository interface {
	CreatePost(post entities.Post) (string, error)
	SearchPost(postID string) (entities.Post, error)
	SearchPosts(tokenUserID string) ([]entities.Post, error)
	UpdatePost(postRequestID string, updatedPost entities.Post) (uint64, error)
	DeletePost(postRequestID string) (uint64, error)
	SearchUserPosts(requestUserId string) ([]entities.Post, error)
	LikePost(postID, tokenUserID string) error
	UnlikePost(postID, tokenUserID string) error
}
