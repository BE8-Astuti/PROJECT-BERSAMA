package config

import (
	"fmt"
	"together/be8/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := InitConfig()
	fmt.Printf("%v", config)
	conString := fmt.Sprintf("%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Host,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
