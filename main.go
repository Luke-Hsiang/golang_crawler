package main

import (
	"fmt"
	"golang_crawler/apps/news"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("error open db: %v\n", err)
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	news.Init(db, v1)
	r.Run()
}
