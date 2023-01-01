package routes

import (
	"github.com/go-resty/resty/v2"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"os"
	"qiuyier/blog/http/controllers"
	"qiuyier/blog/pkg/config"
	"qiuyier/blog/pkg/middlewares"
	"qiuyier/blog/pkg/response"
	"strings"
)

func Routes() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(irisRecover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
		MaxAge:           600,
	}))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "/api/admin") {
			err = ctx.JSON(response.NewErrorResult(ctx.GetStatusCode(), "HTTP ERROR"))
		}
		if err != nil {
			logrus.Error(err)
		}
	})

	app.Any("/", func(i iris.Context) {
		_ = i.JSON(map[string]string{
			"method":  i.Method(),
			"message": "Hello, this is goBlog project",
		})
	})

	// api
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/captcha").Handle(new(controllers.CaptchaController))
		m.Party("/login").Handle(new(controllers.LoginController))
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middlewares.JwtHandler)
		m.Party("/user").Handle(new(controllers.UserCenterController))
	})

	app.Get("/api/img/proxy", func(i iris.Context) {
		url := i.FormValue("url")
		resp, err := resty.New().R().Get(url)
		i.Header("Content-Type", "image/jpg")
		if err == nil {
			_, _ = i.Write(resp.Body())
		} else {
			logrus.Error(err)
		}
	})

	err := app.Listen(":"+config.Instance.Port, iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
