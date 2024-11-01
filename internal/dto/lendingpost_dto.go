package dto

import "time"

type CreateLendingPostRequest struct {
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

type LendingPost struct {
	Id          uint64    `json:"id"`
	ItemName    string    `json:"item_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	OwnerId     uint64    `json:"owner_id"`
	OwnerName   string    `json:"owner_name"`
	CreatedAt   time.Time `json:"created_at"`
}
