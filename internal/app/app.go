package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	server *http.Server
}

func New(server *http.Server) *Application {
	return &Application{
		server: server,
	}
}

func (app *Application) Start() {
	idleConnsClosed := make(chan struct{})

	go app.gracefulShutDown(idleConnsClosed)

	logrus.Info(fmt.Sprintf("HTTP server addr: %s", app.server.Addr))
	if err := app.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Info("HTTP server ListenAndServe")
	}

	<-idleConnsClosed

	logrus.Info("HTTP server stopped")
}

func (app *Application) gracefulShutDown(idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Warn("http-server server shutdown")
	}

	close(idleConnsClosed)
}
