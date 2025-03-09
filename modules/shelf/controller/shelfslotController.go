package controller

import (
	"myvinyl/models"
	"myvinyl/utils"
)

func CreateShelfslot(bookshelfID uint, name string, detail string, column uint, row uint) error {
	utils.Logger.Trace("Starting CreateShelfslot")
	shelfslot := models.Shelfslot{
		BookshelfID: bookshelfID,
		Name:        name,
		Detail:      detail,
		Column:      column,
		Row:         row,
		Enabled:     true,
	}
	err := utils.DB.Create(&shelfslot).Error
	if err != nil {
		utils.Logger.Error("Error creating shelfslot: ", err)
		return err
	}
	utils.Logger.Trace("Shelfslot created. ParentID: ", bookshelfID)
	utils.Logger.Trace("Ending CreateShelfslot")
	return nil
}

func UpdateSlotState(id uint, state bool) error {
	utils.Logger.Trace("Starting UpdateSlotState")
	var slot models.Shelfslot
	if err := utils.DB.Where("id = ?", id).First(&slot).Error; err != nil {
		utils.Logger.Error("Error finding shelfslot: ", err)
		return err
	}
	if err := utils.DB.Model(&slot).Update("Enabled", state).Error; err != nil {
		utils.Logger.Error("Error updating shelfslot state: ", err)
		return err
	}
	utils.Logger.Trace("Shelfslot state updated. ID: ", id)
	utils.Logger.Trace("Ending UpdateSlotState")
	return nil
}

//NOT USING THIS FUNCTION
/*func UpdateSlots(bId uint, cols uint, rows uint) error {
	var temp models.Bookshelf
	err := utils.DB.Where("id = ?", bId).First(&temp).Error
	if err != nil {
		return err
	}
	if temp.Columns == cols && temp.Rows == rows {
		return nil
	}
	if cols > temp.Columns {
		for i := temp.Columns; i < cols; i++ {
			for ii := 0; ii < int(rows); ii++ {
				if err := CreateShelfslot(bId, "", "", i, uint(ii)); err != nil {
					return err
				}
			}
		}
	}
	if cols < temp.Columns {
		for i := cols; i < temp.Columns; i++ {
			var dtemp []models.Shelfslot
			err := utils.DB.Where("bookshelf_id = ? AND `column` = ?", bId, i).Find(&dtemp).Error
			if err != nil {
				return err
			}
			for _, slot := range dtemp {
				utils.DB.Delete(&slot)
			}
		}
	}
	if rows < temp.Rows {
		for i := uint(0); i < temp.Columns; i++ {
			for j := rows; j < temp.Rows; j++ {
				var slot models.Shelfslot
				err := utils.DB.Where("bookshelf_id = ? AND `column` = ? AND `row` = ?", bId, i, j).First(&slot).Error
				if err == nil {
					utils.DB.Delete(&slot)
				}
			}
		}
	}
	utils.Logger.Trace("Ending UpdateSlots")
	return nil
}*/

func GetShelfslotByShelfId(id uint) ([]models.Shelfslot, error) {
	utils.Logger.Trace("Starting GetShelfslotByShelfId")
	var temp []models.Shelfslot
	err := utils.DB.Where("bookshelf_id = ?", id).Find(&temp)
	if err != nil {
		utils.Logger.Error("Error getting shelfslots by shelf ID: ", err)
		return temp, err.Error
	}
	utils.Logger.Trace("Shelfslots retrieved for bookshelf ID: ", id)
	utils.Logger.Trace("Ending GetShelfslotByShelfId")
	return temp, nil
}

func GetShelfslot(id uint) (models.Shelfslot, error) {
	utils.Logger.Trace("Starting GetShelfslot")
	var temp models.Shelfslot
	err := utils.DB.Where("id = ?", id).Find(&temp)
	if err != nil {
		utils.Logger.Error("Error getting shelfslot: ", err)
		return temp, err.Error
	}
	utils.Logger.Trace("Shelfslot retrieved. ID: ", id)
	utils.Logger.Trace("Ending GetShelfslot")
	return temp, nil
}
