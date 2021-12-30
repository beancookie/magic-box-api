package article_service

import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/BeanCookie/magic-box-api/pkg/app"
	"github.com/BeanCookie/magic-box-api/pkg/util"
)

type ArticleRequest struct {
	app.PaginationRequest
	Platform string
}

func (req ArticleRequest) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	maps["platform"] = req.Platform
	return maps
}

func GetArticles(req ArticleRequest) (data []*models.Article, total uint, err error) {
	total, err = models.GetArticleTotal(req.getMaps())
	if err != nil {
		return nil, 0, err
	}

	data, err = models.GetArticles(req.Page, req.Size, req.getMaps())

	data = models.MapArticle(data, func(index int, item *models.Article) *models.Article {
		item.ParstTime = util.PastTime(item.CreatedOn)
		return item
	})

	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
