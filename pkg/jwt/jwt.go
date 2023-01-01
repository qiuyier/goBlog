package jwt

import (
	"github.com/golang-module/carbon/v2"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"qiuyier/blog/model"
	"qiuyier/blog/pkg/cache"
	"qiuyier/blog/pkg/config"
	"qiuyier/blog/pkg/constants"
	"qiuyier/blog/pkg/db"
	"strconv"
)

func GenerateToken(user *model.User) (string, error) {
	jwtExp, _ := cache.Cache.HGetValue(db.Rdb(), "sys_config", "tokenExpireDays")
	jwtExpInt, err := strconv.Atoi(jwtExp)
	if err != nil {
		jwtExpInt = constants.DefaultJwtExp
	}
	if jwtExpInt <= 0 {
		jwtExpInt = constants.DefaultJwtExp
	}

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.Id,
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"iss":      "goBlog",
		"iat":      carbon.Now().Timestamp(),
		"exp":      carbon.Now().AddDays(jwtExpInt).Timestamp(),
	})

	tokenString, err := token.SignedString([]byte(config.Instance.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetInfoFromJwt(ctx iris.Context) *jwt.Token {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token)
	return jwtInfo
}
