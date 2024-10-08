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
		c.JSON(http.StatusBadRequest, dto.NewCardValidationResponse(false, ErrFailedToDecodeBody))
		return
	}

	card := mapper.FromCardRequestDTOToCard(cardToValidate)
	logrus.Info("mapped requests value", card)

	valid, err := card.Validate()
	logrus.Info("card validated", card)

	response := dto.NewCardValidationResponse(valid, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
