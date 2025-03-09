package handler

import (
	"myvinyl/modules/user"
	"myvinyl/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

// Middleware for authorized user
func UserMiddleware(c *fiber.Ctx) error {
	usernameStr, err := user.GetUserBySession(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session data",
		})
	}

	user, err := controller.GetUserById(usernameStr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	c.Locals("user", user)
	return c.Next()
}

// Middleware for admin
func AdminMiddleware(c *fiber.Ctx) error {
	usernameStr, err := user.GetUserBySession(c)
	if err != nil {
		c.Redirect("/")
	}
	user, err := controller.GetUserById(usernameStr)
	if err != nil {
		return c.Redirect("/")
	}
	if !user.IsAdmin {
		return c.Redirect("/")
	}
	return c.Next()
}
