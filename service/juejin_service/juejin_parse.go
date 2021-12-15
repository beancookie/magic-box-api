package juejin_service

import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const ERR_NO = "err_no"
const DATA = "data"
const ARTICLE_INFO = "article_info"
const ITEM_INFO = "item_info"
const AUTHOR_USER_INFO = "author_user_info"
const CATEGORY = "category"
const TAGS = "tags"

const ARTICLE_ID = "article_id"

func ParseArticles(url string) {
	log.Info().Msgf("Juejin ParseArticles")

	client := resty.New()

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"cate_id": "6809637772874219534", "limit": 200}`).
		Post(url)

	if err != nil {
		log.Error().Msgf("%v", err)
	}
	resJson := gjson.Parse(string(res.Body()))
	if resJson.Get(ERR_NO).Int() == 0 {
		resJson.Get(DATA).ForEach(func(index, value gjson.Result) bool {
			article := value.Get(ITEM_INFO).Get(ARTICLE_INFO)
			existed, _ := models.ExistArticleByIdAndPlatform(article.Get(ARTICLE_ID).String(), models.JUEJIN)
			if !existed {
				if article.Value() != nil {
					models.AddJuejinArticle(article)
				}
			}
			return true
		})
	}
}
