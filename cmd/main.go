package main

import (
	"github.com/bpremika/post/db"
	"github.com/bpremika/post/internal/config"
	"github.com/bpremika/post/internal/model"
)

func main() {
	cfg := config.Load()
	// Initialize the DB connection
	db.InitDB(cfg)

	_ = db.DB.AutoMigrate(&model.LendingPost{})
	_ = db.DB.AutoMigrate(&model.BorrowingPost{})

	err := db.ServerInit(cfg, db.DB)
	if err != nil {
		panic(err)
	}
}
