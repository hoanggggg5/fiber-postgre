package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var Store = session.New()
var Database *gorm.DB

func InitConfig() {
	dsn := "host=localhost user=root password=123456 dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	Database = db
}
