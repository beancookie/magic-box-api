package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	ID               string `gorm:"primary_key" json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	JobTitle         string `json:"job_title"`
	GotDiggCount     int    `json:"got_digg_count"`
	GotViewCount     int    `json:"got_view_count"`
	PostArticleCount int    `json:"post_article_count"`
	FolloweeCount    int    `json:"followee_count"`
	FollowerCount    int    `json:"follower_count"`
}

func ExistUserByID(id string) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID != "" {
		return true, nil
	}

	return false, nil
}
