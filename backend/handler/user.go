package handler

import (
	"chatApp/dto"
	"chatApp/entity"
	"chatApp/usecase"
	"net/http"
	"strconv"

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

func (uh UserHandlerImpl) UsecaseGetUserFriends(c *gin.Context){
	sub := c.GetString("subject")
	searchParam := c.Query("search")

	reqBody := entity.ReqUserProfileBody{Email: sub, SearchP: searchParam}

	res, err := uh.uu.UsecaseGetUserFriends(c, reqBody)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	
	var friendList []dto.ResUserFriendProf
	for _,friend := range *res{
		resBody := dto.ResUserFriendProf(friend)
		friendList = append(friendList, resBody)
	}
	resBody := dto.ResUserFriend{
		Friend_list: friendList,
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Error:   nil,
		Data:    resBody,
	})
}

func (uh UserHandlerImpl)HandlerGetUserDetailFriend(c *gin.Context){

	sub := c.GetString("subject")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	reqBody := entity.ReqUserFriendDetailBody{Id: id, CurrentUserEmail: sub}

	res, err := uh.uu.UsecaseGetUserDetailFriend(c, reqBody)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	if res == nil {
		c.JSON(http.StatusOK, dto.Response{
			Success: true,
			Error:   nil,
			Data:    nil,
		})
		return
	}
	resBody := dto.ResUserFriendProf(*res)

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Error:   nil,
		Data:    resBody,
	})
}