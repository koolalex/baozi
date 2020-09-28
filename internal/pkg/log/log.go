package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		opt = new(Options)
	)

	if err = v.UnmarshalKey("log", opt); err != nil {
		return nil, err
	}

	return opt, nil
}

func New(opt *Options) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	if err = level.UnmarshalText([]byte(opt.Level)); err != nil {
		return nil, err
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize, //megabytes
		MaxBackups: opt.MaxBackups,
		MaxAge:     opt.MaxAge, //day
	})

	cw := zapcore.Lock(os.Stdout)

	//file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	// stdout core采用ConsoleEncoder
	if opt.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, err
}

var ProviderSet = wire.NewSet(New, NewOptions)
