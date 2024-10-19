package handler

import (
	"context"

	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	pb "github.com/bpremika/post/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Description: input.Description,
		OwnerID:     input.OwnerId,
	}

	_, err := g.repository.InsertBorrowingPost(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateBorrowingPostResponse{
		Message: "created",
	}

	return &resp, nil
}

func (g *BorrowingPostGRPC) GetBorrowingPostDetail(ctx context.Context, input *pb.GetBorrowingPostDetailRequest) (*pb.BorrowingPost, error) {
	post, err := g.repository.GetBorrowingPostById(uint(input.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.BorrowingPost{
		Id:           uint64(post.ID),
		Description:  post.Description,
		ActiveStatus: post.ActiveStatus,
		UpdatedAt:    timestamppb.New(post.UpdatedAt),
		OwnerId:      post.OwnerID,
	}

	return &resp, nil
}

func (g *BorrowingPostGRPC) SearchBorrowingPost(ctx context.Context, input *pb.SearchBorrowingPostRequest) (*pb.BorrowingPostList, error) {
	posts, err := g.repository.SearchBorrowingPost(input.SearchString)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var resp pb.BorrowingPostList
	for _, post := range *posts {
		resp.BorrowingPost = append(resp.BorrowingPost, &pb.BorrowingPost{
			Description:  post.Description,
			ActiveStatus: post.ActiveStatus,
		})
	}

	return &resp, nil
}

func (g *BorrowingPostGRPC) UpdateBorrowingPost(ctx context.Context, input *pb.UpdateBorrowingPostRequest) (*pb.UpdateBorrowingPostResponse, error) {
	updates := make(map[string]interface{})

	if input.UpdateMask != nil {
		for _, path := range input.UpdateMask.Paths {
			switch path {
			case "description":
				updates["Description"] = input.Description
			case "active_status":
				updates["ActiveStatus"] = input.ActiveStatus
			}
		}
	}

	_, err := g.repository.UpdateBorrowingPost(uint(input.Id), updates)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.UpdateBorrowingPostResponse{
		Message: "updated",
	}

	return &resp, nil
}
