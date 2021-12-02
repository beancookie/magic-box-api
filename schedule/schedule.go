package schedule

import (
	"fmt"
	"time"

	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/service/juejin_service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type JueJinJob struct{}

func (j JueJinJob) Run() {
	host := setting.ThirdPartysSetting.JueJin
	juejin_service.ParseArticles(host + "/recommend_api/v1/article/recommend_all_feed")
	fmt.Println("run juejin job")
	log.Info().Msgf("run juejin job %v", time.Now())
}

func Setup() {
	c := cron.New(cron.WithSeconds())
	c.AddJob("*/60 * * * * *", JueJinJob{})
	c.Start()
}
