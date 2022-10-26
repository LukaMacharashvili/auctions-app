package repositories

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (R *UserRepository) Fetch(id int64) (*sql.Rows, error) {
	rows, err := R.DB.Query("select id, balance from user where id = ?", id)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *UserRepository) DecreaseBalance(id int64, amount float64) (bool, error) {
	_, err := R.DB.Exec("update user set balance = balance - ? where id = ?", amount, id)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *UserRepository) IncreaseBalance(id int64, amount float64) (bool, error) {
	_, err := R.DB.Exec("update user set balance = balance + ? where id = ?", amount, id)

	if err != nil {
		return false, err
	}
	return true, nil
}
