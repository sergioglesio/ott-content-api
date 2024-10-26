package models

type Content struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"size:500"`
	Cast        string `gorm:"size:255"`
	ReleaseYear int
}
