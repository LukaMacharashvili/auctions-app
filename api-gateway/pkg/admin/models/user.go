package models

type User struct {
	Id       int64
	Username string
	Email    string
	Password string
	Roles    string
	Balance  float64
}
