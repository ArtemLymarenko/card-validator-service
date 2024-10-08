package dto

type CardRequest struct {
	Number   string `json:"number"`
	ExpYear  int    `json:"expYear"`
	ExpMonth int    `json:"expMonth"`
}

type CardError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CardResponse struct {
	Valid bool      `json:"valid"`
	Error CardError `json:"error"`
}
