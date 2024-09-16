package handler

import (
	"context"

	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	pb "github.com/bpremika/post/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"fmt"
)

type BorrowingPostGRPC struct {
	pb.UnimplementedBorrowingPostServiceServer
	repository repository.BorrowingPostRepository
}

func NewBorrowingPostGRPC(repo repository.BorrowingPostRepository) *BorrowingPostGRPC {
	return &BorrowingPostGRPC{
		repository: repo,
	}
}

func (g *BorrowingPostGRPC) CreateBorrowingPost(ctx context.Context, input *pb.CreateBorrowingPostRequest) (*pb.CreateBorrowingPostResponse, error) {
	data := model.BorrowingPost{
		Description:  input.Description,
		ActiveStatus: input.ActiveStatus,
		OwnerID:      input.OwnerId,
	}
	fmt.Println(data)

	_, err := g.repository.InsertBorrowingPost(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateBorrowingPostResponse{
		Message: "success ja",
	}

	return &resp, nil
}

func (g *BorrowingPostGRPC) GetBorrowingPostDetail(ctx context.Context, input *pb.GetBorrowingPostDetailRequest) (*pb.GetBorrowingPostDetailResponse, error) {
	post, err := g.repository.GetBorrowingPostById(uint(input.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.GetBorrowingPostDetailResponse{
		Description:  post.Description,
		ActiveStatus: post.ActiveStatus,
	}

	return &resp, nil
}

func (g *BorrowingPostGRPC) SearchBorrowingPost(ctx context.Context, input *pb.SearchBorrowingPostRequest) (*pb.SearchBorrowingPostResponse, error) {
	posts, err := g.repository.SearchBorrowingPost(input.SearchString)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var resp pb.SearchBorrowingPostResponse
	for _, post := range *posts {
		resp.BorrowingPost = append(resp.BorrowingPost, &pb.BorrowingPost{
			Description:  post.Description,
			ActiveStatus: post.ActiveStatus,
		})
	}

	return &resp, nil
}