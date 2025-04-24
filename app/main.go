package main

import (
    database "road_to_mixi/db"
	"net/http"
	"github.com/labstack/echo/v4"
    "github.com/joho/godotenv"
    "log"
    "road_to_mixi/handlers"
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
	defer db.Close()
	e := echo.New()
    handlers.SetDefault(e)
	e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Title": "トップページ",})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
