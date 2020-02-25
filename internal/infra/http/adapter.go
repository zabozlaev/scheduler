package http

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"scheduler/internal/configs"
	"scheduler/internal/infra/api"
)

type Adapter interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type adapter struct {
	logger *zap.Logger
	config *configs.Config
	server *http.Server
}

func NewAdapter(logger *zap.Logger, config *configs.Config, root *api.Adapter) Adapter {
	adapter := &adapter{
		logger: logger,
		config: config,
	}

	r, _ := adapter.newRouter(root)

	adapter.server = &http.Server{
		Handler: r,
		Addr: config.Port,
	}

	return adapter
}

func (a *adapter) Start() error {
	a.logger.Info("Starting http server")
	return a.server.ListenAndServe()
}

func (a *adapter) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
