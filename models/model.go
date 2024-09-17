package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	CreatedBy string
	UpdatedBy string
}
