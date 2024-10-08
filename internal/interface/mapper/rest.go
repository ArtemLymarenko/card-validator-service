package mapper

import (
	"card-validator-service/internal/domain/model/card"
	"card-validator-service/internal/interface/rest/dto"
)

func FromCardRequestDTOToCard(cardDTO dto.CardValidationRequest) card.Card {
	return card.Card{
		Number:   card.Number(cardDTO.Number),
		ExpYear:  card.Year(cardDTO.ExpYear),
		ExpMonth: card.Month(cardDTO.ExpMonth),
	}
}
