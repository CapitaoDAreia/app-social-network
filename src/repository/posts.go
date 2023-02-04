package repository

import (
	"api-dvbk-socialNetwork/src/models"
	"database/sql"
)

type postsRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *postsRepository {
	return &postsRepository{db}
}

func (p postsRepository) CreatePost(post models.Post) (uint64, error) {
	statement, err := p.db.Prepare(`insert into posts (title, content, authorId) value(?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
