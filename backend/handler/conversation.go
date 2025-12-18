package handler

import (
	"chatApp/dto"
	"chatApp/entity"
	"chatApp/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerConvImpl struct {
	cu usecase.ConvUsecaseItf
}

func NewConvHandler(cu usecase.ConvUsecaseItf) HandlerConvImpl{
	return HandlerConvImpl{cu: cu}
}

func (ch HandlerConvImpl) HandlerGetListConv(c *gin.Context) {

	sub := c.GetString("subject")

	reqBody := entity.ReqListConvBody{Email: sub}

	res, err := ch.cu.UsecaseGetListConvs(c, reqBody)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	var resBody []dto.ResListConvBody

	
	for _, conv := range *res {
		resBody = append(resBody, dto.ResListConvBody{
			ConversationId: conv.ConversationId,
			User:           dto.User(conv.User),
			Message:        dto.LastMessage(conv.Message),
		})
	}
	

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Error:   nil,
		Data:    resBody,
	})
}