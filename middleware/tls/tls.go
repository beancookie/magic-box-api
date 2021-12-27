package tls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/BeanCookie/magic-box-api/pkg/setting"

)

func TLSHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
			endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort)

			secureMiddleware := secure.New(secure.Options{
					SSLRedirect: true,
					SSLHost:     endPoint,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
					return
			}

			c.Next()
    }
}