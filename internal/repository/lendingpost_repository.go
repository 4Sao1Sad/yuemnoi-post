package repository

import (
	"fmt"

	"github.com/bpremika/post/internal/model"
	"gorm.io/gorm"
)

type LendingPostRepositoryImpl struct {
	db *gorm.DB
}

type LendingPostRepository interface {
	InsertLendingPost(post model.LendingPost) (*model.LendingPost, error)
}

func NewLendingPostRepository(db *gorm.DB) LendingPostRepository {
	return &LendingPostRepositoryImpl{db}
}

func (r LendingPostRepositoryImpl) InsertLendingPost(post model.LendingPost) (*model.LendingPost, error) {
	fmt.Println("repo")
	err := r.db.Create(&post).Error
	fmt.Println("1", err)
	if err != nil {
		fmt.Println("here")
		return nil, err
	}
	fmt.Println("2", err)
	return &post, nil
}
