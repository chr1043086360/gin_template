package service

import "github.com/gin-gonic/gin"

func ArticleList(c *gin.Context) {
	message := c.PostForm("message")
	// nick :=
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
	})
}
