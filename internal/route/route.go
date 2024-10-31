package route

import (
	"github.com/bpremika/post/internal/config"
	handler "github.com/bpremika/post/internal/handler/rest"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	borrowingPostRestHandler *handler.BorrowingPostRest
	lendingPostRestHandler   *handler.LendingPostRest
}

func NewHandler(borrowingPostRestHandler *handler.BorrowingPostRest, lendingPostRestHandler *handler.LendingPostRest) *Handler {
	return &Handler{
		borrowingPostRestHandler: borrowingPostRestHandler,
		lendingPostRestHandler:   lendingPostRestHandler,
	}
}

func (h *Handler) RegisterRouter(r fiber.Router, cfg *config.Config) {
	{
		borrowingPostRouter := r.Group("/borrowing-posts")
		borrowingPostRouter.Get("/:search", h.borrowingPostRestHandler.SearchBorrowingPost)
		borrowingPostRouter.Get("/me", h.borrowingPostRestHandler.GetMyBorrowingPosts)
		borrowingPostRouter.Post("/", h.borrowingPostRestHandler.CreateBorrowingPost)
	}
	{
		LendingPostRouter := r.Group("/lending-posts")
		LendingPostRouter.Get("/:search", h.lendingPostRestHandler.SearchLendingPost)
		LendingPostRouter.Get("/me", h.lendingPostRestHandler.GetMyLendingPosts)
		LendingPostRouter.Post("/", h.lendingPostRestHandler.CreateLendingPost)
	}
}
