package repositories

import (
	"database/sql"

	"github.com/LukaMacharashvili/wallet/pkg/models"
)

type RaiseReqRepository struct {
	DB *sql.DB
}

func (R *RaiseReqRepository) Fetch(id int64) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from raise_req where id = ?", id)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *RaiseReqRepository) Create(raisereq models.RaiseReq) (bool, error) {
	_, err := R.DB.Exec("INSERT INTO raise_req(ownerId, amount, status) VALUES(?, ?, ?)", raisereq.OwnerId, raisereq.Amount, raisereq.Status)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *RaiseReqRepository) LoadRaiseReq(ownerId int64) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from raise_req where ownerId = ?", ownerId)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *RaiseReqRepository) UpdateStatus(id int64, status string) (bool, error) {
	_, err := R.DB.Exec("update raise_req set status = ? where id = ?", status, id)

	if err != nil {
		return false, err
	}

	return true, nil
}
