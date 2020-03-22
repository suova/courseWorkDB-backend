package models
type User struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
	Country string `json:"country"`
	Password string `json:"password"`
	Is_admin bool `json:"is_admin"`
}
type Exist struct {
	Count    string `json:"count"`
}