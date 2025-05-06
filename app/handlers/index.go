package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/repository"
)

func (h *Handler) Index(c echo.Context) error {
	log.Println(c.RealIP())
    user, err := repository.GetUserByID(h.DB, currentUserID.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
    }

	followers, err := repository.Get_friend_list(h.DB, currentUserID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get followers"})
	}

	recs, err := repository.Get_friend_of_friend_list(h.DB, currentUserID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get recommendations"})
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title":           "ユーザーページ",
		"User":            user,
		"Followers":       followers,
		"Recommendations": recs,
	})
}
