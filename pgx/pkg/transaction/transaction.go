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

type PopularPlace struct {
	MCC int64`json:"mcc"`
	CountOfVisits int64 `json:"count_of_visits"`
	NameOfPlace string `json:"name_of_place"`
}

type BiggestSpending struct {
	MCC int64`json:"mcc"`
	Amount int64 `json:"amount"`
	NameOfPlace string `json:"name_of_place"`
}




