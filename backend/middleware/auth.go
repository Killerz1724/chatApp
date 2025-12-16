package middleware

import (
	"chatApp/constant"
	"chatApp/entity"
	"errors"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenInvalid.Error()}, Log: constant.ErrJWTTokenInvalid}
			c.Error(cusErr).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}
		tokenString := authHeader[7:]
	

		token, err := jwt.Parse(
			tokenString,
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenInvalid.Error()}, Log: constant.ErrJWTTokenInvalid}
					c.Error(cusErr).SetType(gin.ErrorTypePrivate)
					return nil, cusErr
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
			// parser options
			jwt.WithIssuedAt(),
			jwt.WithIssuer(os.Getenv("JWT_USER")),
		)

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) { // error handling
				cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenExpired.Error()}, Log: err}
				c.Error(cusErr).SetType(gin.ErrorTypePrivate)

				c.Abort()
				return
			}
			cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenInvalid.Error()}, Log: err}
			c.Error(cusErr).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenInvalid.Error()}, Log: err}
			c.Error(cusErr).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		sub, err := claims.GetSubject()

		if err != nil {
			cusErr := &entity.CustomError{Msg: constant.TokenErrorType{Msg: constant.ErrJWTTokenInvalid.Error()}, Log: err}
			c.Error(cusErr).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		c.Set("subject", sub)
		c.Next()

	}
}