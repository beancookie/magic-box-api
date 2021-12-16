package hot_new_service

import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/BeanCookie/magic-box-api/pkg/app"
	"github.com/rs/zerolog/log"

)

type HotNewRequest struct {
	app.PaginationRequest
	Platform string
}

func (req HotNewRequest) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	maps["platform"] = req.Platform
	return maps
}

func GetHotNews(req HotNewRequest) (data []*models.HotNew, total uint, err error) {
	total, err = models.GetArticleTotal(req.getMaps())
	if err != nil {
		return nil, 0, err
	}
	log.Info().Msgf("%v", req)
	data, err = models.GetHotNews(req.Page, req.Size, req.getMaps())

	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
