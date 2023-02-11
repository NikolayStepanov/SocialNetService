package service

import (
	"SocialNetHTTPService/internal/repository"
	"context"
)

type ServicesDeps struct {
	Repos *repository.Repositories
}

type Users interface {
	CreateUser(ctx context.Context, name string, age int, friends []int) (int, error)
	DeleteUser(ctx context.Context, userID int) (string, error)
	GetNameUser(ctx context.Context, userID int) (string, error)
	GetAgeUser(ctx context.Context, userID int) (int, error)
	UpdateUserAge(ctx context.Context, userID int, age int) error
}
type Friends interface {
	MakeFriend(ctx context.Context, sourceUserID int, targetUserID int) error
	GetFriendsUser(ctx context.Context, userID int) ([]int, error)
}
type Services struct {
	Users   Users
	Friends Friends
}

func NewServices(deps ServicesDeps) *Services {
	usersService := NewUserService(deps.Repos.Users)
	friendsService := NewFriendsService(deps.Repos.Users)
	return &Services{Users: usersService, Friends: friendsService}
}
