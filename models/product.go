package models

type Product struct {
	Model
	Name   string
	Price  uint32
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
}
