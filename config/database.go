package config

import (
	"os"
	"sync"
)

type Config struct {
	Username         string
	Password         string
	Host             string
	Port             string
	Database         string
	DynamoDbEndpoint string
	S3Endpoint       string
}

var instance *Config
var once sync.Once

func InitDbConf(env string) {
	once.Do(func() {
		testFlag := os.Getenv("TEST_FLAG")
		if testFlag != "" {
			instance = &Config{
				Username:         "root",
				Password:         "root",
				Host:             "127.0.0.1",
				Port:             "3306",
				Database:         "book_recorder_test",
				DynamoDbEndpoint: "http://localhost:3307",
				S3Endpoint:       "http://localhost:9000",
			}
		} else {
			instance = &Config{
				Username:         os.Getenv("USERNAME"),
				Password:         os.Getenv("PASSWORD"),
				Host:             os.Getenv("HOST"),
				Port:             os.Getenv("PORT"),
				Database:         os.Getenv("DBNAME"),
				DynamoDbEndpoint: os.Getenv("DYNAMODB_ENDPOINT"),
				S3Endpoint:       os.Getenv("S3_ENDPOINT"),
			}
		}
	})
}

func GetDbConf() *Config {
	return instance
}
