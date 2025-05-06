package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"road_to_mixi/models"
)

func (h *Handler) GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"Title": "ログイン",
	})
}

func (h *Handler) PostLogin(c echo.Context) error {
	var query models.UserIDQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := h.Validate.Struct(query); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	currentUserID = query
	return c.Redirect(http.StatusSeeOther, "/")
}
