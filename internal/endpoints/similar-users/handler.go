package similarusers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type SimilarUsersService interface {
  GetSimilarUsers(ctx context.Context, queryParams map[string]string) ([]SimilarityScore, error)
}

type Handler struct {
  service SimilarUsersService
}

func NewHandler(service SimilarUsersService) *Handler {
	return &Handler{
    service: service,
  }
}

func (h *Handler) Handle(c *fiber.Ctx) error {
  queryParams := c.Queries()
  result, err := h.service.GetSimilarUsers(c.Context(), queryParams)

  if err != nil {
    log.Printf("Error occured: %s", err)
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Something went wrong",
    })
  }

  return c.Status(fiber.StatusOK).JSON(result)
}
