package dto

import (
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/transaction"
)

type CardDTO struct {
	Error *Error    `json:"error"`
	Card  card.Card `json:"card"`
}

type TransactionDTO struct {
	Error *Error `json:"error"`
	Transaction transaction.Transaction `json:"transaction"`
}

type MostVisitingPlaceDTO struct {
	Error *Error `json:"error"`
	Place *transaction.PopularPlace `json:"place"`
}

type BiggestSpendings struct {
	Error *Error `json:"error"`
	Spending *transaction.BiggestSpending `json:"spending"`
}

type Error struct {
	Code int `json:"error_code"`
	Message string `json:"error_msg"`
}
