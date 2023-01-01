package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"qiuyier/blog/pkg/config"
)

var (
	db    *gorm.DB
	sqlDb *sql.DB
)

func Init(dbConfig config.DB, config *gorm.Config, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   dbConfig.NamingStrategy,
			SingularTable: true,
		}
	}

	if db, err = gorm.Open(mysql.Open(dbConfig.Dsn), config); err != nil {
		logrus.Errorf("database connect failed: %s", err.Error())
		return
	}

	if sqlDb, err = db.DB(); err == nil {
		sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)
	} else {
		logrus.Error(err)
	}

	if err = db.AutoMigrate(models...); err != nil {
		logrus.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

func DB() *gorm.DB {
	return db
}
