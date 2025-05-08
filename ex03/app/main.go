package main

import (
	"log"
	database "road_to_mixi/db"
	"road_to_mixi/router"
	"road_to_mixi/util"

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
	database.InitDatabase(db)
	e := echo.New()
	util.SetDefault(e)
	router.Router(db, e)
	e.Logger.Fatal(e.Start(":1323"))
}
