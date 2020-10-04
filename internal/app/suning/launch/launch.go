package launch

import (
	"fmt"
	"github.com/google/wire"
	"github.com/koolalex/baozi/internal/app/suning/services"
	"github.com/koolalex/baozi/internal/pkg/app"
	"go.uber.org/zap"
)

type AppLauncher struct {
	logger *zap.Logger
}

func (c *AppLauncher) Start() error {
	url := `http://product.suning.com/0000000000/144016246.html`
	price := services.GetGoodPrice(url)
	fmt.Println(price)
	return nil
}

func (c *AppLauncher) Stop() error {
	return nil
}

func NewAppLauncher(logger *zap.Logger) app.Launcher {
	return &AppLauncher{
		logger: logger,
	}
}

var ProviderSet = wire.NewSet(NewAppLauncher)
