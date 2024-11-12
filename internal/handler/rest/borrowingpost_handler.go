package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bpremika/post/internal/dto"
	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	util "github.com/bpremika/post/internal/util"
	"github.com/gofiber/fiber/v2"
)

type BorrowingPostRest struct {
	repository repository.BorrowingPostRepository
}

func NewBorrowingPostRest(repo repository.BorrowingPostRepository) *BorrowingPostRest {
	return &BorrowingPostRest{
		repository: repo,
	}
}

func (g *BorrowingPostRest) CreateBorrowingPost(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
	fmt.Println(userIdString)
	if userIdString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	body := new(dto.CreateBorrowingPostRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	name, err := util.GetUserById(uint(userId))
	if err != nil {
		log.Print("error from get userId", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "This user not exists",
		})
	}

	data := model.BorrowingPost{
		Description: body.Description,
		OwnerID:     uint64(userId),
		OwnerName:   name,
	}

	_, err = g.repository.InsertBorrowingPost(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create borrowing post",
		})
	}

	postId := strconv.FormatUint(uint64(data.ID), 10)
	logDetail := "Post Service: Create post for borrowing, id = " + postId
	err = util.CallActivityLogService(uint64(data.OwnerID), logDetail)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (g *BorrowingPostRest) SearchBorrowingPost(c *fiber.Ctx) error {
	searchString := c.Query("search")

	posts, err := g.repository.SearchBorrowingPost(searchString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search borrowing post",
		})
	}

	var resp []*dto.BorrowingPost
	for _, post := range *posts {
		resp = append(resp, &dto.BorrowingPost{
			Id:          uint64(post.ID),
			OwnerId:     post.OwnerID,
			OwnerName:   post.OwnerName,
			Description: post.Description,
			CreatedAt:   post.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (g *BorrowingPostRest) GetMyBorrowingPosts(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
	fmt.Println("userIdString", userIdString)
	if userIdString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	fmt.Println("userId", userId)
	posts, err := g.repository.GetMyBorrowingPosts(uint64(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get my borrowing posts",
		})
	}

	resp := []*dto.BorrowingPost{}
	for _, post := range *posts {
		resp = append(resp, &dto.BorrowingPost{
			Id:          uint64(post.ID),
			OwnerId:     post.OwnerID,
			OwnerName:   post.OwnerName,
			Description: post.Description,
			CreatedAt:   post.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
