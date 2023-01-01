package controllers

import (
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"qiuyier/blog/pkg/common"
	"qiuyier/blog/pkg/response"
)

type CaptchaController struct {
	Ctx iris.Context
}

func (c *CaptchaController) GetRequest() *response.Result {
	captchaId := captcha.NewLen(4)
	captchaUrl := common.AbsUrl("/api/captcha/show?captchaId=" + captchaId)
	return response.NewEmptyRspBuilder().
		Put("captchaId", captchaId).
		Put("captchaUrl", captchaUrl).
		BuildResult()
}

func (c *CaptchaController) GetShow() *response.Result {
	captchaId := c.Ctx.URLParam("captchaId")

	if captchaId == "" {
		c.Ctx.StatusCode(400)
		return response.NewErrorResult(10001, "缺少验证码ID")
	}

	if !captcha.Reload(captchaId) {
		c.Ctx.StatusCode(400)
		return response.NewErrorResult(10001, "验证码ID错误")
	}

	c.Ctx.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(c.Ctx.ResponseWriter(), captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
		logrus.Error(err)
	}
	return nil
}
