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
	UsecaseAuthRegister(context.Context, entity.RegisterBody) error
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

func (ar AuthUsecaseImpl) UsecaseAuthRegister(c context.Context, req entity.RegisterBody) error {

	hashedPass, err  := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return &entity.CustomError{Msg: constant.LoginErrorType{Msg: constant.ErrCommon.Error()}, Log: err}
	}

	req.Password = string(hashedPass)

	tag, err := utils.GenerateTag()

	if err != nil {
		return &entity.CustomError{Msg: constant.LoginErrorType{Msg: constant.ErrCommon.Error()}, Log: err}
	}

	req.Tag = "#" + tag


	err = ar.ar.RepoAuthRegister(c, req)

	

	if err != nil {
		return err
	}	
	
	
	return nil
}