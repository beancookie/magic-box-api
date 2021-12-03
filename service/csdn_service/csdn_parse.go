package csdn_service

import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const CODE = "code"
const DATA = "data"

func ParseArticles(url string) {
	log.Info().Msgf("csdn ParseArticles")
	client := resty.New()

	res, err := client.R().
		Get(url)

	if err != nil {
		log.Error().Msgf("%v", err)
	}
	resJson := gjson.Parse(string(res.Body()))
	if resJson.Get(CODE).Int() == 200 {
		resJson.Get(DATA).ForEach(func(index, article gjson.Result) bool {
			// existed, _ := models.ExistArticleByIdAndPlatform(article.Get("").String(), models.CSDN)
			// if !existed {
			// 	if article.Value() != nil {
			models.AddCsdnArticle(article)
			// 	}
			// }
			return true
		})
	}
}
