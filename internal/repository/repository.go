package repository

import (
	"SocialNetHTTPService/internal/domain"
	"context"
)

type UsersRep interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	CreateUsers(ctx context.Context, users []domain.User) error
	DeleteUser(ctx context.Context, user domain.User) error
	DeleteUserById(ctx context.Context, userID int) error
	DeleteUsers(ctx context.Context, users []domain.User) error
	ContainsUserById(ctx context.Context, userID int) bool
	UpdateAgeUser(ctx context.Context, userID int, age int) error
	AddFriendUser(ctx context.Context, userID int, friendID int) error
	GetUser(ctx context.Context, userID int) (domain.User, error)
	DeleteFriend(ctx context.Context, sourceUserID int, targetFriendID int) (int, error)
}

type Repositories struct {
	Users UsersRep
}

func NewRepositories(users UsersRep) *Repositories {
	return &Repositories{
		Users: users}
}
