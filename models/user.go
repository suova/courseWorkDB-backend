package models
type User struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
	Country string `json:"country"`
	Password string `json:"password"`
	Role int `json:"role"`
}
type Exist struct {
	Count    string `json:"count"`
}

type Role struct {
	Role    int `json:"role"`
}