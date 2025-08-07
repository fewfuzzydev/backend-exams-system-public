// cli/main.go
package main

//go run internal/cli/cli.go create-module modulesName

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "create-module" {
		fmt.Println("Usage: go run main.go create-module <ModuleName>")
		return
	}

	moduleName := os.Args[2]
	createModule(moduleName)
}

func createModule(name string) {
	modulePath := fmt.Sprintf("internal/modules/%s", name)
	os.MkdirAll(modulePath, os.ModePerm)

	files := []string{"model.go", "repository.go", "service.go", "handler.go", "routes.go"}

	for _, file := range files {
		content := renderTemplate(fileTemplate(file), name)
		os.WriteFile(fmt.Sprintf("%s/%s", modulePath, file), []byte(content), 0644)
	}

	fmt.Printf("✅ Module '%s' created successfully.\n", name)
}

func renderTemplate(tpl string, name string) string {
	var sb strings.Builder
	tmpl := template.Must(template.New("tpl").Funcs(template.FuncMap{
		"title": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToUpper(s[:1]) + s[1:]
		},
	}).Parse(tpl))
	tmpl.Execute(&sb, struct{ ModuleName string }{ModuleName: name})
	return sb.String()
}

func fileTemplate(name string) string {
	switch name {
	case "handler.go":
		return `package {{.ModuleName}}

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
    data := new({{title .ModuleName}})
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

    data := new({{title .ModuleName}})
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
`
	case "model.go":
		return `package {{.ModuleName}}

import (
"exams/internal/common"
)

type {{title .ModuleName}} struct {
{{title .ModuleName}}ID   uint
common.Audit
}`

	case "repository.go":
		return `package {{.ModuleName}}

import "gorm.io/gorm"

type Repository interface {
Create(data *{{title .ModuleName}}) error
FindAll() ([]{{title .ModuleName}}, error)
FindByID(id uint) (*{{title .ModuleName}}, error)
Update(data *{{title .ModuleName}}) error
Delete(id uint) error
}

type repository struct {
db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
db.AutoMigrate(&{{title .ModuleName}}{})
return &repository{db}
}

func (r *repository) Create(data *{{title .ModuleName}}) error {
return r.db.Create(data).Error
}

func (r *repository) FindAll() ([]{{title .ModuleName}}, error) {
var data []{{title .ModuleName}}
err := r.db.Find(&data).Error
return data, err
}

func (r *repository) FindByID(id uint) (*{{title .ModuleName}}, error) {
var data {{title .ModuleName}}
err := r.db.First(&data, id).Error
if err != nil {
	return nil, err
}
return &data, nil
}

func (r *repository) Update(data *{{title .ModuleName}}) error {
return r.db.Save(data).Error
}

func (r *repository) Delete(id uint) error {
return r.db.Delete(&{{title .ModuleName}}{}, id).Error
}`

	case "service.go":
		return `package {{.ModuleName}}

type Service interface {
GetAll() ([]{{title .ModuleName}}, error)
Create(data *{{title .ModuleName}}) error
GetByID(id uint) (*{{title .ModuleName}}, error)
Update(data *{{title .ModuleName}}) error
Delete(id uint) error
}

type service struct {
repo Repository
}

func NewService(repo Repository) Service {
return &service{repo}
}

func (s *service) GetAll() ([]{{title .ModuleName}}, error) {
return s.repo.FindAll()
}

func (s *service) Create(data *{{title .ModuleName}}) error {
return s.repo.Create(data)
}

func (s *service) GetByID(id uint) (*{{title .ModuleName}}, error) {
return s.repo.FindByID(id)
}

func (s *service) Update(data *{{title .ModuleName}}) error {
return s.repo.Update(data)
}

func (s *service) Delete(id uint) error {
return s.repo.Delete(id)
}`

	case "routes.go":
		return `package {{.ModuleName}}

import (
middlewares "exams/internal/middleware"

"github.com/gofiber/fiber/v2"
"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
repo := NewRepository(db)
service := NewService(repo)
handler := NewHandler(service)

api := app.Group("/{{.ModuleName}}", middlewares.JWTProtected())
api.Get("/", handler.Get)
api.Post("/", handler.Create)
api.Get("/:id", handler.GetByID)
api.Put("/:id", handler.Edit)
api.Delete("/:id", handler.Delete)
}`

	}

	return ""
}
