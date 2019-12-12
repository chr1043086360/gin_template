
/*
  Author ： CHR_崔贺然
  Time ： 2019.11.15
  Description ： 
*/

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type myForm struct {
// 	Colors []string `form:"colors[]"`
// }

func FormDemo(c *gin.Context) {
	// var fakeForm myForm
	// c.ShouldBind(&fakeForm)
	// c.JSON(200, gin.H{"color": fakeForm.Colors})
	
	c.HTML(http.StatusOK, "form.html", gin.H{})
}
