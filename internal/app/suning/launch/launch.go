package launch

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type AppLauncher struct {
	logger *zap.Logger
}

func (c *AppLauncher) Start() error {
	return nil
}

func (c *AppLauncher) Stop() error {
	return nil
}

func NewAppLauncher(logger *zap.Logger) *AppLauncher {
	return &AppLauncher{
		logger: logger,
	}
}

var ProviderSet = wire.NewSet(NewAppLauncher)
