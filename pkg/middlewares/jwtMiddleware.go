package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"qiuyier/blog/pkg/config"
	"qiuyier/blog/pkg/response"
)

func JwtHandler(ctx iris.Context) {
	j := jwt.New(jwt.Config{
		Extractor: jwt.FromAuthHeader,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Instance.JwtKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			err = ctx.JSON(response.FailMsgResult("token已失效"))
			if err != nil {
				return
			}
		},
		Expiration: true,
	})
	j.Serve(ctx)
	ctx.Next()
}
