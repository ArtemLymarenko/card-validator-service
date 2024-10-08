package v1Router

import (
	"card-validator-service/internal/config"
	v1Handlers "card-validator-service/internal/interface/rest/v1/handlers"
	"github.com/gin-gonic/gin"
)

const (
	ApiV1 = "/api/v1"
)

func GetGinRouter(env config.Env, handlers *v1Handlers.Handlers) *gin.Engine {
	const (
		GroupCard = "/card"
		Validate  = "/validate"
	)

	gin.SetMode(string(env))
	router := gin.Default()
	v1 := router.Group(ApiV1)
	{
		cardGroup := v1.Group(GroupCard)
		{
			cardGroup.POST(Validate, handlers.CardHandler.ValidateCard)
		}
	}

	return router
}
