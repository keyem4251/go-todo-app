package infrastructure

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB

func Init() *gorm.DB {
	fmt.Println("start db...")
	dsn := "root:password@tcp(mysql-container:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("db setup done")
	return db
}

func GetDB() *gorm.DB {
	return db
}
