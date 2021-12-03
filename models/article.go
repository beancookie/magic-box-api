package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const JUEJIN = "juejin"
const CSDN = "csdn"

type Article struct {
	Model

	ID           string `gorm:"primary_key" json:"id"`
	Title        string `json:"title"`
	CoverImage   string `json:"cover_image"`
	BriefContent string `json:"brief_content"`
	CollectCount int    `json:"collect_count"`
	CommentCount int    `json:"comment_count"`

	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`

	// User   User   `json:"user"`

	// CategoryID string   `json:"category_id"`
	// Category   Category `json:"category"`
	Platform string `json:"Platform"`
}

func AddJuejinArticle(data gjson.Result) error {
	log.Info().Msgf("add article %s", data.Get("title").String())
	article := Article{
		ID:           data.Get("article_id").String(),
		Title:        data.Get("title").String(),
		CoverImage:   data.Get("cover_image").String(),
		BriefContent: data.Get("brief_content").String(),
		CollectCount: int(data.Get("collect_count").Int()),
		CommentCount: int(data.Get("comment_count").Int()),
		Platform:     JUEJIN,
	}
	if err := db.Create(&article).Error; err != nil {
		log.Info().Msgf("%v", article)
		return err
	}
	return nil
}

func AddCsdnArticle(data gjson.Result) error {
	log.Info().Msgf("add article %s", data.Get("articleTitle").String())
	picList := data.Get("picList").Array()
	coverImage := ""
	if len(picList) > 0 {
		coverImage = picList[0].String()
	}
	splitUrl := strings.Split(data.Get("articleDetailUrl").String(), "/")
	id := splitUrl[len(splitUrl)-1]
	article := Article{
		ID:         id,
		Title:      data.Get("articleTitle").String(),
		UserName:   data.Get("userName").String(),
		CoverImage: coverImage,
		Platform:   CSDN,
	}
	if err := db.Create(&article).Error; err != nil {
		log.Info().Msgf("%v", article)
		return err
	}
	return nil
}

func ExistArticleByIdAndPlatform(id string, platform string) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND platform = ? AND deleted_on = ? ", id, platform, 0).First(&article).Error
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
