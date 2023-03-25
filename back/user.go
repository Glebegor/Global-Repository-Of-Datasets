package grod

type User struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Tel       string `json:"tel"`
	Subscribe string `json:"subscribe" binding:"required"`
	Id        int    `json:"-" db:"id"`
}
