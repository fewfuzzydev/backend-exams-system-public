package examsQuestions

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get data"})
	}
	return c.JSON(data)
}
