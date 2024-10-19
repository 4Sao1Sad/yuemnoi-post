package repository

import (
	"github.com/bpremika/post/internal/model"
	"gorm.io/gorm"
)

type BorrowingPostRepositoryImpl struct {
	db *gorm.DB
}

type BorrowingPostRepository interface {
	InsertBorrowingPost(post model.BorrowingPost) (*model.BorrowingPost, error)
	GetBorrowingPostById(id uint) (*model.BorrowingPost, error)
	SearchBorrowingPost(searchString string) (*[]model.BorrowingPost, error)
	UpdateBorrowingPost(id uint, updatedFields map[string]interface{}) (*model.BorrowingPost, error)
}

func NewBorrowingPostRepository(db *gorm.DB) BorrowingPostRepository {
	return &BorrowingPostRepositoryImpl{db}
}

func (r BorrowingPostRepositoryImpl) InsertBorrowingPost(post model.BorrowingPost) (*model.BorrowingPost, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r BorrowingPostRepositoryImpl) GetBorrowingPostById(id uint) (*model.BorrowingPost, error) {
	var post model.BorrowingPost
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r BorrowingPostRepositoryImpl) SearchBorrowingPost(searchString string) (*[]model.BorrowingPost, error) {
	var posts []model.BorrowingPost
	err := r.db.Where("description ILIKE ?", "%"+searchString+"%").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r BorrowingPostRepositoryImpl) UpdateBorrowingPost(id uint, updatedFields map[string]interface{}) (*model.BorrowingPost, error) {
	var post model.BorrowingPost
	err := r.db.Model(&post).Where("id = ?", id).Updates(updatedFields).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}
