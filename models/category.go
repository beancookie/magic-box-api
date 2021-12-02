package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	Model

	ID   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Url  int    `json:"url"`
}

func ExistCategoryByID(id string) (bool, error) {
	var category Category
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if category.ID != "" {
		return true, nil
	}

	return false, nil
}
