package v1

import "github.com/gin-gonic/gin"

const (
	ApiV1 = "/api/v1"
)

func InitializeRouter(handlers interface{}) *gin.Engine {
	const (
		GroupCard = "/card"
	)

	const (
		CardValidate = "/validate"
	)

	router := gin.Default()
	apiV1Routes := router.Group(ApiV1)
	cardGroup := apiV1Routes.Group(GroupCard)

	cardGroup.POST(CardValidate)

	return router
}
