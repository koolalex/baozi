package suning

import (
	"github.com/google/wire"
	"github.com/koolalex/baozi/internal/pkg/app"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	Name string
	Port int
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	opt := new(Options)
	if err := v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return opt, err
}

func NewApp(opt *Options, logger *zap.Logger, launcher app.Launcher) (*app.Application, error) {
	a, err := app.New(opt.Name, logger, app.LauncherOption(launcher))

	if err != nil {
		return nil, errors.Wrap(err, "create new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)
