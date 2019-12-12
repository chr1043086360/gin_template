package service

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Param(c *gin.Context) {
	res := c.Param("name")
	// fmt.Print(res)
	c.String(http.StatusOK, res)
}
