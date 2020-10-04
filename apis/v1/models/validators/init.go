package modelvalidators

import (
	"gin-app/apis/v1/config/validator"
	"gin-app/apis/v1/models"
)

func InitModelValidator() {
	v := validator.GetValidator()
	v.RegisterStructValidation(ValidatePostValuer, models.Post{})
}
