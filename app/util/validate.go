package util

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

func NewValidator() *validator.Validate {
    const minID = 1

	validate := validator.New()
	validate.RegisterValidation("min_id", func(fl validator.FieldLevel) bool {
		idStr := fl.Field().String()
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return false
		}
		return id >= minID
	})
	return validate
}
