package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env ファイルが見つかりませんでしたが、環境変数を使用します")
	}

	// dbhost := os.Getenv("DB_HOST")
	// if dbhost == "" {
	// 	dbhost = "localhost"
	// }

	// dbport := os.Getenv("DB_PORT")
	// if dbport == "" {
	// 	dbport = "3306"
	// }

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		// dbhost,
		// dbport,
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
