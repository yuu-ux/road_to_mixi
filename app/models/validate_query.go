package models

type UserIDQuery struct {
	ID int `query:"id" form:"id" validate:"required"`
}

type UserPagingQuery struct {
	ID    int `query:"id" validate:"required"`
	Page  int `query:"page" validate:"required"`
	Limit int `query:"limit" validate:"required"`
}
