package card

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Card struct {
	Issuer   string `json:"issuer"`
	Number   string `json:"number"`
	Balance  int64  `json:"balance"`
	CardId   int64  `json:"card_id"`
	HolderId int64  `json:"holder_id"`
	Status 	 string `json:"status"`
	Holder 	 string `json:"holder"`
}

type Service struct {
	ids int64
	mu sync.RWMutex
	cards []*Card
}

func NewService() *Service {
	return &Service{
		ids: 3,
		mu:    sync.RWMutex{},
		cards: []*Card{},
	}
}

func (s *Service) CardAdding(yourId int64, issuer string) error {
	var n string
	if yourId > 10 && yourId < 100 {
		n = fmt.Sprintf("00%d", yourId + 1)
	}
	n = fmt.Sprintf("000%d", yourId + 1)

	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63() % 100000

	for _, element := range s.cards{
		if element.HolderId == yourId {
			s.cards = append(s.cards, &Card{
				Issuer:   issuer,
				Number:   n,
				Balance:  0,
				CardId: randId,
				HolderId: yourId,
			})
			return nil
		}
	}

	return errors.New("There is not cardHolder with this Id, before Adding the card, register in our Bank.")

}

func (s *Service) All(ctx context.Context) []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cards
}