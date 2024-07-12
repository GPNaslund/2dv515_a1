package similarusers

import "github.com/gofiber/fiber/v2"

type SimilarUsersService interface {
  getSimilarUsers(map[string]string) (string, error)
}

type Handler struct {
  service SimilarUsersService
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c *fiber.Ctx) error {
  queryParams := c.Queries()
  result, err := h.service.getSimilarUsers(queryParams)

  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "error": "Something went wrong",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "ok": result,
  })
}
