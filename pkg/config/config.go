package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

var Instance *Config

type Config struct {
	PrintSql    bool     `yaml:"PrintSql"`
	DB          DB       `yaml:"DB"`
	Uploader    Uploader `yaml:"Uploader"`
	Redis       Redis    `yaml:"Redis"`
	Env         string   `yaml:"Env"`
	BaseUrl     string   `yaml:"BaseUrl"`
	Port        string   `yaml:"Port"`
	LogFilePath string   `yaml:"LogFilePath"`
	JwtKey      string   `yaml:"JwtKey"`
}

type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}

type DB struct {
	Dsn            string `yaml:"Dsn"`
	MaxIdleConns   int    `yaml:"MaxIdleConns"`
	MaxOpenConns   int    `yaml:"MaxOpenConns"`
	NamingStrategy string `yaml:"NamingStrategy"`
}

type Uploader struct {
	AliYunOss AliYunOss `yaml:"AliYunOss"`
	Local     Local     `yaml:"Local"`
	Channel   string    `yaml:"Channel"`
}

type AliYunOss struct {
	Host          string `yaml:"Host"`
	AccessId      string `yaml:"AccessId"`
	StyleSplitter string `yaml:"StyleSplitter"`
	StylePreview  string `yaml:"StylePreview"`
	StyleSmall    string `yaml:"StyleSmall"`
	Bucket        string `yaml:"Bucket"`
	Endpoint      string `yaml:"Endpoint"`
	AccessSecret  string `yaml:"AccessSecret"`
	StyleAvatar   string `yaml:"StyleAvatar"`
	StyleDetail   string `yaml:"StyleDetail"`
}

type Local struct {
	Host string `yaml:"Host"`
	Path string `yaml:"Path"`
}

func Init(filePath string) *Config {
	Instance = &Config{}

	if yamlFile, err := os.ReadFile(filePath); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}

	return Instance
}
