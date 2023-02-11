package service

import (
	"SocialNetHTTPService/internal/repository"
	"context"
)

type FriendsService struct {
	repo repository.UsersRep
}

func (f *FriendsService) MakeFriend(ctx context.Context, sourceUserID int, targetUserID int) error {
	err := error(nil)
	if err = f.repo.AddFriendUser(ctx, sourceUserID, targetUserID); err != nil {
		return err
	}
	if err = f.repo.AddFriendUser(ctx, targetUserID, sourceUserID); err != nil {
		return err
	}
	return nil
}

func (f *FriendsService) GetFriendsUser(ctx context.Context, userId int) ([]int, error) {
	user, err := f.repo.GetUser(ctx, userId)
	friendsUser := user.Friends()
	return friendsUser, err
}

func NewFriendsService(repo repository.UsersRep) *FriendsService {
	return &FriendsService{repo: repo}
}
