package handler

import (
	"myvinyl/models"
	"myvinyl/modules/shelf"
	"myvinyl/modules/shelf/controller"
	"myvinyl/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateShelfHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting CreateShelfHandler")
	user := c.Locals("user").(models.User)
	var reqForm CreateShelfReq

	// Body Parsing err check
	if err := c.BodyParser(&reqForm); err != nil {
		utils.Logger.Warn("Couldn't create shelf: ", err)
		utils.Logger.Trace("Ending CreateShelfHandler")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check request form
	errValidating := ValidateCreateShelf(reqForm)
	if errValidating != nil {
		utils.Logger.Warn("Couldn't create shelf: ", errValidating)
		utils.Logger.Trace("Ending CreateShelfHandler")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Create Shelf and check
	errCreate := shelf.CreateShelf(user.ID, reqForm.Name, reqForm.Detail, uint(reqForm.Columns), uint(reqForm.Rows))
	if errCreate != nil {
		utils.Logger.Warn("Couldn't create shelf: ", errCreate)
		utils.Logger.Trace("Ending CreateShelfHandler")
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Couldn't create shelf",
		})
	}
	utils.Logger.Trace("Ending CreateShelfHandler")
	return c.SendStatus(fiber.StatusOK)
}

func GetShelvesHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting GetShelvesHandler")
	user := c.Locals("user").(models.User)
	shelves, err := controller.GetBookshelves(user.ID)
	if err != nil {
		utils.Logger.Warn("Couldn't get shelves: ", err)
		utils.Logger.Trace("Ending GetShelvesHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't get shelves",
		})
	}
	utils.Logger.Trace("Ending GetShelvesHandler")
	return c.Status(fiber.StatusOK).JSON(shelves)
}

func GetShelfHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting GetShelfHandler")
	user := c.Locals("user").(models.User)

	// Get Param and convert
	str := c.Params("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't get shelf: ", err)
		utils.Logger.Trace("Ending GetShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't get shelf",
		})
	}

	// Get Shelf and check
	shelf, err := controller.GetBookShelf(uint(id))
	if err != nil {
		utils.Logger.Warn("Couldn't get shelf: ", err)
		utils.Logger.Trace("Ending GetShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't get shelf",
		})
	}

	// Check Ownership
	if shelf.UserID != user.ID {
		utils.Logger.Warn("Couldn't get vinyls: ", shelf.UserID, user.ID)
		utils.Logger.Trace("Ending GetShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't get shelf",
		})
	}

	// Get Shelf and slots
	shelfslots, errGetSlots := controller.GetShelfslotByShelfId(uint(id))
	if errGetSlots != nil {
		utils.Logger.Warn("Couldn't get shelf: ", errGetSlots)
		utils.Logger.Trace("Ending GetShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't get shelf",
		})
	}

	utils.Logger.Trace("Ending GetShelfHandler")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ID":      shelf.ID,
		"Name":    shelf.Name,
		"Detail":  shelf.Detail,
		"Columns": shelf.Columns,
		"Rows":    shelf.Rows,
		"Slots":   shelfslots,
	})
}

func DeleteShelfHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting DeleteShelfHandler")
	// Get Param and convert
	str := c.Params("id")
	gid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't delete shelf: ", err)
		utils.Logger.Trace("Ending DeleteShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete shelf",
		})
	}

	user := c.Locals("user").(models.User)
	shelfT, errGetShelf := controller.GetBookShelf(uint(gid))
	if errGetShelf != nil {
		utils.Logger.Warn("Couldn't delete shelf: ", errGetShelf)
		utils.Logger.Trace("Ending DeleteShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't delete shelf",
		})
	}

	// Check Ownership
	if shelfT.UserID != user.ID {
		utils.Logger.Warn("Couldn't delete shelf: unauthorized access")
		utils.Logger.Trace("Ending DeleteShelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't delete shelf",
		})
	}

	shelf.DeleteShelfByShelfId(shelfT.ID)
	utils.Logger.Trace("Shelf deleted. ID: ", shelfT.ID)
	utils.Logger.Trace("Ending DeleteShelfHandler")
	return c.SendStatus(fiber.StatusOK)
}

func UpdateBookshelfHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting UpdateBookshelfHandler")
	var reqForm UpdateShelfReq
	if err := c.BodyParser(&reqForm); err != nil {
		utils.Logger.Warn("Couldn't update shelf: ", err)
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Get Param and Check
	str := c.Params("id")
	gid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't update shelf: ", err)
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update shelf",
		})
	}
	// Get User And Check Ownership
	user := c.Locals("user").(models.User)
	shelfT, err := controller.GetBookShelf(uint(gid))
	if err != nil {
		utils.Logger.Warn("Couldn't update shelf: ", err)
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't update shelf",
		})
	}
	if shelfT.UserID != user.ID {
		utils.Logger.Warn("Couldn't update shelf: unauthorized access")
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't update shelf",
		})
	}

	errValidating := ValidateUpdateShelf(reqForm)
	if errValidating != nil {
		utils.Logger.Warn("Couldn't update shelf: ", errValidating)
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update and err
	errUpdate := controller.UpdateBookshelf(uint(gid), reqForm.Name, reqForm.Detail)
	if errUpdate != nil {
		utils.Logger.Warn("Couldn't update shelf: ", errUpdate)
		utils.Logger.Trace("Ending UpdateBookshelfHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't update shelf",
		})
	}
	utils.Logger.Trace("Shelf updated. ID: ", gid)
	utils.Logger.Trace("Ending UpdateBookshelfHandler")
	return c.SendStatus(fiber.StatusOK)
}

type changeSlot struct {
	ID    uint
	State bool
}

type changeState struct {
	SlotsID []changeSlot
}

func UpdateShelfSlotStateHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting UpdateShelfSlotStateHandler")
	user := c.Locals("user").(models.User)
	str := c.Params("id")
	ssid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't update shelfslot: ", err)
		utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update shelfslot",
		})
	}
	shelfT, err := controller.GetBookShelf(uint(ssid))
	if err != nil {
		utils.Logger.Warn("Couldn't update shelf: ", err)
		utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't update shelf",
		})
	}
	if shelfT.UserID != user.ID {
		utils.Logger.Warn("Couldn't update shelf: unauthorized access")
		utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't update shelf",
		})
	}
	var reqForm changeState
	if err := c.BodyParser(&reqForm); err != nil {
		utils.Logger.Warn("Couldn't update shelfslot: ", err)
		utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	for _, i := range reqForm.SlotsID {
		slot, err := controller.GetShelfslot(i.ID)
		if err != nil {
			utils.Logger.Warn("Couldn't update shelfslot: ", err)
			utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not update shelfslot",
			})
		}
		if slot.BookshelfID != uint(ssid) {
			utils.Logger.Warn("Couldn't update shelfslot: unauthorized access")
			utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not update shelfslot",
			})
		}
		controller.UpdateSlotState(i.ID, i.State)
	}
	utils.Logger.Trace("Shelfslot state updated. ID: ", ssid)
	utils.Logger.Trace("Ending UpdateShelfSlotStateHandler")
	return c.SendStatus(fiber.StatusOK)
}

func GetShelfslotByIdHandler(c *fiber.Ctx) error {
	utils.Logger.Trace("Starting GetShelfslotByIdHandler")
	str := c.Params("id")
	ssid, err := strconv.Atoi(str)
	if err != nil {
		utils.Logger.Warn("Couldn't get shelfslot: ", err)
		utils.Logger.Trace("Ending GetShelfslotByIdHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get shelfslot",
		})
	}
	ss, err := controller.GetShelfslot(uint(ssid))
	if err != nil {
		utils.Logger.Warn("Couldn't get shelfslot: ", err)
		utils.Logger.Trace("Ending GetShelfslotByIdHandler")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not get shelfslot",
		})
	}
	utils.Logger.Trace("Ending GetShelfslotByIdHandler")
	return c.Status(fiber.StatusOK).JSON(ss)
}
