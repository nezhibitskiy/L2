package server

import (
	"fmt"
	"net"
	"net/http"

	"11/internal"
	"11/internal/logging"
	"11/internal/usecase"
)

const (
	loggingPath = "./internal/logging/logs/logs.json"
)

type Application struct {
	mux    *http.ServeMux
	hash   usecase.Repository
	config *internal.Config
	logger logging.LoggerEx
}

func NewApp(config *internal.Config, repo usecase.Repository) *Application {
	return &Application{
		hash:   repo,
		config: config,
	}
}

func (a *Application) Start() error {

	listener, err := net.Listen("tcp", net.JoinHostPort(a.config.Ip, a.config.Port))
	if err != nil {
		return err
	}

	a.logger, err = logging.NewLogger(loggingPath)

	a.mux = Register(a.hash, a.logger)

	err = a.logger.WriteInfo(fmt.Sprintf("Starting server on %s:%s", a.config.Ip, a.config.Port))
	if err != nil {
		return err
	}

	return http.Serve(listener, a.mux)
}
