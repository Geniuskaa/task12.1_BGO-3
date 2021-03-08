package app

import (
	"encoding/json"
	"errors"
	"github.com/Geniuskaa/task12.1_BGO-3/cmd/bank/app/dto"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/postgreSQL"
	"github.com/Geniuskaa/task12.1_BGO-3/pkg/transaction"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{
		cardSvc: cardSvc,
		mux:     mux,
	}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	mapWithValues := r.URL.Query()
	value1 := mapWithValues["id"]
	id, err := strconv.Atoi(value1[0])
	if err != nil {
		log.Println(err)
		return
	}

	cards, err := postgreSQL.GetCards(int64(id))
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

	transactions, err := postgreSQL.GetTransactions(int64(id))
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
