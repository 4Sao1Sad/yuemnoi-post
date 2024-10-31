package handler

import (
	"context"
	"strconv"

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

func (g *LendingPostGRPC) GetLendingPostsByIds(ctx context.Context, input *pb.GetLendingPostsByIdsRequest) (*pb.LendingPostList, error) {
	posts, err := g.repository.GetLendingPostsByIds(input.Ids)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingPostList{}
	for _, post := range *posts {
		resp.Posts = append(resp.Posts, &pb.LendingPost{
			Id:           uint64(post.ID),
			Description:  post.Description,
			ActiveStatus: post.ActiveStatus,
			UpdatedAt:    timestamppb.New(post.UpdatedAt),
			OwnerId:      post.OwnerID,
			OwnerName:    post.OwnerName,
		})
	}

	return &resp, nil
}

func (g *LendingPostGRPC) UpdateLendingPost(ctx context.Context, input *pb.UpdateLendingPostRequest) (*pb.UpdateLendingPostResponse, error) {
	updates := make(map[string]interface{})

	if input.UpdateMask != nil {
		for _, path := range input.UpdateMask.Paths {
			switch path {
			case "item_name":
				updates["ItemName"] = input.ItemName
			case "description":
				updates["Description"] = input.Description
			case "price":
				updates["Price"] = input.Price
			case "active_status":
				updates["ActiveStatus"] = input.ActiveStatus
			case "image_url":
				updates["ImageURL"] = input.ImageUrl
			}
		}
	}
	_, err := g.repository.UpdateLendingPost(uint(input.Id), updates)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.UpdateLendingPostResponse{
		Message: "updated",
	}

	return &resp, nil
}
