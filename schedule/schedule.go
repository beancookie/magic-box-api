package schedule

import (
	"time"

	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/service/csdn_service"
	"github.com/BeanCookie/magic-box-api/service/juejin_service"
	"github.com/BeanCookie/magic-box-api/service/weibo_service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type ArticleJob struct{}

func (j ArticleJob) Run() {
	juejin := setting.ThirdPartysSetting.JueJin
	csdn := setting.ThirdPartysSetting.Csdn
	juejin_service.ParseArticles(juejin)
	csdn_service.ParseArticles(csdn)

	log.Info().Msgf("run article job %v", time.Now())
}

func Setup() {
	c := cron.New(cron.WithSeconds())
	// ArticleJob{}.Run()
	c.AddJob(setting.AppSetting.ScheduleCron, ArticleJob{})
	c.Start()
	weibo_service.ParseRealtimehot()
}
