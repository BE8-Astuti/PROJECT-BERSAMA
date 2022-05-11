package config

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
<<<<<<< HEAD
	Port     int16
	DBPort   int16
	Host     string
	User     string
	Password string
	DBName   string
=======
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
>>>>>>> 03362f06d487b54d41aeb62a1a3a89dd3f5a3e8b
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func InitConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "layered_db")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "")

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}
<<<<<<< HEAD
	portconv, _ := strconv.Atoi(os.Getenv("PORT"))
	res.Port = int16(portconv)
	conv, _ := strconv.Atoi(os.Getenv("DBPORT"))
	res.DBPort = int16(conv)
	res.Host = os.Getenv("HOST")
	res.User = os.Getenv("NAMEUSER")
	res.Password = os.Getenv("PASSWORD")
	res.DBName = os.Getenv("DBNAME")
	return &res
=======

	return fallback

>>>>>>> 03362f06d487b54d41aeb62a1a3a89dd3f5a3e8b
}
