package model

import (
	"gorm.io/gorm"
)

type BorrowingPost struct {
	gorm.Model
	Description  string `gorm:"type:text;not null"`
	ActiveStatus bool   `gorm:"not null;default:true"`
	OwnerID      uint64 `gorm:"type:varchar(100);not null"`
}
