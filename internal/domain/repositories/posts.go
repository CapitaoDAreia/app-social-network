package repositories

import "api-dvbk-socialNetwork/internal/domain/entities"

type PostsRepository interface {
	CreatePost(post entities.Post) (uint64, error)
	SearchPost(postID uint64) (entities.Post, error)
	SearchPosts(tokenUserID uint64) ([]entities.Post, error)
	UpdatePost(postRequestID uint64, updatedPost entities.Post)
	DeletePost(postRequestID uint64) error
	SearchUserPosts(requestUserId uint64) ([]entities.Post, error)
	LikePost(postID uint64) error
	UnlikePost(postID uint64) error
}
