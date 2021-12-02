package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type Article struct {
	Model

	ID           string `gorm:"primary_key" json:"id"`
	Title        string `json:"title"`
	CoverImage   string `json:"cover_image"`
	BriefContent string `json:"brief_content"`
	CollectCount int    `json:"collect_count"`
	CommentCount int    `json:"comment_count"`

	UserID string `json:"user_id"`
	User   User   `json:"user"`

	CategoryID string   `json:"category_id"`
	Category   Category `json:"category"`
}

func AddArticle(articleData map[string]interface{},
	userData map[string]interface{},
	categoryData map[string]interface{},
) error {
	article := Article{
		ID:           articleData["article_id"].(string),
		Title:        articleData["title"].(string),
		CoverImage:   articleData["cover_image"].(string),
		BriefContent: articleData["brief_content"].(string),
		CollectCount: int(articleData["collect_count"].(float64)),
		CommentCount: int(articleData["comment_count"].(float64)),
		UserID:       userData["user_id"].(string),
		CategoryID:   categoryData["category_id"].(string),
	}
	if err := db.Create(&article).Error; err != nil {
		log.Info().Msgf("%v", article)
		return err
	}
	return nil
}

func ExistArticleById(id string) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if article.ID != "" {
		return true, nil
	}
	return false, nil
}

func GetArticles(page int, size int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Where(maps).Offset(page).Limit(size).Find(&articles).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticleTotal(maps interface{}) (uint, error) {
	var total uint
	if err := db.Model(&Article{}).Where(maps).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
