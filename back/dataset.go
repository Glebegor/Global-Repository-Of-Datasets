package grod

type Dataset struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          int    `json:"-" db:"id"`
}