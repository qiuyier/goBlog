package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"qiuyier/blog/model"
	"qiuyier/blog/pkg/config"
	"qiuyier/blog/pkg/db"
	"qiuyier/blog/routes"
	"time"
)

var configFile = flag.String("config", "./env.yaml", "default config file path")

func init() {
	flag.Parse()

	// 初始化配置
	conf := config.Init(*configFile)

	// gorm配置
	gormConf := &gorm.Config{}

	// 初始化日志
	if file, err := os.OpenFile(conf.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		if conf.PrintSql {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}

	// 连接数据库
	if err := db.Init(conf.DB, gormConf, model.Models...); err != nil {
		logrus.Error(err)
	}

	// 连接redis
	if err := db.InitRedis(conf.Redis); err != nil {
		logrus.Error(err)
	}
}

func main() {
	routes.Routes()
}
