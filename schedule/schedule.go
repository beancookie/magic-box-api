package schedule

import (
	"fmt"
	"time"

	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/service/csdn_service"
	"github.com/BeanCookie/magic-box-api/service/juejin_service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type ArticleJob struct{}

func (j ArticleJob) Run() {
	host := setting.ThirdPartysSetting.JueJin
	juejin_service.ParseArticles(host + "/recommend_api/v1/article/recommend_all_feed")
	csdn_service.ParseArticles(host + "/phoenix/web/blog/hot-rank?page=0&pageSize=25")

	fmt.Println("run article job")
	log.Info().Msgf("run article job %v", time.Now())
}

func Setup() {
	c := cron.New(cron.WithSeconds())
	c.AddJob(setting.AppSetting.ScheduleCron, ArticleJob{})
	c.Start()
}
