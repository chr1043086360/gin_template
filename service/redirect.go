package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
