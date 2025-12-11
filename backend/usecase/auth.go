package usecase

import (
	"chatApp/constant"
	"chatApp/entity"
	"chatApp/repository"
	"chatApp/utils"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseItf interface {
	UsecaseLoginRepo(context.Context, entity.LoginBody) (*entity.LoginRes, error)
}

type AuthUsecaseImpl struct {
	ar repository.AuthRepositoryItf
}

func NewAuthUsecase(ar repository.AuthRepositoryItf) AuthUsecaseImpl{
	return AuthUsecaseImpl{ar: ar}
}

func (au AuthUsecaseImpl)UsecaseLoginRepo(c context.Context, req entity.LoginBody) (*entity.LoginRes, error){

	pass, err := au.ar.RepoAuthLogin(c, req)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(req.Password))

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.LoginErrorType{Msg: constant.ErrLoginNotFound.Error()}, Log: err}
	}


	token, err := utils.GeneratingJWTToken(req.Email)

	if err != nil {
		return nil, err
	}

	return  &entity.LoginRes{
		Token: token,
	}, nil
}
