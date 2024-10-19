package handler

import (
	"context"
	"strconv"

	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	util "github.com/bpremika/post/internal/util"
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

	postId := strconv.FormatUint(uint64(data.ID), 10)
	logDetail := "Post Service: Create post for lending, id = " + postId
	err = util.CallActivityLogService(uint64(data.OwnerID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Post Service: Create post for lending. %v", err)
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

	postId := strconv.FormatUint(uint64(post.ID), 10)
	logDetail := "Post Service: View lending post " + postId + " details"
	err = util.CallActivityLogService(uint64(post.OwnerID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Post Service: View lending post %d details. %v", post.ID, err)
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
