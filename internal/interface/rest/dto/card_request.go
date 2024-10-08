package dto

type CardValidationRequest struct {
	Number   string `json:"number"`
	ExpYear  int    `json:"expYear"`
	ExpMonth int    `json:"expMonth"`
}
