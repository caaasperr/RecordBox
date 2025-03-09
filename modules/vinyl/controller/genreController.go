package controller

import (
	"myvinyl/models"
	"myvinyl/utils"
)

func CreateGenre(name string) error {
	genre := models.Genre{Name: name}
	err := utils.DB.Create(&genre)
	if err.Error != nil {
		return err.Error
	}
	utils.Logger.Trace("Genre created. Name: ", name)
	return nil
}

func GetVinylsByGenre(genreId uint, userId uint) ([]models.Vinyl, error) {
	var temp []models.Vinyl
	err := utils.DB.Where("genre_id = ? and user_id = ?", genreId, userId).Find(&temp)
	if err != nil {
		return temp, err.Error
	}
	return temp, nil
}

func GetGenreNameByID(id uint) (string, error) {
	genre, err := GetGenreByID(id)
	if err != nil {
		return genre.Name, err
	}
	return genre.Name, nil
}

func GetGenres() ([]models.Genre, error) {
	var temp []models.Genre
	err := utils.DB.Find(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}

func GetGenreByID(id uint) (models.Genre, error) {
	var temp models.Genre
	err := utils.DB.Where("ID = ?", id).First(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}

func UpdateGenre(id uint, updateData map[string]interface{}) error {
	genre, err := GetGenreByID(id)
	if err != nil {
		return err
	}
	if err := utils.DB.Model(&genre).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGenre(id uint) error {
	genre, err := GetGenreByID(id)
	if err != nil {
		return err
	}
	utils.DB.Delete(&genre)
	utils.Logger.Trace("Genre deleted. ID: ", id)
	return nil
}
