package controller

import (
	"myvinyl/models"
	"myvinyl/modules/user"
	"myvinyl/utils"
)

// Permission: 0
func CreateUser(username string, password string) bool {
	if IsIdDuplicated(username) {
		return false
	}
	hashedPassword := user.Pass2Hash(password)
	user := models.User{Username: username, Password: hashedPassword}
	utils.DB.Create(&user)
	utils.Logger.Trace("User created. username: ", username)
	return true
}

func GetUserById(username string) (models.User, error) {
	var temp models.User
	err := utils.DB.Where("Username = ?", username).First(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}

func UpdateUserById(username string, cUsername string, password string) error {
	var temp models.User
	if err := utils.DB.First(&temp, username).Error; err != nil {
		return err
	}
	updateFields := map[string]interface{}{}
	if cUsername != "" {
		updateFields["Username"] = cUsername
	}
	if password != "" {
		updateFields["Password"] = user.Pass2Hash(password)
	}
	if err := utils.DB.Model(&temp).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(username string) error {
	var err error
	var temp models.User
	result := utils.DB.Where("Username = ?", username).First(&temp)
	if result.Error != nil {
		return err
	}
	dRes := utils.DB.Delete(&temp)
	if dRes.Error != nil {
		return err
	}
	utils.Logger.Trace("User deleted. Username: ", username)
	return nil
}

func GetUserVinylsAndShelvesLength(userID uint) (int, int, error) {
	utils.Logger.Trace("Starting GetUserVinylsAndShelvesLength")
	var vinyls []models.Vinyl
	var shelves []models.Bookshelf

	// Get user's vinyls
	if err := utils.DB.Where("user_id = ?", userID).Find(&vinyls).Error; err != nil {
		utils.Logger.Error("Error getting user's vinyls: ", err)
		utils.Logger.Trace("Ending GetUserVinylsAndShelvesLength")
		return 0, 0, err
	}

	// Get user's shelves
	if err := utils.DB.Where("user_id = ?", userID).Find(&shelves).Error; err != nil {
		utils.Logger.Error("Error getting user's shelves: ", err)
		utils.Logger.Trace("Ending GetUserVinylsAndShelvesLength")
		return 0, 0, err
	}

	utils.Logger.Trace("Ending GetUserVinylsAndShelvesLength")
	return len(vinyls), len(shelves), nil
}

// Check Functions
func IsIdDuplicated(username string) bool {
	var temp models.User
	result := utils.DB.Where("Username = ?", username).First(&temp)
	return result.Error == nil
}

func CompareHashAndPassword(username string, password string) bool {
	var temp models.User
	utils.DB.Where("Username = ?", username).First(&temp)
	return user.ComparePassword(password, temp.Password)
}
