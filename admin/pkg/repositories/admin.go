package repositories

import (
	"database/sql"

	"github.com/LukaMacharashvili/admin/pkg/models"
)

type AdminRepository struct {
	DB *sql.DB
}

func (R *AdminRepository) Create(user models.User) (bool, error) {
	_, err := R.DB.Exec("INSERT INTO user(username, email, password, role, balance) VALUES(?, ?, ?, ?, ?)", user.Username, user.Email, user.Password, user.Roles, user.Balance)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *AdminRepository) Load() (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from user")

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *AdminRepository) Fetch(email string) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from user where email = ?", email)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *AdminRepository) Delete(id int64) (bool, error) {
	_, err := R.DB.Exec("delete from user where id = ?", id)

	if err != nil {
		return false, err
	}
	return true, nil
}
