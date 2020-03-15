package middleware

import (
	"project2019/jwt"
	"project2019/models"

	"project2019/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, serializer.Response{
				Code:  401,
				Data:  "",
				Msg:   "权限不足",
				Error: "",
			})
			// 将请求抛弃掉
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || token.Valid {
			c.JSON(401, serializer.Response{
				Code:  401,
				Data:  "",
				Msg:   "权限不足",
				Error: "",
			})
			c.Abort()
			return
		}

		// 验证通过获取userId
		userId := claims.UserId
		// 在数据库中查找
		user := models.User{}
		err2 := models.DB.Where("id", userId).First(&user).Error
		if err2 != nil {
			c.JSON(200, serializer.Response{
				Code:  50001,
				Data:  "",
				Msg:   "获取用户失败",
				Error: "",
			})
			c.Abort()
			return
		}
		// 将用户写入上下文
		c.Set("user", user)
		c.Next()
	}
}
