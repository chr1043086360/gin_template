package service

import "github.com/gin-gonic/gin"

func Info(c *gin.Context){
	user, _ := c.Get("user")

	c.JSON(200, gin.H{"user": user})
}
