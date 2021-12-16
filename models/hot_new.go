package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)


const WEIBO = "weibo"

type HotNew struct {
	ID       	string  `gorm:"primary_key" json:"id"`
	OnboardTime int 	`json:"onboard_time"`
	RawHot 		int 	`json:"raw_hot"`
	Category 	string 	`json:"category"`
	KeyWord 	string 	`json:"key_word"`
	Content 	string 	`json:"content"`
}

func ExistHotNewByIdAndPlatform(id string, platform string) (bool, error) {
	var hotNew HotNew
	err := db.Select("id").Where("id = ? AND platform = ? AND deleted_on = ? ", id, platform, 0).First(&hotNew).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if hotNew.ID != "" {
		return true, nil
	}
	return false, nil
}

func AddWeiboHotNew(data gjson.Result) error {
	log.Info().Msgf("add hotNew %s", data.Get("word").String())
	hotNew := HotNew{
		ID: data.Get("mid").String(),
		OnboardTime: int(data.Get("onboard_time").Int()),
		RawHot: int(data.Get("raw_hot").Int()),
		Category: data.Get("category").String(),
		KeyWord: data.Get("word").String(),
		Content: data.Get("mblog").Get("text").String(),
	}
	if err := db.Create(&hotNew).Error; err != nil {
		log.Info().Msgf("%v", hotNew)
		return err
	}
	return nil
}

