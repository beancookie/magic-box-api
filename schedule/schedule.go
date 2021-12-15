package schedule

import (
	"time"

	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/service/juejin_service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type ArticleJob struct{}

func (j ArticleJob) Run() {
	juejin := setting.ThirdPartysSetting.JueJin
	// csdn := setting.ThirdPartysSetting.Csdn
	juejin_service.ParseArticles(juejin)
	// csdn_service.ParseArticles("https://blog.csdn.net/phoenix/web/blog/hot-rank?page=0&pageSize=25")

	log.Info().Msgf("run article job %v", time.Now())
}

func Setup() {
	c := cron.New(cron.WithSeconds())
	c.AddJob(setting.AppSetting.ScheduleCron, ArticleJob{})
	c.Start()
}
