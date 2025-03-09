package handler

import (
	"myvinyl/models"
	shelfCont "myvinyl/modules/shelf"
	"myvinyl/modules/user"
	"myvinyl/modules/user/controller"
	vinylCont "myvinyl/modules/vinyl/controller"
	"myvinyl/utils"

	"github.com/gofiber/fiber/v2"
)

// Sign up handler
// 입력값 검증 알고리즘 추가 예정
func SignUpHandler(c *fiber.Ctx) error {
	var user CreateUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	errValidating := user.Validate()
	if errValidating != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := controller.CreateUser(user.Username, user.Password)
	if !err {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username is duplicated",
		})
	} else {
		return c.SendStatus(fiber.StatusCreated)
	}
}

// Log in handler
func LogInHandler(c *fiber.Ctx) error {
	var user LoginForm
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errValidating := user.Validate()
	if errValidating != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	if !controller.CompareHashAndPassword(user.Username, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect username or password",
		})
	}

	sess, err := utils.SessionManager.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create session",
		})
	}
	sessID := sess.ID()

	sess.Set("username", user.Username)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	// 성공 응답 반환
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"session_id": sessID,
	})
}

// LogOutHandler
func LogOutHandler(c *fiber.Ctx) error {
	sess, err := utils.SessionManager.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving session",
		})
	}
	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error destroying session",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

// Get Profile
func GetProfileHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	vinylsLength, shelvesLength, err := controller.GetUserVinylsAndShelvesLength(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Getting user's vinyls and shelves failed.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Username":  user.Username,
		"Vinyls":    vinylsLength,
		"Shelves":   shelvesLength,
		"CreatedAt": user.CreatedAt,
	})
}

// Update Profile
func UpdateProfileHandler(c *fiber.Ctx) error {
	var updateData UpdateUserRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	errValidating := updateData.Validate()
	if errValidating != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	user := c.Locals("user").(models.User)
	if controller.IsIdDuplicated(updateData.Username) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username is duplicated",
		})
	}
	controller.UpdateUserById(user.Username, updateData.Username, updateData.Password)
	return c.SendStatus(fiber.StatusOK)
}

// Delete User
func DeleteUserHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	errVinyl := vinylCont.DeleteByUserId(user.ID)
	if errVinyl != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Deleting user failed.",
		})
	}
	errShelf := shelfCont.DeleteShelvesByUserId(user.ID)
	if errShelf != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Deleting user failed.",
		})
	}
	err := controller.DeleteUserById(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Deleting user failed.",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func SessionCheckHandelr(c *fiber.Ctx) error {
	username, err := user.GetUserBySession(c)
	if err != nil {
		utils.Logger.Warn("Get Session Error: Can't get user, error")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized.",
		})
	}
	if username == "" {
		utils.Logger.Info("Get Session Error: Username is null")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Session valid.",
	})
}
