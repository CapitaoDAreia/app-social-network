package services

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type PostServices interface {
	CreatePost(post entities.Post) (string, error)
	SearchPost(postID string) (entities.Post, error)
	SearchPosts(tokenUserID string) ([]entities.Post, error)
	UpdatePost(postRequestID string, updatedPost entities.Post) (uint64, error)
	DeletePost(postRequestID string) (uint64, error)
	SearchUserPosts(requestUserId string) ([]entities.Post, error)
	LikePost(postID, tokenUserID string) error
	UnlikePost(postID, tokenUserID string) error
}

type postServices struct {
	postsRepository repositories.PostsRepository
}

func NewPostsServices(postsRepository repositories.PostsRepository) *postServices {
	return &postServices{postsRepository}
}

func (service *postServices) CreatePost(post entities.Post) (string, error) {
	ID, err := service.postsRepository.CreatePost(post)
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (service *postServices) SearchPost(postID string) (entities.Post, error) {
	post, err := service.postsRepository.SearchPost(postID)
	if err != nil {
		return entities.Post{}, err
	}

	return post, nil
}

func (service *postServices) SearchPosts(tokenUserID string) ([]entities.Post, error) {
	posts, err := service.postsRepository.SearchPosts(tokenUserID)
	if err != nil {
		return []entities.Post{}, err
	}

	return posts, nil
}

func (service *postServices) UpdatePost(postRequestID string, updatedPost entities.Post) (uint64, error) {
	modifiedCount, err := service.postsRepository.UpdatePost(postRequestID, updatedPost)
	if err != nil {
		return 0, err
	}

	return modifiedCount, err
}

func (service *postServices) DeletePost(postRequestID string) (uint64, error) {
	modifiedCount, err := service.postsRepository.DeletePost(postRequestID)
	if err != nil {
		return 0, err
	}

	return modifiedCount, nil
}

func (service *postServices) SearchUserPosts(requestUserId string) ([]entities.Post, error) {
	posts, err := service.postsRepository.SearchUserPosts(requestUserId)
	if err != nil {
		return []entities.Post{}, err
	}

	return posts, nil

}

func (service *postServices) LikePost(postID, tokenUserID string) error {
	err := service.postsRepository.LikePost(postID, tokenUserID)
	if err != nil {
		return err
	}

	return nil
}

func (service *postServices) UnlikePost(postID, tokenUserID string) error {
	err := service.postsRepository.UnlikePost(postID, tokenUserID)
	if err != nil {
		return err
	}

	return nil
}
