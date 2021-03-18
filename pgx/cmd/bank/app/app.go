package app

import (
	"encoding/json"
	"errors"
	"github.com/Geniuskaa/task12.1_BGO-3/cmd/bank/app/dto"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/postgreSQL"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/transaction"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
	pool *pgxpool.Pool
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux, pool *pgxpool.Pool) *Server {

	return &Server{
		cardSvc: cardSvc,
		mux:     mux,
		pool: pool,
	}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
	s.mux.HandleFunc("/mostPopularPlace", s.mostPopularPlace)
	s.mux.HandleFunc("/biggestSpending", s.biggestSpending)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	mapWithValues := r.URL.Query()
	value1 := mapWithValues["id"]
	id, err := strconv.Atoi(value1[0])
	if err != nil {
		log.Println(err)
		return
	}

	cards, err := postgreSQL.GetCards(int64(id), s.pool)
	if err != nil {
		log.Println(err)
		return
	}

	if (len(cards) == 0) {
		w.WriteHeader(404)
		dtos := "There are not any cardHolders with this ID."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	dtos := make([]*dto.CardDTO, len(cards))
	counter := 0
	for i, c := range cards {
			counter++
			dtos[i] = &dto.CardDTO{
				Card: card.Card{
					Issuer:   c.Issuer,
					Number:   c.Number,
					Balance:  c.Balance,
					CardId:   c.CardId,
					HolderId: c.HolderId,
					Status:   c.Status,
					Holder:   c.Holder,
				},
			}
	}

	if counter == 0 {
		w.WriteHeader(404)
		dtos := "There are not any cardHolders with this ID."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}

func (s *Server) getTransactions(w http.ResponseWriter, r *http.Request) {
	mapWithValues := r.URL.Query()
	unParsedId := mapWithValues["id"]
	id, err := strconv.Atoi(unParsedId[0])
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	transactions, err := postgreSQL.GetTransactions(int64(id), s.pool)
	if err != nil {
		log.Println(err)
		return
	}

	if (len(transactions) == 0) {
		w.WriteHeader(404)
		dtos := "There are not any cardHolders with this ID."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	dtos := make([]*dto.TransactionDTO, len(transactions))
	counter := 0
	for i, c := range transactions {
		counter++
		dtos[i] = &dto.TransactionDTO{
			Transaction: transaction.Transaction{
				Id:       c.Id,
				Amount:   c.Amount,
				MCC:      c.MCC,
				Date:     c.Date,
				CardId:   c.CardId,
				Receiver: c.Receiver,
			},
		}
	}

	if counter == 0 {
		w.WriteHeader(404)
		dtos := "There are not any cardHolders with this ID."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}

func (s *Server) mostPopularPlace(w http.ResponseWriter, r *http.Request)  {
	place, err := postgreSQL.MostPopularPlace(s.pool)
	if (err != nil) {
		log.Println(err)
		return
	}

	if (place == nil) {
		w.WriteHeader(404)
		dtos := "Something went wrong..."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	dto := dto.MostVisitingPlaceDTO{
		Place: &transaction.PopularPlace{
			MCC:           place.MCC,
			CountOfVisits: place.CountOfVisits,
			NameOfPlace:   findNameOfMCC(place.MCC),
		},
	}

	respBody, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}

func findNameOfMCC(mcc int64) string {
	var mccName string
	switch mcc {
	case 5050:
		mccName = "Grocery"
	case 5010:
		mccName = "Bar"
	default:
		mccName = "Other"
	}
	return mccName
}

func (s *Server) biggestSpending(w http.ResponseWriter, r *http.Request)  {
	spending, err := postgreSQL.BiggestSpendings(s.pool)
	if (err != nil) {
		log.Println(err)
		return
	}

	if (spending == nil) {
		w.WriteHeader(404)
		dtos := "Something went wrong..."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New(dtos))
		return
	}

	dto := dto.BiggestSpendings{
		Spending: &transaction.BiggestSpending{
			MCC:         spending.MCC,
			Amount:      spending.Amount,
			NameOfPlace: findNameOfMCC(spending.MCC),
		},
	}

	respBody, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}
