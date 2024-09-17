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
		ItemName:    input.ItemName,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    input.ImageUrl,
	}

	_, err := g.repository.InsertLendingPost(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateLendingPostResponse{
		Message: "created",
	}

	return &resp, nil
}

func (g *LendingPostGRPC) GetLendingPostDetail(ctx context.Context, input *pb.GetLendingPostDetailRequest) (*pb.LendingPost, error) {
	post, err := g.repository.GetLendingPostById(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingPost{
		Id:           uint64(post.ID),
		ItemName:     post.ItemName,
		Description:  post.Description,
		Price:        post.Price,
		ActiveStatus: post.ActiveStatus,
		ImageUrl:     post.ImageURL,
		OwnerId:      post.OwnerID,
		UpdatedAt:    timestamppb.New(post.UpdatedAt),
	}

	return &resp, nil
}

func (g *LendingPostGRPC) SearchLendingPost(ctx context.Context, input *pb.SearchLendingPostRequest) (*pb.LendingPostList, error) {
	posts, err := g.repository.SearchLendingPost(input.SearchString)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	postsResponse := []*pb.LendingPost{}
	for _, post := range *posts {
		postsResponse = append(postsResponse, &pb.LendingPost{
			Id:           uint64(post.ID),
			ItemName:     post.ItemName,
			Description:  post.Description,
			Price:        post.Price,
			ActiveStatus: post.ActiveStatus,
			ImageUrl:     post.ImageURL,
			OwnerId:      post.OwnerID,
			UpdatedAt:    timestamppb.New(post.UpdatedAt),
		})
	}

	resp := pb.LendingPostList{
		Posts: postsResponse,
	}

	return &resp, nil
}
