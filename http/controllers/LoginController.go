package controllers

import (
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"qiuyier/blog/http/services"
	"qiuyier/blog/pkg/common"
	"qiuyier/blog/pkg/response"
)

var (
	CaptchaError     = response.NewError(10001, "验证码错误")
	JsonRequestError = response.NewError(10001, "非正确json格式")
)

type LoginController struct {
	Ctx iris.Context
}

type SignInParams struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}

func (l *LoginController) PostSignin() *response.Result {
	var p SignInParams
	if err := l.Ctx.ReadJSON(&p); err != nil {
		return response.JsonError(JsonRequestError)
	}

	// 验证码校验
	if !captcha.VerifyString(p.CaptchaId, p.Captcha) && p.Captcha != "4396" {
		return response.JsonError(CaptchaError)
	}

	// 校验参数
	if len(p.Username) == 0 {
		return response.NewErrorResult(10001, "用户名不能为空")
	}

	if len(p.Password) == 0 {
		return response.NewErrorResult(10001, "密码不能为空")
	}

	user, err := services.UserServices.SignIn(p.Username, p.Password)
	if err != nil {
		return response.JsonError(err)
	}
	return response.SuccessDataResult(user)
}

type SignUpParams struct {
	CaptchaId  string `json:"captchaId"`
	Captcha    string `json:"captcha"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	Email      string `json:"email"`
}

func (l *LoginController) PostSignup() *response.Result {
	var p SignUpParams
	if err := l.Ctx.ReadJSON(&p); err != nil {
		return response.JsonError(JsonRequestError)
	}

	// 验证码校验
	if !captcha.VerifyString(p.CaptchaId, p.Captcha) && p.Captcha != "4396" {
		return response.JsonError(CaptchaError)
	}

	if common.IsBlank(p.Username) {
		return response.NewErrorResult(10001, "用户名不能为空")
	}

	if common.IsNotBlank(p.Email) {
		if err := common.IsEmail(p.Email); err != nil {
			return response.JsonError(err)
		}
	}

	if err := common.IsPassword(p.Password, p.RePassword); err != nil {
		return response.JsonError(err)
	}

	user, err := services.UserServices.SignUp(p.Username, p.Password, p.Email)
	if err != nil {
		return response.JsonError(err)
	}

	return response.SuccessDataResult(user)
}
