package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/BeanCookie/magic-box-api/middleware/tls"
	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/routers/api"
	v1 "github.com/BeanCookie/magic-box-api/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(tls.TLSHandler())

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/hot-news", v1.GetHotNews)
	}

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort)


  r.RunTLS(endPoint, fmt.Sprintf("%s/%s.pem", setting.AppSetting.TLSPath, setting.AppSetting.TLSName), fmt.Sprintf("%s/%s.key", setting.AppSetting.TLSPath, setting.AppSetting.TLSName))

	return r
}
