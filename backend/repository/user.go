package repository

import (
	"chatApp/constant"
	"chatApp/entity"
	"context"
	"database/sql"
)

type UserRepoItf interface {
	RepoGetUserProfile(context.Context,entity.ReqUserProfileBody) (*entity.ResUserProfile, error)
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoImpl {
	return UserRepoImpl{db: db}
}

func (ur UserRepoImpl) RepoGetUserProfile(c context.Context, req entity.ReqUserProfileBody) (*entity.ResUserProfile, error){

	row := ur.db.QueryRowContext(c, `
		SELECT email, username, tag, img
		FROM users
		WHERE email = $1;
	`, req.Email)

	var user entity.ResUserProfile
	err := row.Scan(&user.Email, &user.Username, &user.Tag, &user.Img)

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
	}
	
	return &user, nil
}