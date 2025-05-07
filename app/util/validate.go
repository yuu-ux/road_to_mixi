package util

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("min_id", func(fl validator.FieldLevel) bool {
		idStr := fl.Field().String()
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return false
		}
		return id >= 1
	})
	return validate
}
