package repository

import (
	"chatApp/constant"
	"chatApp/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type AuthRepoImpl struct {
	db *sql.DB
}

type AuthRepositoryItf interface {
	RepoAuthLogin(context.Context, entity.LoginBody) (string, error)
	RepoAuthRegister(context.Context, entity.RegisterBody) error
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

func (ar AuthRepoImpl) RepoAuthRegister(c context.Context, req entity.RegisterBody) error {
	_,err := ar.db.ExecContext(c, `
	INSERT INTO users(email, password, username, tag)
	VALUES ($1, $2, $3, $4);
	`, req.Email, req.Password, req.Username, req.Tag)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return &entity.CustomError{Msg: err, Log: err}
	}

	if err != nil {
		return &entity.CustomError{Msg: constant.LoginErrorType{Msg: constant.ErrCommon.Error()}, Log: err}
	}

	return nil
}
