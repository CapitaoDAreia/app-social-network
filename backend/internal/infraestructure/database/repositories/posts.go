package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/infraestructure/database"
	"backend/internal/infraestructure/database/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsRepository struct {
	collection *mongo.Collection
}

func NewPostsRepository(db *mongo.Database) *PostsRepository {
	collection := db.Collection(database.POSTS_COLLECTION)

	return &PostsRepository{
		collection: collection,
	}
}

func (repository PostsRepository) CreatePost(post entities.Post) (string, error) {

	newPost := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		AuthorID:   post.AuthorID,
		AuthorNick: post.AuthorNick,
		Likes:      []string{},
		CreatedAt:  time.Now(),
	}

	result, err := repository.collection.InsertOne(context.Background(), newPost)
	if err != nil {
		return "0", fmt.Errorf("Error on insert a new post: %s", err)
	}

	stringNewInsertedPostID := result.InsertedID.(primitive.ObjectID).Hex()

	return stringNewInsertedPostID, nil
}

// func (p PostsRepository) SearchPost(postID uint64) (entities.Post, error) {
// 	rows, err := p.db.Query(`
// 		 select p.*, u.nick from
// 		 posts p inner join users u
// 		 on u.id = p.authorId where p.id = ?
// 	`, postID)
// 	if err != nil {
// 		return entities.Post{}, err
// 	}
// 	defer rows.Close()

// 	var post entities.Post

// 	if rows.Next() {
// 		if err := rows.Scan(
// 			&post.ID,
// 			&post.Title,
// 			&post.Content,
// 			&post.AuthorID,
// 			&post.Likes,
// 			&post.CreatedAt,
// 			&post.AuthorNick,
// 		); err != nil {
// 			return entities.Post{}, err
// 		}
// 	}

// 	return post, nil
// }

// func (p PostsRepository) SearchPosts(tokenUserID uint64) ([]entities.Post, error) {
// 	rows, err := p.db.Query(`
// 		select distinct p.*, u.nick from posts p
// 		inner join users u on u.id = p.authorId
// 		inner join followers s on p.authorId = s.user_id
// 		where u.id = ? or s.follower_id = ?
// 		order by 1 desc`, tokenUserID, tokenUserID)
// 	if err != nil {
// 		return []entities.Post{}, err
// 	}
// 	defer rows.Close()

// 	var posts []entities.Post

// 	for rows.Next() {

// 		var post entities.Post

// 		if err := rows.Scan(
// 			&post.ID,
// 			&post.Title,
// 			&post.Content,
// 			&post.AuthorID,
// 			&post.Likes,
// 			&post.CreatedAt,
// 			&post.AuthorNick,
// 		); err != nil {
// 			return []entities.Post{}, err
// 		}

// 		posts = append(posts, post)
// 	}

// 	return posts, nil
// }

// func (p PostsRepository) UpdatePost(postRequestID uint64, updatedPost entities.Post) error {
// 	statement, err := p.db.Prepare(`update posts set title = ?, content = ? where id = ?`)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()

// 	if _, err := statement.Exec(
// 		updatedPost.Title,
// 		updatedPost.Content,
// 		postRequestID,
// 	); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p PostsRepository) DeletePost(postRequestID uint64) error {
// 	statement, err := p.db.Prepare(`delete from posts where id = ?`)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()

// 	if _, err := statement.Exec(postRequestID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p PostsRepository) SearchUserPosts(requestUserId uint64) ([]entities.Post, error) {
// 	rows, err := p.db.Query(`
// 		select p.*, u.nick from posts p
// 		join users u on u.id = p.authorId
// 		where p.authorId = ?
// 	`, requestUserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var posts []entities.Post

// 	for rows.Next() {
// 		var post entities.Post
// 		if err := rows.Scan(
// 			&post.ID,
// 			&post.Title,
// 			&post.Content,
// 			&post.AuthorID,
// 			&post.Likes,
// 			&post.CreatedAt,
// 			&post.AuthorNick,
// 		); err != nil {
// 			return nil, err
// 		}

// 		posts = append(posts, post)
// 	}

// 	return posts, nil
// }

func (repository PostsRepository) LikePost(postID, tokenUserID string) error {
	primitiveObjId, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"likes": tokenUserID,
		},
	}

	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error on UpdateOne to LikePost: %s", err)
	}

	return nil
}

// func (p PostsRepository) UnlikePost(postID uint64) error {
// 	statement, err := p.db.Prepare(`
// 		UPDATE posts SET likes =
// 		CASE
// 			WHEN likes > 0 THEN likes -1
// 		ELSE 0 END
// 		WHERE id = ?
// 	`)
// 	if err != nil {
// 		return err
// 	}
// 	defer statement.Close()

// 	if _, err := statement.Exec(postID); err != nil {
// 		return err
// 	}

// 	return nil
// }
