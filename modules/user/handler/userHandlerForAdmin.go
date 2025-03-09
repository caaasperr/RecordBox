package handler

import (
	"myvinyl/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	result, err := controller.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can't get all users.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
