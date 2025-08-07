package examssubmissionanswer

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

func (h *Handler) Create(c *fiber.Ctx) error {
	data := new(ExamSubmissionAnswer)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := h.service.Create(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create data"})
	}
	return c.JSON(data)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Data not found"})
	}

	return c.JSON(data)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	data := new(ExamSubmissionAnswer)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	data.ID = uint(id)
	if err := h.service.Update(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update data"})
	}

	return c.JSON(data)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete"})
	}

	return c.SendStatus(204)
}
