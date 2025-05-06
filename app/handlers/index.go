package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/models"
	"road_to_mixi/repository"
)

func (h *Handler) Index(c echo.Context) error {
	log.Println(c.RealIP())
	var user models.User
	if err := h.DB.First(&user, currentUserID.ID).Error; err != nil {
		return c.String(http.StatusInternalServerError, "ユーザー取得失敗")
	}
	followers, err := repository.Get_friend_list(h.DB, currentUserID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "フォロワー取得失敗"})
	}
	recs, err := repository.Get_friend_of_friend_list(h.DB, currentUserID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "おすすめ取得失敗"})
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title":           "ユーザーページ",
		"User":            user,
		"Followers":       followers,
		"Recommendations": recs,
	})
}
