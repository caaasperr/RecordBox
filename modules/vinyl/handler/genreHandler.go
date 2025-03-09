package handler

import (
	"myvinyl/models"
	"myvinyl/modules/vinyl/controller"
	"myvinyl/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReqForm struct {
	Name string
}

func CreateGenreHandler(c *fiber.Ctx) error {
	var genre ReqForm
	if err := c.BodyParser(&genre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := controller.CreateGenre(genre.Name)
	if err != nil {
		utils.Logger.Warn("Couldn't create genre: ", err)
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not create genre",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func DeleteGenreHandler(c *fiber.Ctx) error {
	str := c.Params("id")
	gid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't delete genre: ", err)
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not delete genre",
		})
	}
	derr := controller.DeleteGenre(uint(gid))
	if derr != nil {
		utils.Logger.Warn("Couldn't delete genre: ", derr)
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not delete genre",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func GetGenresHandler(c *fiber.Ctx) error {
	genres, err := controller.GetGenres()
	if err != nil {
		utils.Logger.Warn("Couldn't get genres: ", err)
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not get genres",
		})
	}
	return c.Status(fiber.StatusOK).JSON(genres)
}

func GetVinylsByGenreHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	str := c.Params("id")
	gid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't get vinyls: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	vinyls, err := controller.GetVinylsByGenre(uint(gid), user.ID)
	if err != nil {
		utils.Logger.Warn("Couldn't get vinyls: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	return c.Status(fiber.StatusOK).JSON(vinyls)
}
