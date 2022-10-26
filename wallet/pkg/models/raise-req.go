package models

type RaiseReq struct {
	Id      int64
	OwnerId int64
	Amount  float64
	Status  string
}
