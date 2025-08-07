package teachers

import (
	middlewares "exams/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := app.Group("/teachers", middlewares.JWTProtected())
	api.Get("/", handler.Get)
	api.Post("/", handler.Create)
	api.Get("/:id", handler.GetByID)
	app.Put("/teachers/:id", handler.Edit)
	app.Delete("/teachers/:id", handler.Delete)

}
