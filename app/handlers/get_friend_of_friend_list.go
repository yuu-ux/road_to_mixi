package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"strconv"
)

func (h *Handler) GetFriendOfFriendList(c echo.Context) error {
	log.Println(c.RealIP())
	const nextPage = 2

	var query models.UserIDQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := h.Validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Parameter"})
	}
	id, err := strconv.Atoi(query.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID must be integer"})
	}
	friends, err := repository.GetFriendOfFriendList(h.DB, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
	}
	return c.Render(http.StatusOK, "recommend_friend.html", map[string]interface{}{
		"Title":    "おすすめの友達",
		"ID":       query.ID,
		"Friends":  friends,
		"HasPrev":  false,
		"HasNext":  true,
		"PrevPage": minPage,
		"NextPage": nextPage,
		"Limit":    defaultLimit,
	})
}
