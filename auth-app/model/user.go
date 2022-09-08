package model

type User struct {
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}
