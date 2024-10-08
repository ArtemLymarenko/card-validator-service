package v1Handlers

import (
	"card-validator-service/internal/interface/mapper"
	"card-validator-service/internal/interface/rest/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CardHandlerImpl struct{}

func NewCardHandler() *CardHandlerImpl {
	return &CardHandlerImpl{}
}

func (handler *CardHandlerImpl) ValidateCard(c *gin.Context) {
	var cardToValidate dto.CardValidationRequest
	if err := c.ShouldBindJSON(&cardToValidate); err != nil {
		logrus.WithError(err).Error("Failed to decode request body")
		c.JSON(http.StatusBadRequest, dto.NewCardValidationResponse(false, ErrFailedToDecodeBody))
		return
	}
	logrus.Info("JSON body was successfully read and decoded")

	card := mapper.FromCardValidationRequestDTOToCard(cardToValidate)
	logrus.Info("Request values were mapped to domain model")

	valid, err := card.Validate()
	if err != nil {
		logrus.WithError(err).Warn("Card validation failed")
		c.JSON(http.StatusBadRequest, dto.NewCardValidationResponse(valid, err))
		return
	}

	c.JSON(http.StatusOK, dto.NewCardValidationResponse(valid, err))
}
