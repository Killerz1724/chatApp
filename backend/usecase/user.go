package usecase

import (
	"chatApp/entity"
	"chatApp/repository"
	"context"
)

type UserUsecaseItf interface {
	UsecaseGetUserProfile(context.Context, entity.ReqUserProfileBody) (*entity.ResUserProfile, error)
	UsecaseGetUserFriends(context.Context, entity.ReqUserProfileBody) (*[]entity.ResUserFriend, error)
	UsecaseGetUserDetailFriend(context.Context, entity.ReqUserFriendDetailBody) (*entity.ResUserFriend, error)
}

type UserUsecaseImpl struct {
	ur repository.UserRepoItf
}

func NewUserUsecase(ur repository.UserRepoItf) UserUsecaseImpl {
	return UserUsecaseImpl{ur: ur}
}

func (uu UserUsecaseImpl) UsecaseGetUserProfile(c context.Context, req entity.ReqUserProfileBody) (*entity.ResUserProfile, error) {

	res, err := uu.ur.RepoGetUserProfile(c, req)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu UserUsecaseImpl) UsecaseGetUserFriends(c context.Context, req entity.ReqUserProfileBody) (*[]entity.ResUserFriend, error){

	res, err := uu.ur.RepoGetUserFriends(c, req)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu UserUsecaseImpl)UsecaseGetUserDetailFriend(c context.Context, req entity.ReqUserFriendDetailBody) (*entity.ResUserFriend, error){

	res, err := uu.ur.RepoGetUserDetailFriend(c, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}