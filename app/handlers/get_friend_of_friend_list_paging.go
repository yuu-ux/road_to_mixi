package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"road_to_mixi/models"
	"road_to_mixi/repository"
)

func (h *Handler) GetFriendOfFriendListPaging(c echo.Context) error {
	log.Println(c.RealIP())
	var query models.UserPagingQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid request"})
	}
	log.Println(query)
	if err := h.Validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Parameter"})
	}
	friends, err := repository.Get_friend_of_friend_list_paging(h.DB, query.ID, query.Page, query.Limit)
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
}
