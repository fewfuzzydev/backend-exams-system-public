package middlewares

import (
	"exams/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])
		return c.Next()
	}
}

func AdminProtected() {
	fmt.Print("AdminProtect")
}
