package models

type UserType string

const (
	Admin   UserType = "admin"
	Regular UserType = "regular"
	Guest   UserType = "guest"
)

type User struct {
	ID       int      `json:"id"`
	UserName string   `json:"user_name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Phone    string   `json:"phone"`
	UserType UserType `json:"user_type"`
}
