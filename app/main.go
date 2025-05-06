package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	database "road_to_mixi/db"
	"road_to_mixi/handlers"
	"road_to_mixi/models"
	"road_to_mixi/repository"
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

	currentUserID := models.UserIDQuery{ID: 1}
	const defaultLimit = 1

	e.GET("/", func(c echo.Context) error {
		log.Println(c.RealIP())
		var user models.User
		if err := db.First(&user, currentUserID.ID).Error; err != nil {
			return c.String(http.StatusInternalServerError, "ユーザー取得失敗")
		}
		followers, err := repository.Get_friend_list(db, currentUserID.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "フォロワー取得失敗"})
		}
		recs, err := repository.Get_friend_of_friend_list(db, currentUserID.ID)
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
		var query models.UserIDQuery
		if err := c.Bind(&query); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}
		if err := validate.Struct(query); err != nil {
			return c.JSON(http.StatusBadRequest, "id は必須です")
		}
		friends, err := repository.Get_friend_list(db, query.ID)
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
		var query models.UserIDQuery
		if err := c.Bind(&query); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		if err := validate.Struct(query); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}
		friends, err := repository.Get_friend_of_friend_list(db, query.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
		}
		return c.Render(http.StatusOK, "recommend_friend.html", map[string]interface{}{
			"Title":    "おすすめの友達",
			"ID":       query.ID,
			"Friends":  friends,
			"HasPrev":  false,
			"HasNext":  true,
			"PrevPage": 0,
			"NextPage": 2,
			"Limit":    defaultLimit,
		})
	})

	e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
		log.Println(c.RealIP())
		var query models.UserPagingQuery
		if err != c.Bind(&query) {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid request"})
		}
		log.Println(query)
		if err := validate.Struct(query); err != nil {
			return c.JSON(http.StatusBadRequest, "必須パラメータです")
		}
		friends, err := repository.Get_friend_of_friend_list_paging(db, query.ID, query.Page, query.Limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
		}
		HasPrev := query.Page > 1
		HasNext := len(friends) == query.Limit
		return c.Render(http.StatusOK, "recommend_friend.html", map[string]interface{}{
			"Title":    "おすすめの友達",
			"ID":       query.ID,
			"Friends":  friends,
			"HasPrev":  HasPrev,
			"HasNext":  HasNext,
			"PrevPage": query.Page - 1,
			"NextPage": query.Page + 1,
			"Limit":    query.Limit,
		})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"Title": "ログイン",
		})
	})

	e.POST("/login", func(c echo.Context) error {
		var query models.UserIDQuery
		if err := c.Bind(&query); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		if err := validate.Struct(query); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}
		currentUserID = query
		return c.Redirect(http.StatusSeeOther, "/")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
