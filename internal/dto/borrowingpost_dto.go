package dto

import "time"

type CreateBorrowingPostRequest struct {
	Description string
}

type BorrowingPost struct {
	Id          uint64
	OwnerId     uint64
	OwnerName   string
	Description string
	CreatedAt   time.Time
}
