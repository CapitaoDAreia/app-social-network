package services

import (
	"backend/internal/domain/entities"
	repository "backend/internal/infraestructure/database/repositories"
)

type PostServices interface {
	CreatePost(post entities.Post) (uint64, error)
	SearchPost(postID uint64) (entities.Post, error)
	SearchPosts(tokenUserID uint64) ([]entities.Post, error)
	UpdatePost(postRequestID uint64, updatedPost entities.Post) error
	DeletePost(postRequestID uint64) error
	SearchUserPosts(requestUserId uint64) ([]entities.Post, error)
	LikePost(postID uint64) error
	UnlikePost(postID uint64) error
}

type postServices struct {
	postsRepository repository.PostsRepository
}

func NewPostsServices(postsRepository repository.PostsRepository) *postServices {
	return &postServices{postsRepository}
}

func (service *postServices) CreatePost(post entities.Post) (uint64, error) {
	ID, err := service.postsRepository.CreatePost(post)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (service *postServices) SearchPost(postID uint64) (entities.Post, error) {
	post, err := service.postsRepository.SearchPost(postID)
	if err != nil {
		return entities.Post{}, err
	}

	return post, nil
}

func (service *postServices) SearchPosts(tokenUserID uint64) ([]entities.Post, error) {
	posts, err := service.postsRepository.SearchPosts(tokenUserID)
	if err != nil {
		return []entities.Post{}, err
	}

	return posts, nil
}

func (service *postServices) UpdatePost(postRequestID uint64, updatedPost entities.Post) error {
	err := service.postsRepository.UpdatePost(postRequestID, updatedPost)
	if err != nil {
		return err
	}

	return nil
}

func (service *postServices) DeletePost(postRequestID uint64) error {
	err := service.postsRepository.DeletePost(postRequestID)
	if err != nil {
		return err
	}

	return nil
}

func (service *postServices) SearchUserPosts(requestUserId uint64) ([]entities.Post, error) {
	posts, err := service.postsRepository.SearchUserPosts(requestUserId)
	if err != nil {
		return []entities.Post{}, err
	}

	return posts, nil

}

func (service *postServices) LikePost(postID uint64) error {
	err := service.postsRepository.LikePost(postID)
	if err != nil {
		return err
	}

	return nil
}

func (service *postServices) UnlikePost(postID uint64) error {
	err := service.postsRepository.UnlikePost(postID)
	if err != nil {
		return err
	}

	return nil
}
