package repositories

import (
	"database/sql"

	"github.com/LukaMacharashvili/wallet/pkg/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func (R *TransactionRepository) Create(transaction models.Transaction) (bool, error) {
	_, err := R.DB.Exec("INSERT INTO transaction(ownerId, amount) VALUES(?, ?)", transaction.OwnerId, transaction.Amount)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *TransactionRepository) LoadTransaction(ownerId int64) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from transaction where ownerId = ?", ownerId)

	if err != nil {
		return nil, err
	}
	return rows, nil
}
