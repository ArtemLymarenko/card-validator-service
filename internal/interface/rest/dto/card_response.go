package dto

import (
	"card-validator-service/internal/domain/model/card"
	"errors"
)

type StatusCode string

const (
	StatusInvalidCardNumber StatusCode = "001"
	StatusInvalidCardYear   StatusCode = "002"
	StatusInvalidCardMonth  StatusCode = "003"
	StatusBadRequest        StatusCode = "004"
)

type CardValidationError struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message"`
}

type CardValidationResponse struct {
	Valid bool                 `json:"valid"`
	Error *CardValidationError `json:"error,omitempty"`
}

func NewCardValidationError(err error) *CardValidationError {
	var code StatusCode

	switch {
	case errors.Is(err, card.ErrInvalidLuhnCheck):
		code = StatusInvalidCardNumber
	case errors.Is(err, card.ErrInvalidCardYear):
		code = StatusInvalidCardYear
	case errors.Is(err, card.ErrInvalidCardMonth):
		code = StatusInvalidCardMonth
	default:
		code = StatusBadRequest
	}

	if err != nil {
		return &CardValidationError{
			Code:    code,
			Message: err.Error(),
		}
	}

	return nil
}

func NewCardValidationResponse(valid bool, err error) CardValidationResponse {
	return CardValidationResponse{
		valid,
		NewCardValidationError(err),
	}
}
