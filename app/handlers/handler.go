package handlers

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var currentUserID = 1

const defaultLimit = 1

type Handler struct {
	DB       *gorm.DB
	Validate *validator.Validate
}
