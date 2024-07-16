package movierecommendations

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MovieRecommendationsService interface {
  GetMovieRecommendations(ctx context.Context, queryParams map[string]string) ([]MovieRecommendation, error)
}

type Handler struct {
  service MovieRecommendationsService
}

func NewHandler(service MovieRecommendationsService) *Handler {
	return &Handler{
    service: service,
  }
}

func (h *Handler) Handle(c *fiber.Ctx) error {
	queryParams := c.Queries()
	result, err := h.service.GetMovieRecommendations(c.Context(), queryParams)

	if err != nil {
		log.Printf("Error occured: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"movies": result,
	})
}
