package juejin_service

import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/BeanCookie/magic-box-api/pkg/app"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const ERR_NO = "err_no"
const DATA = "data"
const ITEM_INFO = "item_info"
const ARTICLE_INFO = "article_info"
const AUTHOR_USER_INFO = "author_user_info"
const CATEGORY = "category"
const TAGS = "tags"

const ARTICLE_ID = "article_id"

type ArticleRequest struct {
	app.PaginationRequest
}

func (req ArticleRequest) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}

func GetArticles(req ArticleRequest) (data []*models.Article, total uint, err error) {
	total, err = models.GetArticleTotal(req.getMaps())
	if err != nil {
		return nil, 0, err
	}
	log.Info().Msgf("%d", total)

	data, err = models.GetArticles(req.Page, req.Size, req.getMaps())
	log.Info().Msgf("%v", data)

	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func ParseArticles(url string) {

	client := resty.New()

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"limit": 0}`).
		Post(url)

	if err != nil {
		log.Error().Msgf("%v", err)
	}
	resJson := gjson.Parse(string(res.Body()))
	if resJson.Get(ERR_NO).Int() == 0 {
		resJson.Get(DATA).ForEach(func(index, value gjson.Result) bool {
			item := value.Get(ITEM_INFO)
			article := item.Get(ARTICLE_INFO)
			user := item.Get(AUTHOR_USER_INFO)
			category := item.Get(CATEGORY)
			existed, _ := models.ExistArticleById(article.Get(ARTICLE_ID).String())
			if !existed {
				if article.Value() != nil {
					models.AddArticle(article.Value().(map[string]interface{}),
						user.Value().(map[string]interface{}),
						category.Value().(map[string]interface{}),
					)
				}
			}
			return true
		})
	}
}
