package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"qiuyier/blog/pkg/config"
	"regexp"
	"unicode"
)

func AbsUrl(path string) string {
	return config.Instance.BaseUrl + ":" + config.Instance.Port + path
}

func GeneratePassword(password string) string {
	passwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwd)
}

func ValidatePassword(dbPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	return err == nil
}

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsEmail(email string) (err error) {
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

// RuneLen 字符长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

func IsPassword(password, rePassword string) error {
	if IsBlank(password) {
		return errors.New("请输入密码")
	}
	if RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}

// IsUsername 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
func IsUsername(username string) error {
	if IsBlank(username) {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

// Struct2map 将struct转换成map
func Struct2map(content interface{}) (map[string]interface{}, error) {
	jsonStr, _ := json.Marshal(content)
	result := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
