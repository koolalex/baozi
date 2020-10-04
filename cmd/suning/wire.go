// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/koolalex/baozi/internal/app/suning"
	"github.com/koolalex/baozi/internal/app/suning/launch"
	"github.com/koolalex/baozi/internal/pkg/app"
	"github.com/koolalex/baozi/internal/pkg/config"
	"github.com/koolalex/baozi/internal/pkg/log"
)

var providerSet = wire.NewSet(
	//common
	config.ProviderSet,
	log.ProviderSet,

	//biz
	launch.ProviderSet,
	suning.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	//panic防止被执行，这里是要用wire来生成代码的，不能被执行
	panic(wire.Build(providerSet))
}
