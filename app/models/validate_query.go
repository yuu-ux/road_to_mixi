package models

type UserIDQuery struct {
	ID int `query:"id" form:"id" validate:"required,min=1"`
}

type UserPagingQuery struct {
	ID    int `query:"id" validate:"required,min=1"`
	Page  int `query:"page" validate:"required,min=1"`
	Limit int `query:"limit" validate:"required,min=1,max=100"`
}
