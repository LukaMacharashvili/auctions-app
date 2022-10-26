package repositories

import (
	"database/sql"

	"github.com/LukaMacharashvili/auth/pkg/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func (R *AuthRepository) Create(user models.User) (bool, error) {
	_, err := R.DB.Exec("INSERT INTO user(username, email, password, balance) VALUES(?, ?, ?, ?)", user.Username, user.Email, user.Password, user.Balance)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *AuthRepository) Fetch(email string) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from user where email = ?", email)

	if err != nil {
		return nil, err
	}
	return rows, nil
}
