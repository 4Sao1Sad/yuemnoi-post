package model

import (
	"gorm.io/gorm"
)

type BorrowingPost struct {
    gorm.Model
    Description  string        `gorm:"type:text;not null"`
    ActiveStatus bool          `gorm:"not null"`
    ItemList     []LendingPost `gorm:"many2many:borrowing_post_item_list;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    OwnerID      string        `gorm:"type:varchar(100);not null"`
}