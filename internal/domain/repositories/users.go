package repositories

import "api-dvbk-socialNetwork/internal/domain/entities"

type UsersRepository interface {
	CreateUser(user entities.User) (uint64, error)
	SearchUsers(usernameOrNickQuery string) ([]entities.User, error)
	SearchUser(requestID uint64) (entities.User, error)
	UpdateUser(ID uint64, user entities.User) error
	DeleteUser(ID uint64) error
	SearchUserByEmail(email string) (entities.User, error)
	Follow(followedID, followerID uint64) error
	UnFollow(followedID, followerID uint64) error
	SearchFollowersOfnAnUser(userID uint64) ([]entities.User, error)
	SearchWhoAnUserFollow(userID uint64) ([]entities.User, error)
	SearchUserPassword(userID uint64) (string, error)
	UpdateUserPassword(requestUserId uint64, hashedNewPasswordStringed string) error
}
