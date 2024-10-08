package v1Router

import (
	v1Handlers "card-validator-service/internal/interface/rest/v1/handlers"
	"github.com/gin-gonic/gin"
)

const (
	ApiV1 = "/api/v1"
)

func GetGinRouter(handlers *v1Handlers.Handlers) *gin.Engine {
	const (
		GroupCard = "/card"
		Validate  = "/validate"
	)

	router := gin.Default()
	apiV1Routes := router.Group(ApiV1)
	cardGroup := apiV1Routes.Group(GroupCard)

	//GROUP: /card
	//Method: /validate
	cardGroup.POST(Validate, handlers.CardHandler.ValidateCard)

	return router
}
