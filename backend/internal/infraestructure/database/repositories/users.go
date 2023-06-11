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

type UsersRepository struct {
	collection *mongo.Collection
}

// NewUserRepository Receives a database opened in controller and instances it in users struct.
func NewUsersRepository(db *mongo.Database) *UsersRepository {

	collection := db.Collection(database.USERS_COLLECTION)
	return &UsersRepository{
		collection: collection,
	}
}

// CreateUser Creates a user on database.
// This is a method of users struct.
func (repository UsersRepository) CreateUser(user entities.User) (string, error) {

	newUser := models.User{
		Nick:      user.Nick,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Followers: []string{},
		Following: []string{},
		Posts:     []string{},
		CreatedAt: time.Now(),
	}

	result, err := repository.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return "0", fmt.Errorf("Error on insert a new user: %s", err)
	}

	stringNewInsertedUserID := result.InsertedID.(primitive.ObjectID).Hex()

	return stringNewInsertedUserID, nil
}

// Search for users by username or nick
func (repository UsersRepository) SearchUsers(usernameOrNickQuery string) ([]entities.User, error) {

	filter := bson.M{
		"$or": []bson.M{
			{"nick": bson.M{
				"$regex":   usernameOrNickQuery,
				"$options": "i",
			}},
			{"username": bson.M{
				"$regex":   usernameOrNickQuery,
				"$options": "i",
			}},
		},
	}

	result, err := repository.collection.Find(context.Background(), filter)
	if err != nil {
		return []entities.User{}, fmt.Errorf("Error on SearchUsers: %s", err)
	}

	var users []entities.User

	err = result.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository UsersRepository) SearchUser(requestID string) (entities.User, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(requestID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	var user entities.User

	result := repository.collection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err().Error() == mongo.ErrNoDocuments.Error() {
			return entities.User{}, fmt.Errorf("User not found")
		}
	}

	if err := result.Decode(&user); err != nil {
		return entities.User{}, fmt.Errorf("Error on decode SearchUser: %s", err)
	}

	return user, nil
}

func (repository UsersRepository) SearchUserByEmail(email string) (entities.User, error) {

	filter := bson.M{
		"email": email,
	}

	var user entities.User

	result := repository.collection.FindOne(context.Background(), filter)

	if err := result.Decode(&user); err != nil {
		return entities.User{}, fmt.Errorf("Error on FindOne to SearchUserByEmail: %s", err)
	}

	return user, nil
}

func (repository UsersRepository) UpdateUser(ID string, user entities.User) (uint64, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	updateOptions := bson.M{
		"$set": bson.M{
			"nick":      user.Nick,
			"username":  user.Username,
			"email":     user.Email,
			"updatedAt": time.Now(),
		},
	}

	result, err := repository.collection.UpdateOne(context.Background(), filter, updateOptions)
	if err != nil {
		return 0, fmt.Errorf("Error on UpdateOne: %s", err)
	}

	modifiedCount := uint64(result.ModifiedCount)

	fmt.Println(modifiedCount)

	return modifiedCount, nil
}

func (repository UsersRepository) DeleteUser(ID string) (uint64, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	result, err := repository.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, fmt.Errorf("Error on UpdateOne to DeleteUser: %s", err)
	}

	modifiedCount := uint64(result.DeletedCount)

	return modifiedCount, nil
}

func (repository UsersRepository) Follow(followedID, followerID string) error {

	primitiveObjId, _ := primitive.ObjectIDFromHex(followedID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	update := bson.M{
		"$push": bson.M{
			"followers": followerID,
		},
	}

	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error on UpdateOne to Follow: %s", err)
	}

	return nil
}

func (repository UsersRepository) UnFollow(followedID, followerID string) error {

	primitiveObjId, _ := primitive.ObjectIDFromHex(followedID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	update := bson.M{
		"$pull": bson.M{
			"followers": followerID,
		},
	}

	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("Error on UpdateOne to UnFollow: %s", err)
	}

	return nil
}

func (repository UsersRepository) SearchFollowersOfAnUser(userID string) ([]string, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	var user entities.User

	result := repository.collection.FindOne(context.Background(), filter)

	if err := result.Decode(&user); err != nil {
		return []string{}, fmt.Errorf("Erro on decode user to SearchFollowersOfAnUser: %s", err)
	}

	return user.Followers, nil
}

func (repository UsersRepository) SearchWhoAnUserFollow(userID string) ([]string, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	var user entities.User

	result := repository.collection.FindOne(context.Background(), filter)

	if err := result.Decode(&user); err != nil {
		return []string{}, fmt.Errorf("Erro on decode user to SearchWhoAnUserFollow: %s", err)
	}

	return user.Following, nil
}

func (repository UsersRepository) SearchUserPassword(userID string) (string, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	var user entities.User

	result := repository.collection.FindOne(context.Background(), filter)

	if err := result.Decode(&user); err != nil {
		return "", fmt.Errorf("Erro on decode user to SearchUserPassword: %s", err)
	}

	return user.Password, nil
}

func (repository UsersRepository) UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error) {

	primitiveObjId, _ := primitive.ObjectIDFromHex(requestUserId)

	filter := bson.M{
		"_id": primitiveObjId,
	}

	updateOptions := bson.M{
		"$set": bson.M{
			"password": hashedNewPasswordStringed,
		},
	}

	result, err := repository.collection.UpdateOne(context.Background(), filter, updateOptions)
	if err != nil {
		return 0, fmt.Errorf("Error on UpdateOne to UpdateUserPassword: %s", err)
	}

	modifiedCount := uint64(result.ModifiedCount)

	return modifiedCount, nil
}
