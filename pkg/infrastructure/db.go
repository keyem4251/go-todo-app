package infrastructure

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func Init() *gorm.DB {
	fmt.Println("start db...")
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("db setup done")
	return db
}

func GetDB() *gorm.DB {
	return db
}
