package user

import (
	"errors"
	"myvinyl/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Pass2Hash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "0"
	}
	return string(hashed)
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserBySession(c *fiber.Ctx) (string, error) {
	sess, err := utils.SessionManager.Get(c)
	if err != nil {
		utils.Logger.Error("Error retrieving session: ", err)
		return "", err
	}

	// Get the "username" from the session
	username := sess.Get("username")
	if username == nil {
		utils.Logger.Warn("Username not found in session")
		return "", errors.New("username not found in session")
	}

	usernameStr, ok := username.(string)
	if !ok {
		utils.Logger.Warn("Invalid type for username in session")
		return "", errors.New("username is not a string")
	}

	return usernameStr, nil
}
