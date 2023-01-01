package controllers

import (
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"qiuyier/blog/http/services"
	"qiuyier/blog/pkg/jwt"
	"qiuyier/blog/pkg/response"
	"strconv"
)

type UserCenterController struct {
	Ctx iris.Context
}

func (u *UserCenterController) GetProfile() *response.Result {
	jwtInfo := jwt.GetInfoFromJwt(u.Ctx)
	userId := jwtInfo.Claims.(jwt2.MapClaims)["userId"].(float64)
	user, err := services.UserServices.GetUserProfile(userId)
	if err != nil {
		return response.JsonError(err)
	}
	return response.SuccessDataResult(user)
}

type FansParam struct {
	Cursor int `json:"cursor"`
}

func (u *UserCenterController) GetFans() *response.Result {
	jwtInfo := jwt.GetInfoFromJwt(u.Ctx)
	userId := jwtInfo.Claims.(jwt2.MapClaims)["userId"].(float64)
	var p FansParam
	if err := u.Ctx.ReadJSON(&p); err != nil {
		return response.JsonError(JsonRequestError)
	}
	fansList, nextCursor, hasMore, err := services.UserServices.GetFansList(userId, p.Cursor, 10)
	if err != nil {
		return response.JsonError(err)
	}
	return response.CursorDataResult(fansList, strconv.FormatInt(nextCursor, 10), hasMore)
}
