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
<<<<<<< HEAD
	conString := fmt.Sprintf("%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		// config.Password,
		config.Host,
		config.DBPort,
		config.DBName,
=======
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
>>>>>>> 03362f06d487b54d41aeb62a1a3a89dd3f5a3e8b
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
func Migrate() {
	db := InitDB()
	db.AutoMigrate(&entities.User{}, &entities.Address{}, &entities.Cart{}, &entities.Category{}, entities.Product{}, entities.Transaction{})
}
