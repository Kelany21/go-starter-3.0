package Validations

import (
	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func RequiredRule() validation.Rule {
	return validation.Required.Error(gotrans.T("required"))
}

func MinMaxRule() validation.Rule {
	return validation.Length(2, 50).Error(gotrans.T("min_max"))
}

func IsEmailRule() validation.Rule {
	return is.Email.Error(gotrans.T("email"))
}
