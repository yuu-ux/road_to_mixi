package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"strconv"
)

func (h *Handler) GetFriendList(c echo.Context) error {
	log.Println(c.RealIP())
	var query models.UserIDQuery
	if err := c.Bind(&query); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	if err := h.Validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Parameter"})
	}
	id, err := strconv.Atoi(query.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID must be integer"})
	}
	friends, err := repository.GetFriendList(h.DB, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
	}
	return c.Render(http.StatusOK, "friend_list.html", map[string]interface{}{
		"Title":   "フレンドリスト",
		"Friends": friends,
	})
}
