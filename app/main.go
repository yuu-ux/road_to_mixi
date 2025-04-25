package main

import (
    database "road_to_mixi/db"
	"net/http"
	"github.com/labstack/echo/v4"
    "github.com/joho/godotenv"
    "log"
    "road_to_mixi/handlers"
    "road_to_mixi/models"
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

    var results []struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }

	e.GET("/", func(c echo.Context) error {
        var users []models.User
        db.Find(&users)
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Title": "トップページ", "Users": users})
	})

    e.GET("/get_friend_list", func(c echo.Context) error {
        param := c.QueryParam("id")
        if err := db.Model(&models.FriendLink{}).
            Select("User2.user_id AS id, User2.name").
            Joins("User2").
            Where("friend_links.user1_id = ?", param).
            Scan(&results).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
        }
        if len(results) == 0 {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
        }
        return c.JSON(http.StatusOK, results)
    })

    e.GET("get_friend_of_friend_list", func(c echo.Context) error {
        param := c.QueryParam("id")
        subQuery := db.Model(&models.FriendLink{}).
            Select("user2_id").
            Where("user1_id = ?", param)
        if err := db.Model(&models.FriendLink{}).
            Select("User2.user_id AS id, User2.name AS name").
            Joins("User2").
            Where("friend_links.user1_id IN (?)", subQuery).
            Scan(&results).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
            }
        return c.JSON(http.StatusOK, results)
    })

	e.Logger.Fatal(e.Start(":1323"))
}
