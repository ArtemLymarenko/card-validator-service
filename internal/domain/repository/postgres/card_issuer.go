package postgres

import (
	"database/sql"
)

type Cache interface {
	Add(iin string, issuer interface{})
	Get(iin string) interface{}
}

type CardIssuerRepository struct {
	db    *sql.DB
	cache Cache
}

func NewCardIssuerRepository(db *sql.DB, cache Cache) *CardIssuerRepository {
	return &CardIssuerRepository{
		db:    db,
		cache: cache,
	}
}

func (c *CardIssuerRepository) FindByIIn(iin string) (iss interface{}) {
	//Impl
	return iss
}
