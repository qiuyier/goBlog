package model

import (
	"qiuyier/blog/pkg/constants"
	"time"
)

var Models = []interface{}{
	&User{}, &UserToken{},
}

// DefaultModel 每张表固定有的字段
type DefaultModel struct {
	Id         int64 `gorm:"primaryKey;autoIncrement;type:int(1)" json:"id" form:"id"`
	IsDeleted  int   `gorm:"type:tinyint(1);default:0;comment:是否删除,1-已删除 0-未删除;" json:"IsDeleted" form:"isDeleted"`
	CreateTime int64 `gorm:"not null;type:int(1);autoCreateTime;comment:创建时间;" json:"createTime" form:"createTime"`
	UpdateTime int64 `gorm:"not null;type:int(1);autoUpdateTime;comment:更新时间;" json:"updateTime" form:"updateTime"`
}

// User 用户表
type User struct {
	DefaultModel
	Username         string           `gorm:"size:15;unique;not null;comment:用户名;default:'';" json:"username" form:"username"`
	Email            string           `gorm:"size:128;not null;unique;comment:邮箱;default:'';" json:"email" form:"email"`
	EmailVerified    bool             `gorm:"not null;default:false;comment:邮箱是否验证;" json:"emailVerified" form:"emailVerified"`
	Nickname         string           `gorm:"size:16;comment:昵称;default:'';" json:"nickname" form:"nickname"`
	Avatar           string           `gorm:"size:250;comment:头像;default:'';" json:"avatar" form:"avatar"`
	Gender           constants.Gender `gorm:"type:tinyint;size:1;default:0;;comment:性别,1-男 2-女 0-未知;" json:"gender" form:"gender"`
	Birthday         *time.Time       `gorm:"comment:生日;" json:"birthday" form:"birthday"`
	BackgroundImage  string           `gorm:"size:250;comment:个人中心背景图片;default:'';" json:"backgroundImage" form:"backgroundImage"`
	Password         string           `gorm:"size:512;comment:密码;not null;" json:"password" form:"password"`
	HomePage         string           `gorm:"size:1024;comment:个人主页;default:'';" json:"homePage" form:"homePage"`
	Description      string           `gorm:"size:250;comment:个人描述;default:'';" json:"description" form:"description"`
	Score            int              `gorm:"not null;index:idx_user_score;comment:积分;default:0" json:"score" form:"score"`
	Status           int              `gorm:"index:idx_user_status;not null;comment:状态,1-正常,-1-禁用;default:1" json:"status" form:"status"`
	TopicCount       int              `gorm:"not null;comment:帖子数量;default:0" json:"topicCount" form:"topicCount"`
	CommentCount     int              `gorm:"not null;comment:跟帖数量;default:0" json:"commentCount" form:"commentCount"`
	FollowCount      int              `gorm:"not null;comment:关注数量;default:0" json:"followCount" form:"followCount"`
	FansCount        int              `gorm:"not null;comment:粉丝数量;default:0" json:"fansCount" form:"fansCount"`
	Roles            string           `gorm:"size:250;comment:角色;default:'';" json:"roles" form:"roles"`
	ForbiddenEndTime int64            `gorm:"not null;default:0;comment:禁言结束时间;" json:"forbiddenEndTime" form:"forbiddenEndTime"`
}

// UserToken 用户Token表
type UserToken struct {
	DefaultModel
	Token     string `gorm:"size:32;unique;not null;comment:用户token;" json:"token" form:"token"`
	UserId    int64  `gorm:"not null;index:idx_user_token_user_id;comment:用户ID;" json:"userId" form:"userId"`
	Ip        string `gorm:"size:15;not null;comment:ip地址;" json:"ip" form:"ip"`
	ExpiredAt int64  `gorm:"not null;comment:失效时间;" json:"expiredAt" form:"expiredAt"`
}

// UserFollow 粉丝关注
type UserFollow struct {
	DefaultModel
	UserId   int64 `gorm:"not null;uniqueIndex:idx_user_id" json:"userId"`   // 用户编号
	FollowId int64 `gorm:"not null;uniqueIndex:idx_user_id" json:"followId"` // 对方的ID（被关注用户编号）
}
