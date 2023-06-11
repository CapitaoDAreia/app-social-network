package repositories

import "backend/internal/domain/entities"

type UsersRepository interface {
	CreateUser(user entities.User) (string, error)
	SearchUsers(usernameOrNickQuery string) ([]entities.User, error)
	SearchUser(requestID string) (entities.User, error)
	UpdateUser(ID string, user entities.User) (uint64, error)
	DeleteUser(ID string) (uint64, error)
	SearchUserByEmail(email string) (entities.User, error)
	Follow(followedID, followerID string) error
	UnFollow(followedID, followerID string) error
	SearchFollowersOfAnUser(userID string) ([]string, error)
	SearchWhoAnUserFollow(userID string) ([]string, error)
	SearchUserPassword(userID string) (string, error)
	UpdateUserPassword(requestUserId string, hashedNewPasswordStringed string) (uint64, error)
}
