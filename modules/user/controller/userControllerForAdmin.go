package controller

import (
	"myvinyl/models"
	"myvinyl/utils"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := utils.DB.Select("ID", "Username", "CreatedAt", "IsAdmin").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
