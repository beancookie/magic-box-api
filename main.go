package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BeanCookie/magic-box-api/models"
	"github.com/BeanCookie/magic-box-api/pkg/gredis"
	"github.com/BeanCookie/magic-box-api/pkg/setting"
	"github.com/BeanCookie/magic-box-api/pkg/util"
	"github.com/BeanCookie/magic-box-api/routers"
	"github.com/BeanCookie/magic-box-api/schedule"
)

func init() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	util.Setup()
	schedule.Setup()
	// csdn_service.ParseArticles("https://blog.csdn.net/phoenix/web/blog/hot-rank?page=0&pageSize=25")
}

// @title Golang Gin API
// @version 1.0
// @description An example of gins
// @termsOfService https://github.com/BeanCookie/magic-box-api
// @license.name MIT
// @license.url https://github.com/BeanCookie/magic-box-api/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
