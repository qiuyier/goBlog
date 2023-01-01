package dao

import (
	"gorm.io/gorm"
	"qiuyier/blog/model"
)

type userDao struct {
}

func newUserDao() *userDao {
	return &userDao{}
}

var UserDao = newUserDao()

func (u *userDao) Take(db *gorm.DB, where ...interface{}) *model.User {
	result := &model.User{}
	if err := db.Take(result, where...).Error; err != nil {
		return nil
	}
	return result
}

func (u *userDao) Create(db *gorm.DB, data *model.User) (err error) {
	err = db.Create(data).Error
	return
}

func (u *userDao) GetByUsername(db *gorm.DB, username string) *model.User {
	return u.Take(db, "username = ?", username)
}

func (u *userDao) GetByEmail(db *gorm.DB, email string) *model.User {
	return u.Take(db, "email = ?", email)
}

func (u *userDao) GetById(db *gorm.DB, id int64) *model.User {
	result := &model.User{}
	if err := db.First(result, "id = ?", id).Error; err != nil {
		return nil
	}
	return result
}
