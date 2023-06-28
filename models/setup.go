package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	var dns = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns, // 账号密码地址端口
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		return
	}

	DB = db
}
