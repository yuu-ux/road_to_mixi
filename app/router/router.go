package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"road_to_mixi/handlers"
)

func Router(db *gorm.DB, e *echo.Echo) {
	validate := validator.New()
	h := &handlers.Handler{DB: db, Validate: validate}
	e.GET("/", h.Index)
	e.GET("/login", h.GetLogin)
	e.POST("/login", h.PostLogin)
	e.GET("/get_friend_list", h.GetFriendList)
	e.GET("/get_friend_of_friend_list", h.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", h.GetFriendOfFriendListPaging)
}
