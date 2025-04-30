package main

import (
	"log"
	"net/http"
	database "road_to_mixi/db"
	"road_to_mixi/handlers"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"strconv"
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
	e := echo.New()
    handlers.SetDefault(e)
    database.InitDatabase(db)

	e.GET("/", func(c echo.Context) error {
        var users []models.User
        db.Find(&users)
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Title": "トップページ", "Users": users})
	})

    e.GET("/get_friend_list", func(c echo.Context) error {
        id := c.QueryParam("id")
        if id == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found id"})
        }
        friends, err := repository.Get_friend_list(db, id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
        }
        return c.JSON(http.StatusOK, friends)
    })

    e.GET("get_friend_of_friend_list", func(c echo.Context) error {
        id := c.QueryParam("id")
        if id == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found id"})
        }
        friends, err := repository.Get_friend_of_friend_list(db, id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
        }
        return c.JSON(http.StatusOK, friends)
    })

    e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
        id := c.QueryParam("id")
        limitStr := c.QueryParam("limit")
        pageStr := c.QueryParam("page")
        if id == "" || limitStr == "" || pageStr == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found id or limit or page"})
        }
        limit, err := strconv.Atoi(limitStr)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "conver error"})
        }
        page, err := strconv.Atoi(pageStr)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "conver error"})
        }
        friends, err := repository.Get_friend_of_friend_list_paging(db, id, page, limit)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
        }
        return c.JSON(http.StatusOK, friends)
    })

	e.Logger.Fatal(e.Start(":1323"))
}
