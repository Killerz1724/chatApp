package repository

import (
	"chatApp/constant"
	"chatApp/entity"
	"context"
	"database/sql"
)

type ConvRepositoryItf interface{
	RepoGetListConvs(c context.Context, req entity.ReqListConvBody) (*[]entity.ResListConvBody, error)
}

type ConvRepositoryImpl struct {
	db *sql.DB
}

func NewConvRepository(db *sql.DB) ConvRepositoryImpl {
	return ConvRepositoryImpl{db: db}
}

func (cr ConvRepositoryImpl) RepoGetListConvs(c context.Context, req entity.ReqListConvBody) (*[]entity.ResListConvBody, error) {

	var userId int8

	row := cr.db.QueryRowContext(c, `
		SELECT id
		FROM users
		WHERE email = $1;
	`, req.Email)

	err := row.Scan(&userId)

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
	}
	rows, err := cr.db.QueryContext(c, `
	WITH last_message AS (
  	SELECT DISTINCT ON (conversation_id)
    conversation_id,
    content,
    created_at
  	FROM messages
  	WHERE deleted_at IS NULL
  	ORDER BY conversation_id, created_at DESC
	)
	SELECT DISTINCT ON (conversation_id)
  c.id AS conversation_id,
  u.id AS user_id,
	c.is_group,
  u.username,
  u.tag,
  u.img,
  lm.content AS last_message,
  lm.created_at AS last_message_at
	FROM conversations c
	JOIN conversationsMembers cm_self
  ON cm_self.conversation_id = c.id
	JOIN conversationsMembers cm_other
  ON cm_other.conversation_id = c.id
	JOIN users u
  ON u.id = cm_other.user_id
	LEFT JOIN last_message lm
  ON lm.conversation_id = c.id
	WHERE cm_self.user_id = $1
  AND cm_other.user_id <> $1
  AND cm_self.deleted_at IS NULL
  AND cm_other.deleted_at IS NULL
  AND c.deleted_at IS NULL
	ORDER BY c.id, lm.created_at DESC NULLS LAST;

	`, userId)

	if err != nil {
		return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
	}

	var res []entity.ResListConvBody

	for rows.Next() {
		var conv entity.ResListConvBody
		err = rows.Scan(&conv.ConversationId, &conv.User.Id,&conv.IsGroup, &conv.User.Username, &conv.User.Tag, &conv.User.Img, &conv.Message.Message, &conv.Message.Time)
		if err != nil {
			return nil, &entity.CustomError{Msg: constant.ErrCommon, Log: err}
		}
		res = append(res, conv)
	}

	return &res, nil
}