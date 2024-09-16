package handler

import (
	"context"

	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	pb "github.com/bpremika/post/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LendingPostGRPC struct {
	pb.UnimplementedLendingPostServiceServer
	repository repository.LendingPostRepository
}

func NewLendingPostGRPC(repo repository.LendingPostRepository) *LendingPostGRPC {
	return &LendingPostGRPC{
		repository: repo,
	}
}

func (g *LendingPostGRPC) CreateLendingPost(ctx context.Context, input *pb.CreateLendingPostRequest) (*pb.CreateLendingPostResponse, error) {
	data := model.LendingPost{
		ItemName:     input.ItemName,
		Price:        input.Price,
		ActiveStatus: input.ActiveStatus,
	}

	_, err := g.repository.InsertLendingPost(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateLendingPostResponse{
		Message: "success ja",
	}

	return &resp, nil
}

// func (g *LendingPostGRPC) GetLendingPostById(ctx context.Context, id uint) (*)