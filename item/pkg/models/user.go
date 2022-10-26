package models

type Item struct {
	Id            int64
	Name          string
	Description   string
	StartingPrice float64
	Image         string
	Status        string
	HighestBid    float64
	OwnerId       int64
	HighestBidder int64
}

type User struct {
	Id      int64
	Balance float64
}
