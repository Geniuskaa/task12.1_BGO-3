package transaction

import "time"

type Transaction struct {
	Id int64 `json:"id"`
	Amount int64 `json:"amount"`
	MCC int64 `json:"mcc"`
	Date time.Time `json:"date"`
	CardId int64 `json:"card_id"`
	Receiver string `json:"receiver"`
}




