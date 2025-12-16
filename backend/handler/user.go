package handler

import (
	"chatApp/dto"
	"chatApp/entity"
	"chatApp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	uu usecase.UserUsecaseItf
}

func NewUserHandler(uu usecase.UserUsecaseItf) UserHandlerImpl {
	return UserHandlerImpl{uu: uu}
}

func (uh UserHandlerImpl) HandlerGetUserProfile(c *gin.Context) {
	sub := c.GetString("subject")

	reqBody := entity.ReqUserProfileBody{Email: sub}

	res, err := uh.uu.UsecaseGetUserProfile(c, reqBody)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resBody := dto.ResUserProfile(*res)

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Error:   nil,
		Data:    resBody,
	})
}