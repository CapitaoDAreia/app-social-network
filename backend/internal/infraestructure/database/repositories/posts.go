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
	postsCollection *mongo.Collection
	usersCollection *mongo.Collection
}

func NewPostsRepository(db *mongo.Database) *PostsRepository {
	postsCollection := db.Collection(database.POSTS_COLLECTION)
	usersCollection := db.Collection(database.USERS_COLLECTION)

	return &PostsRepository{
		postsCollection: postsCollection,
		usersCollection: usersCollection,
	}
}

// Register a post in the database
func (repository PostsRepository) CreatePost(post entities.Post) (string, error) {

	newPost := models.Post{
		Title:      post.Title,
		Content:    post.Content,
		AuthorID:   post.AuthorID,
		AuthorNick: post.AuthorNick,
		Likes:      []string{},
		CreatedAt:  time.Now(),
	}

	postsResult, err := repository.postsCollection.InsertOne(context.Background(), newPost)
	if err != nil {
		return "0", fmt.Errorf("Error on insert a new post into posts collection: %s", err)
	}

	stringNewInsertedPostID := postsResult.InsertedID.(primitive.ObjectID).Hex()

	primitiveAuthorID, err := primitive.ObjectIDFromHex(newPost.AuthorID)

	updateAuthorFilter := bson.M{
		"_id": primitiveAuthorID,
	}

	updateAuthorCriteria := bson.M{
		"$addToSet": bson.M{
			"posts": stringNewInsertedPostID,
		},
	}

	_, err = repository.usersCollection.UpdateOne(context.Background(), updateAuthorFilter, updateAuthorCriteria)
	if err != nil {
		return "0", fmt.Errorf("Post was created but an error was occur on insert a new post into user document: %s", err)
	}

	return stringNewInsertedPostID, nil
}

// Search for a post in the database based on a post ID
func (repository PostsRepository) SearchPost(postID string) (entities.Post, error) {

	primitivePostID, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.M{
		"_id": primitivePostID,
	}

	var post entities.Post

	postsResult := repository.postsCollection.FindOne(context.Background(), filter)

	if err := postsResult.Decode(&post); err != nil {
		return entities.Post{}, fmt.Errorf("Error on FindOne to SearchPost: %s", err)
	}

	return post, nil
}

// Search for posts in the database based on an user ID
func (repository PostsRepository) SearchPosts(tokenUserID string) ([]entities.Post, error) {

	filter := bson.M{
		"authorid": tokenUserID,
	}

	postsResult, err := repository.postsCollection.Find(context.Background(), filter)
	if err != nil {
		return []entities.Post{}, fmt.Errorf("Error on Find to SearchPosts")
	}

	var posts []entities.Post

	if err := postsResult.All(context.Background(), &posts); err != nil {
		return []entities.Post{}, err
	}

	return posts, nil
}

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

	_, err := repository.postsCollection.UpdateOne(context.Background(), filter, update)
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
