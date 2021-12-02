package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BeanCookie/magic-box-api/pkg/app"
	"github.com/BeanCookie/magic-box-api/pkg/e"
	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/pkg/util"
	"github.com/BeanCookie/magic-box-api/service/juejin_service"
)

func GetJuejinArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	// juejin_service.ParseArticles("https://api.juejin.cn/recommend_api/v1/article/recommend_all_feed")
	req := juejin_service.ArticleRequest{app.PaginationRequest{Page: util.GetPage(c), Size: setting.AppSetting.PageSize}}
	list, total, _ := juejin_service.GetArticles(req)
	data := make(map[string]interface{})
	data["list"] = list
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
