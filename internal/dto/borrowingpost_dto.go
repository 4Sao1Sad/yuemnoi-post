package dto

import "time"

type CreateBorrowingPostRequest struct {
	Description string `json:"description"`
}

type BorrowingPost struct {
	Id          uint64    `json:"id"`
	OwnerId     uint64    `json:"owner_id"`
	OwnerName   string    `json:"owner_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
