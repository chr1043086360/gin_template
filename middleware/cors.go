package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
// func Cors() gin.HandlerFunc {
// 	config := cors.DefaultConfig()
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
// 	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Access-Control-Allow-Origin"}
// 	if gin.Mode() == gin.ReleaseMode {
// 		// 生产环境需要配置跨域域名，否则403
// 		config.AllowOrigins = []string{"http://www.zhulin.com"}
// 	} else {
// 		// 测试环境下模糊匹配本地开头的请求
// 		config.AllowOriginFunc = func(origin string) bool {
// 			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
// 				return true
// 			}
// 			if regexp.MustCompile(`^http://localhost\.0\.0\.1:\d+$`).MatchString(origin) {
// 				return true
// 			}
// 			return false
// 		}
// 	}
// 	config.AllowCredentials = true
// 	return cors.New(config)
// }

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
