package repositories

import (
	"database/sql"

	"github.com/LukaMacharashvili/item/pkg/models"
)

type ItemRepository struct {
	DB *sql.DB
}

func (R *ItemRepository) Create(item models.Item) (bool, error) {
	_, err := R.DB.Exec("INSERT INTO item(name, description, startingPrice, image, ownerId, highestBidder) VALUES(?, ?, ?, ?, ?, ?)", item.Name, item.Description, item.StartingPrice, item.Image, item.OwnerId, item.HighestBidder)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *ItemRepository) Load() (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from item")

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *ItemRepository) Fetch(id int64) (*sql.Rows, error) {
	rows, err := R.DB.Query("select * from item where id = ?", id)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (R *ItemRepository) Delete(id int64) (bool, error) {
	_, err := R.DB.Exec("delete from item where id = ?", id)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (R *ItemRepository) UpdateHighestBid(id int64, bidder int64, amount float64) (bool, error) {
	_, err := R.DB.Exec("update item set highestBidder = ?, highestBid = ? where id = ?", bidder, amount, id)

	if err != nil {
		return false, err
	}
	return true, nil
}
