package Auth

import (
	"app/Application"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	r, auth := Application.AuthRequest(c)
	if !auth {
		return
	}
	r.Ok(r.User.Transform())
}
