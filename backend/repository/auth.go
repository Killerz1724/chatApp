package repository

import (
	"chatApp/constant"
	"chatApp/entity"
	"context"
	"database/sql"
)

type AuthRepoImpl struct {
	db *sql.DB
}

type AuthRepositoryItf interface {
	RepoAuthLogin(context.Context, entity.LoginBody) (string, error)
}

func NewAuthRepo(db *sql.DB) AuthRepoImpl {
	return AuthRepoImpl{db: db}
}

func (ar AuthRepoImpl) RepoAuthLogin(c context.Context, req entity.LoginBody) (string, error) {

	row := ar.db.QueryRowContext(c, `
	SELECT password
	FROM users
	WHERE email = $1;
	`, req.Email)

	var userPassword string
	err := row.Scan(&userPassword)

	if err != nil {
		return "", &entity.CustomError{Msg: constant.LoginErrorType{Msg: constant.ErrLoginNotFound.Error()}, Log: err}
	}

	return userPassword, nil
}
