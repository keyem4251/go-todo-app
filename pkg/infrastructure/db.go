package infrastructure

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	fmt.Println("start db...")
	dsn := "root:password@tcp(mysql:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	
	retry_count := 0
	for {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("failed to connect database and retry count %d\n", retry_count)
			retry_count += 1
			time.Sleep(time.Duration(retry_count * 10) * time.Second)
		} else {
			fmt.Println("db setup done")
			return db
		}
		if retry_count > 3 {
			break
		}
	}
	panic("failed to connect database")
}

func GetDB() *gorm.DB {
	return db
}
