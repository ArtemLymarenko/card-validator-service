package mapper

import (
	"card-validator-service/internal/domain/model/card"
	"card-validator-service/internal/interface/rest/dto"
)

func FromCardValidationRequestDTOToCard(cardRequestDTO dto.CardValidationRequest) card.Card {
	return card.Card{
		Number:   card.Number(cardRequestDTO.Number),
		ExpYear:  card.Year(cardRequestDTO.ExpYear),
		ExpMonth: card.Month(cardRequestDTO.ExpMonth),
	}
}
