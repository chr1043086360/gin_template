package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project2019/qiniu"
)

func QiniuTest(c *gin.Context) {
	qiniu.UpFile()
	c.String(http.StatusOK, "success! ", qiniu.RET, qiniu.HASH)
}
