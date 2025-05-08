package models

type UserIDQuery struct {
	ID string `query:"id" form:"id" validate:"required,numeric,min_id"`
}

type UserPagingQuery struct {
	ID    string `query:"id" validate:"required,numeric,min_id"`
	Page  string `query:"page" validate:"required,numeric,min_id"`
	Limit string `query:"limit" validate:"required,numeric,min_id"`
}
