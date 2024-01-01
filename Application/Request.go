package Application

import (
	"app/Models"
	"database/sql"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type ShareResources interface {
	Share()
}

type Request struct {
	Context          *gin.Context
	DB               *gorm.DB
	Connection       *sql.DB
	User             *Models.User
	IsAuth           bool
	IsAdmin          bool
	Lang             string
	ValidateionError error
}

func (req *Request) Share() {}

func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDataBase(&request)
		setLang(&request)
		return &request
	}
}

func (req *Request) ValidateRequest(errors validation.Errors) *Request {
	req.ValidateionError = errors.Filter()
	return req
}

func (req *Request) Fails() bool {
	if req.ValidateionError != nil {
		req.BadRequest(req.ValidateionError)
		return true
	}
	return false
}

func setLang(req *Request) {
	lang := gotrans.DetectLanguage(req.Context.GetHeader("Accept-Language"))
	gotrans.SetDefaultLocale(lang)
	req.Lang = lang
}

func (req Request) Resonse(code int, body map[string]interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

func NewReuqest(c *gin.Context) *Request {
	reuqest := req()
	return reuqest(c)
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewReuqest(c).Auth()
}

func AuthRequest(c *gin.Context) (*Request, bool) {
	r := NewRequestWithAuth(c)
	if !r.IsAuth {
		r.NotAuth()
		return r, false
	}
	return r, true
}

func (req *Request) Auth() *Request {
	req.IsAuth = false
	req.IsAdmin = false
	authHeader := req.Context.GetHeader("Authorization")
	if authHeader != "" {
		req.DB.Where("token = ? ", authHeader).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
			if req.User.Group == "admin" {
				req.IsAdmin = true
			}
		}
	}
	return req
}
