package config

import (
	"os"
	"sync"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var instance *Config
var once sync.Once

func InitDbConf(env string) {
	once.Do(func() {
		if env == "test" {
			instance = &Config{
				Username: "root",
				Password: "root",
				Host:     "127.0.0.1",
				Port:     "3306",
				Database: "book_recorder_test",
			}
		} else {
			instance = &Config{
				Username: os.Getenv("USERNAME"),
				Password: os.Getenv("PASSWORD"),
				Host:     os.Getenv("HOST"),
				Port:     os.Getenv("PORT"),
				Database: os.Getenv("DBNAME"),
			}
		}
	})
}

func GetDbConf() *Config {
	return instance
}
