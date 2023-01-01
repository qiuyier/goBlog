package services

import (
	"errors"
	"github.com/mlogclub/simple/sqls"
	"qiuyier/blog/http/dao"
	"qiuyier/blog/model"
	"qiuyier/blog/pkg/cache"
	"qiuyier/blog/pkg/common"
	"qiuyier/blog/pkg/constants"
	"qiuyier/blog/pkg/db"
	jwt2 "qiuyier/blog/pkg/jwt"
	"qiuyier/blog/pkg/response"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type userServices struct {
}

func newUserService() *userServices {
	return &userServices{}
}

var UserServices = newUserService()

func (u *userServices) SignIn(username, password string) (map[string]interface{}, error) {
	var userModel *model.User = u.GetByUsername(username)

	if userModel == nil {
		return nil, errors.New("用户不存在")
	} else if userModel.Status == constants.StatusForbidden || userModel.IsDeleted == constants.StatusDeleted {
		return nil, errors.New("用户被禁用")
	}

	if !common.ValidatePassword(userModel.Password, password) {
		return nil, errors.New("密码不正确")
	}

	userMap, _ := common.Struct2map(userModel)
	userId := strconv.Itoa(int(userModel.Id))
	if err := cache.Cache.HSet(db.Rdb(), "user_"+userId, userMap, constants.DefaultCacheTime); err != nil {
		return nil, err
	}

	tokenString, err := jwt2.GenerateToken(userModel)
	if err != nil {
		return nil, err
	}
	result := response.NewEmptyRspBuilder().
		Put("userId", userModel.Id).
		Put("nickname", userModel.Nickname).
		Put("avatar", userModel.Avatar).
		Put("token", tokenString).
		GetStruct()
	return result, nil
}

func (u *userServices) GetByUsername(username string) *model.User {
	return dao.UserDao.GetByUsername(db.DB(), username)
}

func (u *userServices) SignUp(username, password, email string) (map[string]interface{}, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	if u.GetByUsername(username) != nil {
		return nil, errors.New("用户名已被占用")
	}

	if common.IsNotBlank(email) {
		if u.GetByEmail(email) != nil {
			return nil, errors.New("邮箱已被占用")
		}
	}

	user := &model.User{
		Username: username,
		Password: common.GeneratePassword(password),
		Email:    email,
	}

	if err := dao.UserDao.Create(db.DB(), user); err != nil {
		return nil, err
	}

	userMap, _ := common.Struct2map(user)
	userId := strconv.Itoa(int(user.Id))
	if err := cache.Cache.HSet(db.Rdb(), "user_"+userId, userMap, constants.DefaultCacheTime); err != nil {
		return nil, err
	}

	tokenString, err := jwt2.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	result := response.NewEmptyRspBuilder().
		Put("userId", user.Id).
		Put("nickname", user.Nickname).
		Put("avatar", user.Avatar).
		Put("token", tokenString).
		GetStruct()
	return result, nil
}

func (u *userServices) GetByEmail(email string) *model.User {
	return dao.UserDao.GetByEmail(db.DB(), email)
}

func (u *userServices) GetUserProfile(userId float64) (map[string]string, error) {
	result := make(map[string]string)
	userIdStr := strconv.Itoa(int(userId))
	key := "user_" + userIdStr
	userInfo, err := cache.Cache.HGetAll(db.Rdb(), key)

	if err != nil {
		return nil, err
	}
	if len(userInfo) == 0 {
		user := dao.UserDao.GetById(db.DB(), int64(userId))
		if user == nil {
			return nil, errors.New("用户不存在")
		}
		userMap, err := common.Struct2map(user)
		if err != nil {
			return nil, err
		}
		if err := cache.Cache.HSet(db.Rdb(), "user_"+userIdStr, userMap, constants.DefaultCacheTime); err != nil {
			return nil, err
		}

		result["userId"] = strconv.FormatInt(user.Id, 10)
		result["nickname"] = user.Nickname
		result["description"] = user.Description
		result["avatar"] = user.Avatar
		result["email"] = user.Email
		result["homePage"] = user.HomePage
		return result, nil
	}

	result["userId"] = userInfo["id"]
	result["nickname"] = userInfo["nickname"]
	result["description"] = userInfo["description"]
	result["avatar"] = userInfo["avatar"]
	result["email"] = userInfo["email"]
	result["homePage"] = userInfo["homePage"]

	return result, nil
}

func (u *userServices) GetFansList(userId float64, cursor, limit int) (fanList []map[string]interface{}, nextCursor int64, hasMore bool, err error) {
	// 获取关注自己的粉丝
	cnd := sqls.NewCnd().Eq("follow_id")
	if cursor > 0 {
		cnd.Lt("id", cursor)
	}
	cnd.Desc("id").Limit(limit)
	userList := dao.UserFollowDao.Find(db.DB(), cnd)
	var followIds []int64
	for _, v := range userList {
		followIds = append(followIds, v.UserId)
	}

	// 获取自己关注的人
	var myFollowIds hashset.Set
	myFollowIds = u.MyFollowUserIds(int64(userId), followIds...)

	if len(userList) > 0 {
		nextCursor = userList[len(userList)-1].Id
		hasMore = len(userList) >= limit
		for k, v := range userList {
			userInfo, _ := u.GetUserProfile(float64(v.UserId))
			fanList[k]["userId"] = userInfo["userId"]
			fanList[k]["nickname"] = userInfo["nickname"]
			fanList[k]["description"] = userInfo["description"]
			fanList[k]["avatar"] = userInfo["avatar"]
			fanList[k]["follow"] = myFollowIds.Contains(v.UserId)
		}
	} else {
		nextCursor = int64(cursor)
	}
	return
}

func (u *userServices) MyFollowUserIds(userId int64, followIds ...int64) hashset.Set {
	set := hashset.New()

	cnd := sqls.NewCnd().Eq("user_id", userId).In("follow_id", followIds)
	followUserList := dao.UserFollowDao.Find(db.DB(), cnd)

	for _, follow := range followUserList {
		set.Add(follow.FollowId)
	}
	return *set
}
