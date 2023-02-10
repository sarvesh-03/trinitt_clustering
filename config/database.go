package config

import (
	"fmt"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database
var db *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var er error
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		DbHost, DbUser, DbPassword, DbName, DbPort)

	db, er = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if er != nil {
		fmt.Println(color.RedString("Error connecting to database"))
	} else {
		fmt.Println(color.GreenString("Database connected"))
	}
}

// GetDB returns the database
func GetDB() *gorm.DB {
	return db
}
