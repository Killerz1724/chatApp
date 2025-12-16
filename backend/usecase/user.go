package usecase

import (
	"chatApp/entity"
	"chatApp/repository"
	"context"
)

type UserUsecaseItf interface {
	UsecaseGetUserProfile(c context.Context, req entity.ReqUserProfileBody) (*entity.ResUserProfile, error)
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