package v1Handlers

import (
	"card-validator-service/internal/config"
	"card-validator-service/internal/domain/repository/cache"
	"card-validator-service/internal/domain/repository/postgres"
	"github.com/gin-gonic/gin"
)

type CardHandler interface {
	ValidateCard(c *gin.Context)
}

type Handlers struct {
	CardHandler CardHandler
}

func New(cfg *config.Config) *Handlers {
	//repos to show how I'd extend the logic
	cardIssCache := cache.NewCardIssuerCache(cfg.Cache.Cap)
	_ = postgres.NewCardIssuerRepository(nil, cardIssCache)

	//handlers
	cardHandler := NewCardHandler()

	return &Handlers{
		CardHandler: cardHandler,
	}
}
