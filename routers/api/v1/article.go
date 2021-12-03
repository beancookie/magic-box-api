package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/BeanCookie/magic-box-api/pkg/app"
	"github.com/BeanCookie/magic-box-api/pkg/e"
	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/pkg/util"
	"github.com/BeanCookie/magic-box-api/service/article_service"
)

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	platform := ""
	if arg := c.Query("platform"); arg != "" {
		platform = c.Query("platform")
	}
	valid.Required(platform, "platform")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	req := article_service.ArticleRequest{
		PaginationRequest: app.PaginationRequest{Page: util.GetPage(c), Size: setting.AppSetting.PageSize},
		Platform:          platform,
	}
	list, total, _ := article_service.GetArticles(req)
	data := make(map[string]interface{})
	data["list"] = list
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
