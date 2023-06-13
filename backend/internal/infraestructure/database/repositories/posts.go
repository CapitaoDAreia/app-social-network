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

func (repository PostsRepository) UpdatePost(postRequestID string, updatedPost entities.Post) (uint64, error) {

	postRequestPrimitiveID, _ := primitive.ObjectIDFromHex(postRequestID)

	filter := bson.M{
		"_id": postRequestPrimitiveID,
	}

	update := bson.M{
		"$set": bson.M{
			"title":     updatedPost.Title,
			"content":   updatedPost.Content,
			"updatedAt": time.Now(),
		},
	}

	result, err := repository.postsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, fmt.Errorf("Error on UpdateOne to UpdatePost: %s", err)
	}

	modifiedCount := uint64(result.ModifiedCount)

	return modifiedCount, nil
}

func (repository PostsRepository) DeletePost(postRequestID string) (uint64, error) {

	postRequestPrimitiveID, _ := primitive.ObjectIDFromHex(postRequestID)

	postsFilter := bson.M{
		"_id": postRequestPrimitiveID,
	}

	var post entities.Post

	searchPostResult := repository.postsCollection.FindOne(context.Background(), postsFilter)
	if err := searchPostResult.Decode(&post); err != nil {
		return 0, fmt.Errorf("Error on Decode to DeletePost: %s", err)
	}

	userPrimitiveID, _ := primitive.ObjectIDFromHex(post.AuthorID)

	usersFilter := bson.M{
		"_id": userPrimitiveID,
	}

	usersUpdateCriteria := bson.M{
		"$pull": bson.M{
			"posts": postRequestID,
		},
	}

	_, err := repository.usersCollection.UpdateOne(context.Background(), usersFilter, usersUpdateCriteria)
	if err != nil {
		return 0, fmt.Errorf("Error on UpdateMany to DeletePost: %s", err)
	}

	postDeleteResult, err := repository.postsCollection.DeleteOne(context.Background(), postsFilter)
	if err != nil {
		return 0, fmt.Errorf("Error on DeleteOne to DeletePost: %s", err)
	}

	deletedCountInPostsCollection := uint64(postDeleteResult.DeletedCount)

	return deletedCountInPostsCollection, nil
}

func (repository PostsRepository) SearchUserPosts(requestUserId string) ([]entities.Post, error) {

	filter := bson.M{
		"authorid": requestUserId,
	}

	var posts []entities.Post

	result, err := repository.postsCollection.Find(context.Background(), filter)
	if err != nil {
		return []entities.Post{}, fmt.Errorf("Error on Find to SearchUserPosts: %s", err)
	}

	if err := result.All(context.Background(), &posts); err != nil {
		return []entities.Post{}, err
	}

	return posts, nil
}

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

func (repository PostsRepository) UnlikePost(postID, tokenUserID string) error {
	primitiveObjId, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	update := bson.M{
		"$pull": bson.M{
			"likes": tokenUserID,
		},
	}

	_, err := repository.postsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error on UpdateOne to UnlikePost: %s", err)
	}

	return nil
}
