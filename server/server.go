package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/attiss/rpi-switch/config"
	"github.com/attiss/rpi-switch/relaycontroller"
	"go.uber.org/zap"
)

type APIServer struct {
	handler    *requestHandler
	httpServer *http.Server
	logger     *zap.Logger
}

func (s *APIServer) Start(shutdownChannel chan os.Signal) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				s.logger.Error("listen and serve failure", zap.Error(err))
				panic(err)
			}
		}
	}()

	s.logger.Info("http server is running", zap.String("address", s.httpServer.Addr))

	<-shutdownChannel

	s.logger.Info("shutting down http server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown http server", zap.Error(err))
		panic(err)
	}

	s.handler.relayController.Close()
}

func NewAPIServer(config config.Config, logger *zap.Logger) (*APIServer, error) {
	relayController, err := relaycontroller.New(config, logger)
	if err != nil {
		return nil, err
	}
	if err := relayController.Init(); err != nil {
		return nil, err
	}

	handler := newRequestHandler(&relayController, logger)

	server := APIServer{
		handler: &handler,
		httpServer: &http.Server{
			Addr:         config.GetListenAddress(),
			WriteTimeout: time.Second * 5,
			ReadTimeout:  time.Second * 5,
			IdleTimeout:  time.Second * 5,
			Handler:      handler.GetRouter(),
		},
		logger: logger,
	}

	return &server, nil
}
