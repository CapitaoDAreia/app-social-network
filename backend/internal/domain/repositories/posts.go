package repositories

import "backend/internal/domain/entities"

type PostsRepository interface {
	CreatePost(post entities.Post) (string, error)
	SearchPost(postID string) (entities.Post, error)
	SearchPosts(tokenUserID string) ([]entities.Post, error)
	// UpdatePost(postRequestID uint64, updatedPost entities.Post)
	// DeletePost(postRequestID uint64) error
	// SearchUserPosts(requestUserId uint64) ([]entities.Post, error)
	LikePost(postID, tokenUserID string) error
	// UnlikePost(postID uint64) error
}
