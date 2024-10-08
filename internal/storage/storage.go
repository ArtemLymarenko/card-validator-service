package storage

import (
	"card-validator-service/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	sslDisable    = "disable"
	sslVerifyFull = "verify-full"
)

func getSSLConfig(env config.Env) string {
	var sslConfig string
	if env == config.EnvLocal {
		sslConfig = sslDisable
	} else {
		sslConfig = sslVerifyFull
	}

	return sslConfig
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres(
	ps config.Postgres,
	env config.Env,
) (*Postgres, error) {
	connectionPath := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		ps.User,
		ps.Password,
		ps.Host,
		ps.Port,
		ps.Name,
		getSSLConfig(env),
	)

	db, err := sql.Open(ps.Dialect, connectionPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(ps.PoolMin)
	db.SetMaxOpenConns(ps.PoolMax)

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, ErrHealthCheckPostgres
	}

	logrus.Info("connected to database")

	m, err := migrate.New(
		"file://migrations",
		connectionPath,
	)

	if err != nil {
		_ = db.Close()
		return nil, ErrCreateMigrationPostgres
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		_ = db.Close()
		return nil, ErrMigratePostgres
	}

	logrus.Info("migrations created")

	return &Postgres{db}, nil
}

func (p *Postgres) CloseConnection() error {
	if err := p.db.Close(); err != nil {
		return ErrClosePostgresConnection
	}

	return nil
}

func (p *Postgres) GetConnection() (*sql.DB, error) {
	if p.db != nil {
		return p.db, nil
	}

	return nil, ErrGetPostgresConnection
}
