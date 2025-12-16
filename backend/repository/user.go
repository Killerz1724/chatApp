package repository

import (
	"chatApp/constant"
	"chatApp/entity"
	"context"
	"database/sql"
)

type UserRepoItf interface {
	RepoGetUserProfile(context.Context,entity.ReqUserProfileBody) (*entity.ResUserProfile, error)
	RepoGetUserFriends(context.Context, entity.ReqUserProfileBody) (*[]entity.ResUserFriend, error)
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

func (ur UserRepoImpl) RepoGetUserFriends(c context.Context, req entity.ReqUserProfileBody) (*[]entity.ResUserFriend, error) {

	rows, err := ur.db.QueryContext(c, `
		SELECT  u.username, u.tag, u.img
		FROM users u
		JOIN friends f
		ON f.user_id = u.id
		WHERE u.email = $1;
	`, req.Email)

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
	}

	var users []entity.ResUserFriend
	for rows.Next(){
		var user entity.ResUserFriend
		err := rows.Scan(&user.Username, &user.Tag, &user.Img)
		if err != nil {
			return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
		}
		users = append(users, user)
	}

	return &users, nil
}