package dao

import (
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
	"qiuyier/blog/model"
)

type userFollowDao struct {
}

func newUserFollowDao() *userFollowDao {
	return &userFollowDao{}
}

var UserFollowDao = newUserFollowDao()

func (u *userFollowDao) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.UserFollow) {
	cnd.Find(db, &list)
	return
}
