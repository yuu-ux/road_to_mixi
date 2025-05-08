package test

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	database "road_to_mixi/db"
	"road_to_mixi/handlers"
	"road_to_mixi/util"
	"testing"
)

type MockRenderer struct{}

func (m *MockRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return nil
}

func InitHandler() *handlers.Handler {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.New()
	if err != nil {
		log.Fatal("Error DB")
	}
	validate := util.NewValidator()
	h := &handlers.Handler{DB: db, Validate: validate}
	return h
}

func testGetFriendListRequest(t *testing.T, e *echo.Echo, path string, expectedStatus int) {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != expectedStatus {
		t.Errorf("for path %q, expected status %d, got %d", path, expectedStatus, rec.Code)
	}
}

func TestGetFriendListHandler(t *testing.T) {
	e := echo.New()
	h := InitHandler()
	e.Renderer = &MockRenderer{}
	e.GET("/get_friend_list", h.GetFriendList)
	e.GET("/get_friend_of_friend_list", h.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", h.GetFriendOfFriendListPaging)

	testCases := []struct {
		path           string
		expectedStatus int
	}{
		// 正常系
		{"/get_friend_list?id=1", http.StatusOK},
		{"/get_friend_of_friend_list?id=1", http.StatusOK},
		{"/get_friend_of_friend_list_paging?id=1&page=1&limit=1", http.StatusOK},

		// パラメータなし
		{"/get_friend_list", http.StatusBadRequest},
		{"/get_friend_of_friend_list", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging", http.StatusBadRequest},

		// 文字のパラメータ
		{"/get_friend_list?id=a", http.StatusBadRequest},
		{"/get_friend_of_friend_list?id=a", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=a&page=1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=a&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=1&limit=a", http.StatusBadRequest},

		// 1より小さい
		{"/get_friend_list?id=-1", http.StatusBadRequest},
		{"/get_friend_list?id=0", http.StatusBadRequest},
		{"/get_friend_of_friend_list?id=-1", http.StatusBadRequest},
		{"/get_friend_of_friend_list?id=0", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=0&page=1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=0&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=1&limit=0", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=-1&page=1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=-1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=1&limit=-1", http.StatusBadRequest},

		// キーしかない
		{"/get_friend_list?id=", http.StatusBadRequest},
		{"/get_friend_of_friend_list?id=", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=", http.StatusBadRequest},

		// 大きすぎる値
		{"/get_friend_list?id=99999999999999999999", http.StatusBadRequest},
		{"/get_friend_of_friend_list?id=99999999999999999999", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=99999999999999999999", http.StatusBadRequest},

		// パラメータが足りない
		{"/get_friend_of_friend_list_paging?page=1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&limit=1", http.StatusBadRequest},
		{"/get_friend_of_friend_list_paging?id=1&page=1", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			testGetFriendListRequest(t, e, tc.path, tc.expectedStatus)
		})
	}
}
