package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", c.Params("*"))
	return c.SendString(msg) // => ✋ register
}
