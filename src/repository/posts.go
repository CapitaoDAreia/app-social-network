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

func (p postsRepository) SearchPost(postID uint64) (models.Post, error) {
	rows, err := p.db.Query(`
		 select p.*, u.nick from
		 posts p inner join users u
		 on u.id = p.authorId where p.id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post

	if rows.Next() {
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
