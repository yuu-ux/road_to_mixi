package main

import (
    database "road_to_mixi/db"
    "strconv"
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

    type friend struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }

	e.GET("/", func(c echo.Context) error {
        var users []models.User
        db.Find(&users)
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{"Title": "トップページ", "Users": users})
	})

    e.GET("/get_friend_list", func(c echo.Context) error {
        var friends []friend
        id := c.QueryParam("id")
        if err := db.Model(&models.FriendLink{}).
            Select("User2.user_id AS id, User2.name").
            Joins("User2").
            Where("friend_links.user1_id = ?", id).
            Scan(&friends).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
        }
        return c.JSON(http.StatusOK, friends)
    })

    e.GET("get_friend_of_friend_list", func(c echo.Context) error {
        var friends []friend
        id := c.QueryParam("id")
        subQuery := db.Model(&models.FriendLink{}).
            Select("user2_id").
            Where("user1_id = ?", id)
        var blockees []int
        var blockers []int
        db.Model(&models.BlockList{}).
            Where("user1_id = ?", id).
            Pluck("user2_id", &blockees)
        db.Model(&models.BlockList{}).
            Where("user2_id = ?", id).
            Pluck("user1_id", &blockers)
        blockedIDs := append(blockees, blockers...)
        if err := db.Model(&models.FriendLink{}).
            Select("User2.user_id AS id, User2.name AS name").
            Joins("User2").
            Where("friend_links.user1_id IN (?)", subQuery).
            Where("friend_links.user2_id NOT IN (?)", subQuery).
            Where("friend_links.user2_id NOT IN (?)", blockedIDs).
            Scan(&friends).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
            }
        return c.JSON(http.StatusOK, friends)
    })

    e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
        id := c.QueryParam("id")
        limitStr := c.QueryParam("limit")
        pageStr := c.QueryParam("page")
        limit, err := strconv.Atoi(limitStr)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "conver error"})
        }
        page, err := strconv.Atoi(pageStr)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "conver error"})
        }
        var friends []friend

        subQuery := db.Model(&models.FriendLink{}).
            Select("user2_id").
            Where("user1_id = ?", id)
        var blockees []int
        var blockers []int
        db.Model(&models.BlockList{}).
            Where("user1_id = ?", id).
            Pluck("user2_id", &blockees)
        db.Model(&models.BlockList{}).
            Where("user2_id = ?", id).
            Pluck("user1_id", &blockers)
        blockedIDs := append(blockees, blockers...)
        if err := db.Model(&models.FriendLink{}).
            Select("User2.user_id AS id, User2.name AS name").
            Joins("User2").
            Where("friend_links.user1_id IN (?)", subQuery).
            Where("friend_links.user2_id NOT IN (?)", subQuery).
            Where("friend_links.user2_id NOT IN (?)", blockedIDs).
            Offset((page - 1) * limit).
            Limit(limit).
            Scan(&friends).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
            }
        return c.JSON(http.StatusOK, friends)
    })

	e.Logger.Fatal(e.Start(":1323"))
}
