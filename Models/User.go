package Models

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50)"`
	Email    string `json:"email" gorm:"type:varchar(50)"`
	Password string `json:"password" gorm:"type:varchar(50)"`
	Token    string `json:"token" gorm:"type:varchar(100)"`
	Group    string `json:"group" gorm:"type:varchar(20)"`
}

func (user User) Transform() map[string]interface{} {
	m := make(map[string]interface{})
	m["username"] = user.Username
	m["email"] = user.Email
	m["token"] = user.Token
	return m
}

func (users Users) Transform() []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, user := range users {
		m = append(m, user.Transform())
	}
	return m
}
