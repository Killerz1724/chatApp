package middleware

import (
	"chatApp/constant"
	"chatApp/dto"
	"chatApp/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]

		var cerr *entity.CustomError

		if errors.As(err, &cerr) {
			c.Error(cerr.Log).SetType(gin.ErrorTypePrivate)

			switch {
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