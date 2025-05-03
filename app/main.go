package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	database "road_to_mixi/db"
	"road_to_mixi/handlers"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"strconv"
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
	handlers.SetDefault(e)

	const currentUserID = "1"

	e.GET("/", func(c echo.Context) error {
		log.Println(c.RealIP())
		var user models.User
		if err := db.First(&user, currentUserID).Error; err != nil {
			return c.String(http.StatusInternalServerError, "ユーザー取得失敗")
		}
		followers, err := repository.Get_friend_list(db, currentUserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "フォロワー取得失敗"})
		}
		recs, err := repository.Get_friend_of_friend_list(db, currentUserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "おすすめ取得失敗"})
		}
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Title":           "ユーザーページ",
			"User":            user,
			"Followers":       followers,
			"Recommendations": recs,
		})
	})

	e.GET("/get_friend_list", func(c echo.Context) error {
		log.Println(c.RealIP())
		id := c.QueryParam("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found id"})
		}
		friends, err := repository.Get_friend_list(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
		}
		return c.Render(http.StatusOK, "friend_list.html", map[string]interface{}{
			"Title":   "フレンドリスト",
			"Friends": friends,
		})
	})

	e.GET("get_friend_of_friend_list", func(c echo.Context) error {
		log.Println(c.RealIP())
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
		log.Println(c.RealIP())
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
