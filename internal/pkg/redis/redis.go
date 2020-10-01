package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options is  configuration of database
type Options struct {
	Addr string `yaml:"url"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

// Init 初始化数据库
func New(o *Options) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     o.Addr,
		Password: o.Pass, // no password set
		DB:       o.DB,   // use default DB
	})

	var ctx = context.Background()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	if err != nil {
		return nil, errors.Wrap(err, "redis open database connection error")
	}

	return rdb, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
