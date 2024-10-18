package repository

import (
	"github.com/bpremika/post/internal/model"
	"gorm.io/gorm"
)

type LendingPostRepositoryImpl struct {
	db *gorm.DB
}

type LendingPostRepository interface {
	InsertLendingPost(post model.LendingPost) (*model.LendingPost, error)
	GetLendingPostById(id uint64) (*model.LendingPost, error)
	SearchLendingPost(searchString string) (*[]model.LendingPost, error)
	UpdateLendingPost(updatedFields map[string]interface{}) (*model.LendingPost, error)
}

func NewLendingPostRepository(db *gorm.DB) LendingPostRepository {
	return &LendingPostRepositoryImpl{db}
}

func (r LendingPostRepositoryImpl) InsertLendingPost(post model.LendingPost) (*model.LendingPost, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r LendingPostRepositoryImpl) GetLendingPostById(id uint64) (*model.LendingPost, error) {
	var post model.LendingPost
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r LendingPostRepositoryImpl) SearchLendingPost(searchString string) (*[]model.LendingPost, error) {
	var posts []model.LendingPost
	err := r.db.Where("item_name ILIKE ?", "%"+searchString+"%").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r LendingPostRepositoryImpl) UpdateLendingPost(updatedFields map[string]interface{}) (*model.LendingPost, error) {
	var post model.LendingPost
	err := r.db.Model(&post).Updates(updatedFields).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}
