package Visitors

import (
	"app/Models"
	"app/Validations"
	validation "github.com/go-ozzo/ozzo-validation"
)

func RegisterValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"name":     validation.Validate(user.Username, Validations.RequiredRule(), Validations.MinMaxRule()),
		"email":    validation.Validate(user.Email, Validations.RequiredRule(), Validations.IsEmailRule(), Validations.MinMaxRule()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
	}
}

func LoginValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"email":    validation.Validate(user.Email, Validations.RequiredRule(), Validations.IsEmailRule(), Validations.MinMaxRule()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
	}
}

