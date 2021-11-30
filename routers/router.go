package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/BeanCookie/magic-box-api/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/BeanCookie/magic-box-api/middleware/jwt"
	"github.com/BeanCookie/magic-box-api/routers/api"
	"github.com/BeanCookie/magic-box-api/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
	}

	return r
}
