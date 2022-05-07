package main

import (
	"gorm.io/gorm"
)

func main() {
	fmt.Println(config.InitDB())
}
