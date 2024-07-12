package movierecommendations

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c *fiber.Ctx) error {
  return fmt.Errorf("Not implemented")
}
