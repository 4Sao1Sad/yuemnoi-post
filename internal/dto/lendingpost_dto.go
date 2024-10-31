package dto

import "time"

type CreateLendingPostRequest struct {
	ItemName    string
	Description string
	Price       float64
	ImageURL    string
}

type LendingPost struct {
	Id          uint64
	ItemName    string
	Description string
	Price       float64
	ImageURL    string
	OwnerId     uint64
	OwnerName   string
	CreatedAt   time.Time
}
