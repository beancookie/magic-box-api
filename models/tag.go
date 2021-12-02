package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	ID               string `gorm:"primary_key" json:"string"`
	Name             string `json:"name"`
	ConcernUserCount int    `json:"concern_user_count"`
	PostArticleCount int    `json:"post_article_count"`
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID != "" {
		return true, nil
	}

	return false, nil
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID != "" {
		return true, nil
	}

	return false, nil
}

// AddTag Add a Tag
func AddTag(name string, concernUserCount int, postArticleCount int) error {
	tag := Tag{
		Name:             name,
		ConcernUserCount: concernUserCount,
		PostArticleCount: postArticleCount,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
