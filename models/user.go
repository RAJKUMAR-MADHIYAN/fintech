package models

type UserType string

const (
	Admin UserType = "admin"
	Users UserType = "user"
)

type User struct {
	ID        int  `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	Phone     string   `json:"phone"`
	UserType  UserType `json:"user_type"`
}
