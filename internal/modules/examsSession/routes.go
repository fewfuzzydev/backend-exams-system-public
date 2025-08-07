package examssession

import (
	middlewares "exams/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := app.Group("/examssession", middlewares.JWTProtected())
	api.Get("/", handler.Get)
	api.Post("/", handler.Create)
	api.Get("/:id", handler.GetByID)
	app.Put("/examssession/:id", handler.Edit)
	app.Delete("/examssession/:id", handler.Delete)

}
