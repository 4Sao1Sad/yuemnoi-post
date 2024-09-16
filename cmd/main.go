package main

import (
	"log"

	"github.com/bpremika/post/db"
	"github.com/bpremika/post/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the DB connection
	db.InitDB()

	_ = db.DB.AutoMigrate(&model.LendingPost{})
	_ = db.DB.AutoMigrate(&model.BorrowingPost{})

	err = db.ServerInit(db.DB)
	if err != nil {
		panic(err)
	}
}
