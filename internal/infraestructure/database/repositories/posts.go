package repository

import (
	"api-dvbk-socialNetwork/internal/infraestructure/database/models"
	"database/sql"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db}
}

func (p PostsRepository) CreatePost(post models.Post) (uint64, error) {
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

func (p PostsRepository) SearchPost(postID uint64) (models.Post, error) {
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

func (p PostsRepository) SearchPosts(tokenUserID uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.authorId
		inner join followers s on p.authorId = s.user_id
		where u.id = ? or s.follower_id = ?
		order by 1 desc`, tokenUserID, tokenUserID)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {

		var post models.Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p PostsRepository) UpdatePost(postRequestID uint64, updatedPost models.Post) error {
	statement, err := p.db.Prepare(`update posts set title = ?, content = ? where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(
		updatedPost.Title,
		updatedPost.Content,
		postRequestID,
	); err != nil {
		return err
	}

	return nil
}

func (p PostsRepository) DeletePost(postRequestID uint64) error {
	statement, err := p.db.Prepare(`delete from posts where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postRequestID); err != nil {
		return err
	}

	return nil
}

func (p PostsRepository) SearchUserPosts(requestUserId uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
		select p.*, u.nick from posts p
		join users u on u.id = p.authorId
		where p.authorId = ?
	`, requestUserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p PostsRepository) LikePost(postID uint64) error {
	statement, err := p.db.Prepare(`update posts set likes = likes + 1 where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (p PostsRepository) UnlikePost(postID uint64) error {
	statement, err := p.db.Prepare(`
		UPDATE posts SET likes = 
		CASE 
			WHEN likes > 0 THEN likes -1
		ELSE 0 END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
