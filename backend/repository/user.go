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
	

	var userId int8

	row := ur.db.QueryRowContext(c, `
		SELECT id
		FROM users
		WHERE email = $1;
	`, req.Email)

	err := row.Scan(&userId)

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
	}

	params := []any{userId}
	q :=`
		SELECT  u.username, u.tag, u.img
		FROM users u
		JOIN friends f
		ON (
       (f.user_id = $1 AND f.friend_id = u.id)
    	OR (f.friend_id = $1 AND f.user_id = u.id)
  	)
		WHERE f.deleted_at IS NULL;
	`

	if(req.SearchP != ""){
		q += `AND u.username ILIKE '%' || $2 || '%'`
		params = append(params, req.SearchP)
	}

	q += ";"

	rows, err := ur.db.QueryContext(c, q, params...)

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