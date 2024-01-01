package Visitors

import (
	"app/Application"
	"github.com/gin-gonic/gin"
)



func Welcome(c *gin.Context) {
	Application.NewReuqest(c).Ok("welcome in test app")
}
