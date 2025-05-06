package main

import (
	"log"
	database "road_to_mixi/db"
	"road_to_mixi/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	// database.InitDatabase(db)

	e := echo.New()
	handlers.SetDefault(e)
	validate := validator.New()
	h := &handlers.Handler{DB: db, Validate: validate}
	e.GET("/", h.Index)
	e.GET("/login", h.GetLogin)
	e.POST("/login", h.PostLogin)
	e.GET("/get_friend_list", h.GetFriendList)
	e.GET("get_friend_of_friend_list", h.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", h.GetFriendOfFriendListPaging)

	e.Logger.Fatal(e.Start(":1323"))
}
