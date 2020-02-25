package logger

import (
	"fmt"
	"go.uber.org/zap"
	"scheduler/internal/configs"
)

const (
	development = "development"
	production = "production"
)

func NewLogger(config *configs.Config) (*zap.Logger, error) {

	switch config.Mode {
	case development:
		return zap.NewDevelopment()
	case production:
		return zap.NewProduction()
	default:
		return nil, fmt.Errorf("unknown environment '%s', check the '%s' environment variable", config.Mode, "MODE")
	}
}
