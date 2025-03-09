package models

type Bookshelf struct {
	ID      uint `gorm:"primaryKey;<-:create"`
	UserID  uint `gorm:"<-:create"`
	Name    string
	Detail  string
	Columns uint
	Rows    uint
}

type Shelfslot struct {
	ID          uint `gorm:"primaryKey;<-:create"`
	BookshelfID uint `gorm:"<-:create"`
	Name        string
	Detail      string
	Column      uint
	Row         uint
	Enabled     bool

	Bookshelf Bookshelf `gorm:"foreignKey:BookshelfID"`
}
