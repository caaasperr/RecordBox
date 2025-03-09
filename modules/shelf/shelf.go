package shelf

import (
	"myvinyl/modules/shelf/controller"
	"myvinyl/utils"
)

func CreateShelf(userId uint, name string, detail string, cols uint, rows uint) error {
	id, err := controller.CreateBookshelf(userId, name, detail, cols, rows)
	if err != nil {
		return err
	}
	for i := 0; i < int(cols); i++ {
		for ii := 0; ii < int(rows); ii++ {
			controller.CreateShelfslot(id, "", detail, uint(i), uint(ii))
		}
	}
	return nil
}

func DeleteShelfByShelfId(id uint) error {
	slots, err := controller.GetShelfslotByShelfId(id)
	if err != nil {
		return err
	}
	shelf, err := controller.GetBookShelf(id)
	if err != nil {
		return err
	}
	for _, i := range slots {
		utils.DB.Delete(i)
	}
	utils.DB.Delete(shelf)
	return nil
}

func DeleteShelvesByUserId(id uint) error {
	bookshelfs, err := controller.GetBookshelves(id)
	if err != nil {
		return err
	}
	for _, i := range bookshelfs {
		DeleteShelfByShelfId(i.ID)
	}
	return nil
}
