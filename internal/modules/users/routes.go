package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := app.Group("/users")
	api.Get("/", handler.GetUsers)
	api.Get("/pagination", handler.GetUsersByPage)
	api.Post("/", handler.CreateUser)
	api.Get("/:id", handler.GetUserByID)
	api.Put("/:id", handler.EditUser)
	api.Delete("/:id", handler.DeleteUser)
}
