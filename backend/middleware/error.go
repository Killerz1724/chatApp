package middleware

import (
	"chatApp/constant"
	"chatApp/dto"
	"chatApp/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]

	

		var pgErr *pgconn.PgError
		var cerr *entity.CustomError

		if errors.As(err, &cerr) {
			c.Error(cerr.Log).SetType(gin.ErrorTypePrivate)

			switch {
			case errors.As(cerr.Msg, &pgErr):
				msg := ""
				if pgErr.Code == "23505" {
					msg = constant.ErrUserAlreadyExist.Error()
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
					Success: false,
					Error: &dto.ErrorResponse{
						Message: msg,
					},
					Data: nil,
				})
				return
			case errors.As(cerr.Msg, &constant.LoginErrorType{}):
				c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
					Success: false,
					Error: &dto.ErrorResponse{
						Message: cerr.Error(),
					},
					Data: nil,
				})
				return
			}

			
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: &dto.ErrorResponse{
				Message: constant.ErrCommon.Error(),
			},
			Data: nil,
		})
	}

}