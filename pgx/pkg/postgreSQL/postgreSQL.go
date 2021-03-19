package postgreSQL

import (
	"context"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/transaction"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client struct {
	Id int64
	FullName string
	Birthday time.Time
}

func GetCards(OwnerId int64, pool *pgxpool.Pool) ([]*card.Card, error){

	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Release()

	cards := make([]*card.Card, 0)
	rows, err := conn.Query(ctx, `
		SELECT id, number, balance, issuer, holder, status FROM cards WHERE owner_id = $1 LIMIT 5`, OwnerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		card := &card.Card{}
		err = rows.Scan(&card.CardId, &card.Number, &card.Balance, &card.Issuer, &card.Holder, &card.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		card.HolderId = OwnerId;
		cards = append(cards, card)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return cards, nil
}

func GetTransactions(CardId int64, pool *pgxpool.Pool) ([]*transaction.Transaction, error){

	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Release()

	transactions := make([]*transaction.Transaction, 0)
	rows, err := conn.Query(ctx, `
		SELECT id, sum, mcc, receiver, created FROM transactions WHERE card_id = $1 LIMIT 10`, CardId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transaction := &transaction.Transaction{}
		err = rows.Scan(&transaction.Id, &transaction.Amount, &transaction.MCC, &transaction.Receiver, &transaction.Date)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		transaction.CardId = CardId;
		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transactions, nil
}

func MostPopularPlace(pool *pgxpool.Pool, id int64) (*transaction.PopularPlace, error){ // mcc, count, err

	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Release()

	place := &transaction.PopularPlace{}
	err = conn.QueryRow(ctx, `SELECT mcc, count(*) OthenVisitingPlaces FROM transactions WHERE card_id = $1
	GROUP BY mcc ORDER BY OthenVisitingPlaces DESC LIMIT 1`, id).Scan(&place.MCC, &place.CountOfVisits)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return place, nil
}

func BiggestSpendings(pool *pgxpool.Pool, id int64) (*transaction.BiggestSpending, error) {

	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Release()

	spending := &transaction.BiggestSpending{}
	err = conn.QueryRow(ctx, `SELECT mcc, sum(sum) total FROM transactions WHERE card_id = $1 
	GROUP BY mcc ORDER BY total DESC LIMIT 1`, id).Scan(&spending.MCC, &spending.Amount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return spending, nil
}

