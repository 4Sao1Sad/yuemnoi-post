package handler

import (
	"strconv"

	"github.com/bpremika/post/internal/dto"
	"github.com/bpremika/post/internal/model"
	"github.com/bpremika/post/internal/repository"
	util "github.com/bpremika/post/internal/util"
	"github.com/gofiber/fiber/v2"
)

type LendingPostRest struct {
	repository repository.LendingPostRepository
}

func NewLendingPostRest(repo repository.LendingPostRepository) *LendingPostRest {
	return &LendingPostRest{
		repository: repo,
	}
}

func (g *LendingPostRest) CreateLendingPost(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
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

	body := new(dto.CreateLendingPostRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	name, err := util.GetUserById(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "This user not exists",
		})
	}

	data := model.LendingPost{
		ItemName:    body.ItemName,
		Description: body.Description,
		Price:       body.Price,
		ImageURL:    body.ImageURL,
		OwnerID:     uint64(userId),
		OwnerName:   name,
	}

	_, err = g.repository.InsertLendingPost(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create lending post",
		})
	}

	postId := strconv.FormatUint(uint64(data.ID), 10)
	logDetail := "Post Service: Create post for lending, id = " + postId
	err = util.CallActivityLogService(uint64(data.OwnerID), logDetail)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (g *LendingPostRest) SearchLendingPost(c *fiber.Ctx) error {
	searchString := c.Query("search")

	posts, err := g.repository.SearchLendingPost(searchString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search lending post",
		})
	}

	resp := []*dto.LendingPost{}
	for _, post := range *posts {
		resp = append(resp, &dto.LendingPost{
			Id:          uint64(post.ID),
			ItemName:    post.ItemName,
			Description: post.Description,
			Price:       post.Price,
			ImageURL:    post.ImageURL,
			OwnerId:     post.OwnerID,
			OwnerName:   post.OwnerName,
			CreatedAt:   post.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

func (g *LendingPostRest) GetMyLendingPosts(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
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

	posts, err := g.repository.GetMyLendingPosts(uint64(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get my lending posts",
		})
	}

	resp := []*dto.LendingPost{}
	for _, post := range *posts {
		resp = append(resp, &dto.LendingPost{
			Id:          uint64(post.ID),
			ItemName:    post.ItemName,
			Description: post.Description,
			Price:       post.Price,
			ImageURL:    post.ImageURL,
			OwnerId:     post.OwnerID,
			OwnerName:   post.OwnerName,
			CreatedAt:   post.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
