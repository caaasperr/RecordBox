package controller

import (
	"myvinyl/models"
	"myvinyl/utils"
)

type ForBookshelves struct {
	ID      uint
	Name    string
	Detail  string
	Columns uint
	Rows    uint
}

func CreateBookshelf(userId uint, name string, detail string, columns uint, rows uint) (uint, error) {
	utils.Logger.Trace("Starting CreateBookshelf")
	bookshelf := models.Bookshelf{
		UserID:  userId,
		Name:    name,
		Detail:  detail,
		Columns: columns,
		Rows:    rows,
	}
	err := utils.DB.Create(&bookshelf).Error
	if err != nil {
		utils.Logger.Error("Error creating bookshelf: ", err)
		return bookshelf.ID, err
	}
	utils.Logger.Trace("Bookshelf created. Name: ", name)
	utils.Logger.Trace("Ending CreateBookshelf")
	return bookshelf.ID, nil
}

func GetBookshelves(userId uint) ([]ForBookshelves, error) {
	utils.Logger.Trace("Starting GetBookshelves")
	var temp []ForBookshelves
	err := utils.DB.Model(&models.Bookshelf{}).Select("ID", "Name", "Detail", "Columns", "Rows").Where("user_id = ?", userId).Find(&temp).Error
	if err != nil {
		utils.Logger.Error("Error getting bookshelves: ", err)
		return temp, err
	}
	utils.Logger.Trace("Bookshelves retrieved for user ID: ", userId)
	utils.Logger.Trace("Ending GetBookshelves")
	return temp, nil
}

func GetBookShelf(id uint) (models.Bookshelf, error) {
	utils.Logger.Trace("Starting GetBookShelf")
	var temp models.Bookshelf
	err := utils.DB.Where("id = ?", id).First(&temp).Error
	if err != nil {
		utils.Logger.Error("Error getting bookshelf: ", err)
		return temp, err
	}
	utils.Logger.Trace("Bookshelf retrieved. ID: ", id)
	utils.Logger.Trace("Ending GetBookShelf")
	return temp, nil
}

func UpdateBookshelf(id uint, name string, detail string) error {
	utils.Logger.Trace("Starting UpdateBookshelf")
	var temp models.Bookshelf
	if err := utils.DB.Where("id = ?", id).First(&temp).Error; err != nil {
		utils.Logger.Error("Error finding bookshelf: ", err)
		return err
	}
	updateFields := make(map[string]interface{})

	if name != "" {
		updateFields["name"] = name
	}
	if detail != "" {
		updateFields["detail"] = detail
	}

	if err := utils.DB.Model(&temp).Updates(updateFields).Error; err != nil {
		utils.Logger.Error("Error updating bookshelf: ", err)
		return err
	}
	utils.Logger.Trace("Bookshelf updated. ID: ", id)
	utils.Logger.Trace("Ending UpdateBookshelf")
	return nil
}
