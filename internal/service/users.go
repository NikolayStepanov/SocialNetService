package service

import (
	"SocialNetHTTPService/internal/domain"
	"SocialNetHTTPService/internal/repository"
	"context"
	"log"
)

type UsersService struct {
	repo repository.UsersRep
}

func (us *UsersService) GetNameUser(ctx context.Context, userId int) (string, error) {
	userName := ""
	user := domain.User{}
	err := error(nil)
	user, err = us.repo.GetUser(ctx, userId)
	userName = user.Name()
	return userName, err
}

func (us *UsersService) GetAgeUser(ctx context.Context, userId int) (int, error) {
	ageUser := 0
	user := domain.User{}
	err := error(nil)
	user, err = us.repo.GetUser(ctx, userId)
	ageUser = int(user.Age())
	return ageUser, err
}

func (us *UsersService) CreateUser(ctx context.Context, name string, age int, friends []int) (int, error) {
	user := domain.NewUser(0, name, uint8(age), friends)
	userId, err := us.repo.CreateUser(ctx, user)
	return userId, err
}

func (us *UsersService) DeleteUser(ctx context.Context, userID int) (string, error) {
	var (
		err  error
		user domain.User
	)
	if user, err = us.repo.GetUser(ctx, userID); err != nil {
		log.Println(err)
	} else {
		friendsUser := user.Friends()
		for _, friendId := range friendsUser {
			us.repo.DeleteFriend(ctx, friendId, userID)
		}
		if err = us.repo.DeleteUserById(ctx, userID); err != nil {
			log.Println(err)
		}
	}
	return user.Name(), err
}

func (us *UsersService) UpdateUserAge(ctx context.Context, userID int, age int) error {
	err := us.repo.UpdateAgeUser(ctx, userID, age)
	return err
}

func NewUserService(repo repository.UsersRep) *UsersService {
	return &UsersService{repo: repo}
}
