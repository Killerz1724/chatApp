package handler

import (
	"chatApp/dto"
	"chatApp/entity"
	"chatApp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	au usecase.AuthUsecaseItf
}

func NewAuthHandler(au usecase.AuthUsecaseItf) AuthHandlerImpl {
	return AuthHandlerImpl{au: au}
}

func (ah AuthHandlerImpl) HandlerLogin(c *gin.Context) {
	var reqBody dto.LoginBody
	err := c.ShouldBindBodyWithJSON(&reqBody)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	req := entity.LoginBody(reqBody)

	res, err := ah.au.UsecaseLoginRepo(c, req)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resBody := dto.LoginRes(*res)

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Error:   nil,
		Data:    resBody,
	})
}