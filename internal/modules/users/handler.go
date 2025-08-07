package users

import (
	"exams/internal/upload"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	role := c.FormValue("role")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	department := c.FormValue("department")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse multipart form"})
	}

	files := form.File["files"]
	var savedPaths []string
	for _, file := range files {
		path, err := upload.SaveSingleFile(file, "./uploads/multifile")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
		}
		savedPaths = append(savedPaths, path)
	}

	var filePath string

	file, err := c.FormFile("profile_image")
	if err == nil {
		path, err := upload.SaveSingleFile(file, "./uploads/profiles")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "upload failed"})
		}
		fmt.Println("Uploaded path:", path)
		filePath = path
	}

	req := CreateUserRequest{
		Username:         username,
		Password:         password,
		Role:             role,
		Firstname:        firstname,
		Lastname:         lastname,
		Email:            email,
		Phone:            phone,
		Department:       department,
		ProfileImagePath: filePath,
		Files:            savedPaths,
	}

	if err := h.service.CreateUserWithTeacher(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user & teacher"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}
	return c.JSON(users)
}

func (h *Handler) GetUsersByPage(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, total, err := h.service.GetUsersPaginated(page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"total": total,
	})
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (h *Handler) EditUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	user.ID = uint(id)
	if err := h.service.UpdateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(204)
}
