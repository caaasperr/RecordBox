package handler

import (
	"myvinyl/models"
	"myvinyl/modules/shelf/controller"
	"myvinyl/modules/vinyl"
	vCont "myvinyl/modules/vinyl/controller"
	"myvinyl/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateVinyl(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	var reqForm CreateVinylRequest
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errValidating := reqForm.Validate()
	if errValidating != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := vCont.CreateVinyl(user.ID, reqForm.GenreID, reqForm.ShelfslotID, reqForm.Name, reqForm.Artist, reqForm.Detail, reqForm.Price, reqForm.ImageURL, reqForm.Format, reqForm.Sleeve, reqForm.Media, reqForm.ReleasedDate)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Couldn't create vinyl",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetVinylsHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	vinyls, err := vCont.GetAllVinylsByUserID(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	return c.Status(fiber.StatusOK).JSON(vinyls)
}

func GetVinylHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	str := c.Params("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyl",
		})
	}
	tVinyl, terr := vCont.GetVinylByIDForHandler(uint(id))
	if terr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyl",
		})
	}
	if vinyl.IsOwned(user, uint(id)) != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	return c.Status(fiber.StatusOK).JSON(tVinyl)
}

func DeleteVinylHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	str := c.Params("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete vinyls",
		})
	}
	if vinyl.IsOwned(user, uint(id)) != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not delete vinyls",
		})
	}
	vCont.DeleteVinyl(uint(id))
	return c.SendStatus(fiber.StatusAccepted)
}

func UpdateVinylHandler(c *fiber.Ctx) error {
	var updateData UpdateVinylRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	errValidating := updateData.Validate()
	if errValidating != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body (Vallidation)",
		})
	}
	user := c.Locals("user").(models.User)

	str := c.Params("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	if vinyl.IsOwned(user, uint(id)) != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	errU := vCont.UpdateVinyl(uint(id), updateData.GenreID, updateData.ShelfslotID, updateData.Name, updateData.Artist, updateData.Detail, updateData.Price, updateData.ImageURL, updateData.Format, updateData.Sleeve, updateData.Media, updateData.ReleasedDate)
	if errU != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func UpdateVinylSlotHandler(c *fiber.Ctx) error {
	var updateData UpdateVinylSlotRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	user := c.Locals("user").(models.User)

	str := c.Params("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	if vinyl.IsOwned(user, uint(id)) != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	errU := vCont.UpdateVinylSlot(uint(id), updateData.ShelfslotID)
	if errU != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update vinyls",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func GetVinylsByShelfSlotHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	str := c.Params("id")
	str2 := c.Params("value")
	sid, errrr := strconv.Atoi(str)
	if errrr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	id, err := strconv.Atoi(str2)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}

	ss, err := controller.GetBookShelf(uint(sid))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	if ss.UserID != user.ID {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}

	ssl, err := controller.GetShelfslot(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	if sid != int(ssl.BookshelfID) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}

	vinyls, err := vCont.GetVinylsByShelfSlot(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get vinyls",
		})
	}
	return c.Status(fiber.StatusOK).JSON(vinyls)
}

func GetAlbumCoversFromLastFmByNameHandler(c *fiber.Ctx) error {
	var reqForm GetAlbumCoversRequest
	if err := c.BodyParser(&reqForm); err != nil {
		utils.Logger.Warn("Invalid request body", reqForm)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	var albumResponseS vinyl.AlbumResponse
	albumResponseS, err := vinyl.GetAlbumCoverFromSpecificInformation(reqForm.Name, reqForm.Artist)
	if err != nil {
		albumResponse, err := vinyl.GetAlbumCoversFromLastFmByName(reqForm.Name)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		albumImages := []fiber.Map{}
		for _, album := range albumResponse.Results.Albummatches.Album {
			for _, image := range album.Image {
				if image.Size == "extralarge" && image.Text != "" {
					albumImages = append(albumImages, fiber.Map{
						"name":   album.Name,
						"artist": album.Artist,
						"url":    image.Text,
					})
				}
			}
		}
		return c.JSON(fiber.Map{
			"albums": albumImages,
		})
	}
	var albumCover string
	for _, image := range albumResponseS.Album.Image {
		if image.Size == "extralarge" {
			albumCover = image.Text
			break
		}
	}

	return c.JSON(fiber.Map{
		"albums": append([]fiber.Map{}, fiber.Map{"name": reqForm.Name, "artist": reqForm.Artist, "url": albumCover}),
	})
}
