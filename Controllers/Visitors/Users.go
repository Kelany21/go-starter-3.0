package Visitors

import (
	"app/Application"
	"app/Models"
	"app/Validations/Visitors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	r := Application.NewReuqest(c)
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	if r.ValidateRequest(Visitors.RegisterValidation(user)).Fails() {
		return
	}
	user.Token = user.Email
	user.Group = "user"
	r.DB.Create(&user)
	r.Created(user.Transform())
}

func Login(c *gin.Context) {
	r := Application.NewReuqest(c)
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	if r.ValidateRequest(Visitors.LoginValidation(user)).Fails() {
		return
	}
	r.DB.Where("email = ? ", user.Email).Where("password = ? ", user.Password).First(&user)
	if user.ID == 0 {
		r.UserNotFound()
		return
	}
	r.Ok(user.Transform())
}
