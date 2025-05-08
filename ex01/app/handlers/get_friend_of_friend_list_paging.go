package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/models"
	"road_to_mixi/repository"
	"strconv"
)

func (h *Handler) GetFriendOfFriendListPaging(c echo.Context) error {
	log.Println(c.RealIP())

	var query models.UserPagingQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid request"})
	}
	if err := h.Validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Parameter"})
	}
	id, err := strconv.Atoi(query.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID must be integer"})
	}
	page, err := strconv.Atoi(query.Page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Page must be integer"})
	}
	limit, err := strconv.Atoi(query.Limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Limit must be integer"})
	}
	friends, err := repository.GetFriendOfFriendListPaging(h.DB, id, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed"})
	}
	HasPrev := page > minPage
	HasNext := len(friends) == limit
	return c.Render(http.StatusOK, "recommend_friend.html", map[string]interface{}{
		"Title":    "おすすめの友達",
		"ID":       query.ID,
		"Friends":  friends,
		"HasPrev":  HasPrev,
		"HasNext":  HasNext,
		"PrevPage": page - 1,
		"NextPage": page + 1,
		"Limit":    query.Limit,
	})
}
