package model

import (
	"gorm.io/gorm"
)

type LendingPost struct {
	gorm.Model
	ItemName     string  `gorm:"type:varchar(255);not null"`
	Description  string  `gorm:"type:text"`
	Price        float64 `gorm:"type:numeric;"`
	ActiveStatus bool    `gorm:"not null"`
	ImageURL     string  `gorm:"type:varchar(255);"`
	OwnerID      string  `gorm:"type:varchar(255);"`
}
