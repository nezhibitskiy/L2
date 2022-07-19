package server

import (
	"fmt"
	"net"
	"net/http"

	cfg "11/internal/config"
	"11/internal/logging"
	"11/internal/usercases"
)

const (
	loggingPath = "./internal/logging/logs/logs.json"
)

type Application struct {
	mux    *http.ServeMux
	hash   usercases.Repository
	config *cfg.Cfg
	logger logging.LoggerEx
}

func NewApp(config *cfg.Cfg, repo usercases.Repository) *Application {
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

	a.mux = NewServ(a.hash, a.logger)

	err = a.logger.WriteInfo(fmt.Sprintf("Starting server on %s:%s", a.config.Ip, a.config.Port))
	if err != nil {
		return err
	}

	return http.Serve(listener, a.mux)
}
