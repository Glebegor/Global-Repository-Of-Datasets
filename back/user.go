package grod

type User struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Tel           string `json:"tel"`
	TimeSub       int    `json:"time_of_sub"`
	Subscribe     string `json:"subscribe"`
	Id            int    `json:"id" db:"id"`
	CountRequests int    `json: "count_requests" db:"count_requests"`
}
