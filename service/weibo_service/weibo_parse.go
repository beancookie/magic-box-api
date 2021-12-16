package weibo_service


import (
	"github.com/BeanCookie/magic-box-api/models"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"	
)

const OK = "ok"
const DATA = "data"
const BAND_LIST = "band_list"
const M_ID = "mid"

func ParseRealtimehot() {
	client := resty.New()

	res, err := client.R().
		Get("https://weibo.com/ajax/statuses/hot_band")

	if err != nil {
		log.Error().Msgf("%v", err)
	}
	resJson := gjson.Parse(string(res.Body()))

	if resJson.Get(OK).Int() == 1 {
		resJson.Get(DATA).Get(BAND_LIST).ForEach(func(index, value gjson.Result) bool {
			existed, _ := models.ExistHotNewByIdAndPlatform(value.Get(M_ID).String(), models.WEIBO)
			if !existed && value.Get(M_ID).Value() != nil {
				models.AddWeiboHotNew(value)
			}
			return true
		})
	}
}