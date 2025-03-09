package controller

import (
	"myvinyl/models"
	"myvinyl/utils"
	"time"
)

type ForVinyl struct {
	ID           uint
	GenreID      uint
	ShelfslotID  uint
	Shelfslot    models.Shelfslot
	Name         string
	Artist       string
	Detail       string
	Price        uint
	ImageURL     string
	Format       uint
	Sleeve       uint
	Media        uint
	ReleasedDate string
	CreatedAt    time.Time
}

func CreateVinyl(userID uint, genreID uint, shelfslotID uint, name string, artist string, detail string, price uint, imageUrl string, format uint, sleeve uint, media uint, releaseDate string) error {
	vinyl := models.Vinyl{
		UserID:       userID,
		GenreID:      genreID,
		ShelfslotID:  &shelfslotID,
		Name:         name,
		Artist:       artist,
		Detail:       detail,
		Price:        price,
		ImageURL:     imageUrl,
		Format:       format,
		Sleeve:       sleeve,
		Media:        media,
		ReleasedDate: releaseDate,
	}
	err := utils.DB.Create(&vinyl)
	if err != nil {
		return err.Error
	}
	utils.Logger.Trace("Vinyl created. ID: ", vinyl.ID)
	return nil
}

func GetAllVinylsByUserID(userId uint) ([]ForVinyl, error) {
	var temp []ForVinyl
	err := utils.DB.Model(&models.Vinyl{}).
		Where("user_id = ?", userId).
		Preload("Shelfslot").           // Preload Shelfslot
		Preload("Shelfslot.Bookshelf"). // Preload Bookshelf from Shelfslot
		Select("id", "genre_id", "shelfslot_id", "name", "artist", "detail", "price", "image_url", "format", "sleeve", "media", "released_date", "created_at").
		Find(&temp).Error // Corrected to use Error on the result of Find

	if err != nil {
		return temp, err
	}
	return temp, nil
}

func GetVinylByID(id uint) (models.Vinyl, error) {
	var temp models.Vinyl
	err := utils.DB.Where("ID = ?", id).First(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}

func GetVinylByIDForHandler(id uint) (ForVinyl, error) {
	var temp ForVinyl
	err := utils.DB.Model(&models.Vinyl{}).Where("ID = ?", id).First(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}

func UpdateVinyl(id uint, genreId uint, shelfslotId uint, name string, artist string, detail string, price uint, imageUrl string, format uint, sleeve uint, media uint, releasedDate string) error {
	vinyl, err := GetVinylByID(id)
	if err != nil {
		return err
	}
	updateFields := map[string]interface{}{}
	updateFields["genre_id"] = genreId
	updateFields["shelfslot_id"] = shelfslotId
	if name != "" {
		updateFields["name"] = name
	}
	if artist != "" {
		updateFields["artist"] = artist
	}
	if detail != "" {
		updateFields["detail"] = detail
	}
	updateFields["price"] = price
	if imageUrl != "" {
		updateFields["image_url"] = imageUrl
	}
	updateFields["format"] = format
	updateFields["sleeve"] = sleeve
	updateFields["media"] = media
	if releasedDate != "" {
		updateFields["released_date"] = releasedDate
	}
	if err := utils.DB.Model(&vinyl).Updates(updateFields).Error; err != nil {
		return err
	}

	return nil
}

func UpdateVinylSlot(id uint, shelfslotId uint) error {
	vinyl, err := GetVinylByID(id)
	if err != nil {
		return err
	}
	updateFields := map[string]interface{}{}
	updateFields["shelfslot_id"] = shelfslotId
	if err := utils.DB.Model(&vinyl).Updates(updateFields).Error; err != nil {
		return err
	}

	return nil
}

func DeleteVinyl(id uint) error {
	vinyl, err := GetVinylByID(id)
	if err != nil {
		return err
	}
	utils.DB.Delete(&vinyl)
	utils.Logger.Trace("Vinyl deleted. ID: ", id)
	return nil
}

func DeleteByUserId(id uint) error {
	var temp []models.Vinyl
	err := utils.DB.Model(&models.Vinyl{}).Where("user_id = ?", id).Find(&temp)
	if err.Error != nil {
		return err.Error
	}
	for _, i := range temp {
		DeleteVinyl(i.ID)
	}
	return nil
}

func GetVinylsByShelfSlot(ssId uint) ([]ForVinyl, error) {
	var temp []ForVinyl
	err := utils.DB.Model(&models.Vinyl{}).Where("shelfslot_id = ?", ssId).Find(&temp)
	if err.Error != nil {
		return temp, err.Error
	}
	return temp, nil
}
