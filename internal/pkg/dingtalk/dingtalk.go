package dingtalk

import (
	"github.com/blinkbean/dingtalk"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options is  configuration of database
type Options struct {
	token string `yaml:"token"`
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("dingtalk", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal dingtalk option error")
	}

	logger.Info("load dingtalk options success", zap.String("token", o.token))

	return o, err
}

// Init 初始化dingtalk
func New(o *Options) *dingtalk.DingTalk {
	operator := dingtalk.InitDingTalk([]string{o.token}, ".")

	return operator
}

var ProviderSet = wire.NewSet(New, NewOptions)
